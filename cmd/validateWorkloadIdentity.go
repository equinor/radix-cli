// Copyright © 2023
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	radixapiclient "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/client/configuration"
	"github.com/equinor/radix-cli/generated/radixapi/client/deployment"
	"github.com/equinor/radix-cli/generated/radixapi/client/platform"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/auth"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/workloadidentity"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/fatih/color"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var validateWorkloadIdentityCmd = &cobra.Command{
	Use:     "workload-identity",
	Short:   "Validate radixconfig.yaml",
	Long:    `Valida workload identity configuration`,
	Example: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		var appNames []string
		if len(appName) > 0 {
			appNames = []string{appName}
		} else {
			apps, err := apiClient.Platform.ShowApplications(&platform.ShowApplicationsParams{Context: cmd.Context()}, nil)
			if err != nil {
				return err
			}

			appNames = slice.Map(apps.Payload, func(app *models.ApplicationSummary) string { return *app.Name })
			completion.UpdateAppNamesCache(appNames)
		}

		fedCredValidator := federatedCredentialsValidationHelper{
			apiClient: apiClient,
			logger: func(msg string) {
				fmt.Fprintln(os.Stderr, msg)
			}}
		clientFedCreds, err := fedCredValidator.ValidateFederatedCredentialsDetails(cmd.Context(), appNames)

		radixAuth, err := auth.New()
		if err != nil {
			return err
		}

		graphHelper, err := workloadidentity.NewAzureServicePrincipalService(&tokenCredentialAdapter{auth: radixAuth})
		if err != nil {
			return fmt.Errorf("error initializing client: %w", err)
		}

		for clientId, expectedFedCreds := range clientFedCreds {

			fmt.Fprintf(os.Stderr, "Analyzing service principal with client id %v: ", clientId)
			sp, err := graphHelper.GetServicePrincipal(cmd.Context(), clientId)
			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "%s (%s). ", sp.DisplayName, sp.Type)

			existingWorkloadIdentityFedCreds := slice.FindAll(sp.FederatedCredentials, func(fedCred workloadidentity.FederatedCredential) bool {
				subjectParts := strings.Split(fedCred.Subject, ":")
				if len(subjectParts) != 4 {
					return false
				}

				return subjectParts[0] == "system" && subjectParts[1] == "serviceaccount"
			})
			mappedExistingWorkloadIdentityFedCreds := slice.Map(existingWorkloadIdentityFedCreds, func(c workloadidentity.FederatedCredential) FederatedCredential {
				return FederatedCredential{FederatedCredential: c}
			})
			missingFedCreds := findMissingFedCreds(expectedFedCreds, mappedExistingWorkloadIdentityFedCreds)
			extraFedCreds := findMissingFedCreds(mappedExistingWorkloadIdentityFedCreds, expectedFedCreds)
			fmt.Fprintf(os.Stderr, "(total: %v, missing: %v, extra: %v)\n", len(expectedFedCreds), len(missingFedCreds), len(extraFedCreds))

			for _, missingFedCred := range missingFedCreds {
				command, err := generateCreateFederatedCredentialAzureCLICommand(*sp, missingFedCred)
				if err != nil {
					return err
				}
				color.RGB(128, 128, 128).Fprintf(os.Stderr, "# %s\n", missingFedCred.Reason)
				color.Green(command)
			}

			for _, extraFedCred := range extraFedCreds {
				command, err := generateDeleteFederatedCredentialAzureCLICommand(*sp, extraFedCred)
				if err != nil {
					return err
				}
				color.Red(command)
			}
		}

		return nil
	},
}

type FederatedCredential struct {
	workloadidentity.FederatedCredential
	Reason string
}
type clientFedCredMap map[string][]FederatedCredential

func (target clientFedCredMap) MergeFrom(source clientFedCredMap) {
	for clientId, fedCred := range source {
		if _, ok := target[clientId]; !ok {
			target[clientId] = make([]FederatedCredential, 0, len(fedCred))
		}

		target[clientId] = append(target[clientId], fedCred...)
	}
}

type federatedCredentialsValidationHelper struct {
	apiClient               *radixapiclient.Radixapi
	servicePrincipalService *workloadidentity.AzureServicePrincipalService
	logger                  func(msg string)
}

func (v *federatedCredentialsValidationHelper) ValidateFederatedCredentialsDetails(ctx context.Context, appNames []string) (clientFedCredMap, error) {
	appDeploymentsMap := map[string][]models.Deployment{}

	for _, appName := range appNames {
		v.log(fmt.Sprintf("Reading deployments for application %s", appName))
		appDeployments, err := v.getActiveDeploymentsForApplication(ctx, appName)
		if err != nil {
			return nil, err
		}

		appDeploymentsMap[appName] = appDeployments
	}

	cfg, err := v.apiClient.Configuration.GetConfiguration(&configuration.GetConfigurationParams{Context: ctx}, nil)
	if err != nil {
		return nil, err
	}

	fedCreds := clientFedCredMap{}

	for appName, appDeployments := range appDeploymentsMap {
		for _, deployment := range appDeployments {
			fedCreds.MergeFrom(v.buildFederatedCredentialsMapForDeployment(appName, deployment, cfg.Payload.ClusterOidcIssuers))
		}
	}

	return fedCreds, nil
}

// func (v *federatedCredentialsValidationHelper) buildMissingAndExtraFederatedCredentials(ctx context.Context, fedCredMap clientFedCredMap) (missing, extra []FederatedCredential, err error) {
// 	for clientId, expectedFedCreds := range clientFedCreds {

// 		fmt.Fprintf(os.Stderr, "Analyzing service principal with client id %v: ", clientId)
// 		sp, err := graphHelper.GetServicePrincipal(cmd.Context(), clientId)
// 		if err != nil {
// 			return err
// 		}

// 		fmt.Fprintf(os.Stderr, "%s (%s). ", sp.DisplayName, sp.Type)

// 		existingWorkloadIdentityFedCreds := slice.FindAll(sp.FederatedCredentials, func(fedCred workloadidentity.FederatedCredential) bool {
// 			subjectParts := strings.Split(fedCred.Subject, ":")
// 			if len(subjectParts) != 4 {
// 				return false
// 			}

// 			return subjectParts[0] == "system" && subjectParts[1] == "serviceaccount"
// 		})
// 		mappedExistingWorkloadIdentityFedCreds := slice.Map(existingWorkloadIdentityFedCreds, func(c workloadidentity.FederatedCredential) FederatedCredential {
// 			return FederatedCredential{FederatedCredential: c}
// 		})
// 		missingFedCreds := findMissingFedCreds(expectedFedCreds, mappedExistingWorkloadIdentityFedCreds)
// 		extraFedCreds := findMissingFedCreds(mappedExistingWorkloadIdentityFedCreds, expectedFedCreds)
// 		fmt.Fprintf(os.Stderr, "(total: %v, missing: %v, extra: %v)\n", len(expectedFedCreds), len(missingFedCreds), len(extraFedCreds))

// 		for _, missingFedCred := range missingFedCreds {
// 			command, err := generateCreateFederatedCredentialAzureCLICommand(*sp, missingFedCred)
// 			if err != nil {
// 				return err
// 			}
// 			color.RGB(128, 128, 128).Fprintf(os.Stderr, "# %s\n", missingFedCred.Reason)
// 			color.Green(command)
// 		}

// 		for _, extraFedCred := range extraFedCreds {
// 			command, err := generateDeleteFederatedCredentialAzureCLICommand(*sp, extraFedCred)
// 			if err != nil {
// 				return err
// 			}
// 			color.Red(command)
// 		}
// 	}
// }

func (v *federatedCredentialsValidationHelper) buildFederatedCredentialsMapForDeployment(appName string, deployment models.Deployment, issuers []string) clientFedCredMap {
	fedCreds := clientFedCredMap{}

	for _, component := range deployment.Components {
		reason := fmt.Sprintf("Federated credential for %s %s in application %s, environment %s", strings.ToLower(*component.Type), *component.Name, appName, *deployment.Environment)
		if fedCred := v.buildFederatedCredentialsMapForIdentity(component.Identity, *deployment.Namespace, issuers, reason); fedCred != nil {
			fedCreds.MergeFrom(fedCred)
		}

		if component.Oauth2 != nil {
			reason := fmt.Sprintf("Federated credential for %s %s OAuth2 service in application %s, environment %s", strings.ToLower(*component.Type), *component.Name, appName, *deployment.Environment)
			if fedCred := v.buildFederatedCredentialsMapForIdentity(component.Oauth2.Identity, *deployment.Namespace, issuers, reason); fedCred != nil {
				fedCreds.MergeFrom(fedCred)
			}
		}

		// if horizontalScaling := component.HorizontalScalingSummary; horizontalScaling != nil {
		// 	for _, trigger := range horizontalScaling.Triggers {
		// 		if fedCred := v.buildFederatedCredentialsMapForIdentity(trigger.Identity, *deployment.Namespace); fedCred != nil {
		// 			fedCreds.MergeFrom(fedCred)
		// 		}
		// 	}
		// }
	}

	return fedCreds
}

func (v *federatedCredentialsValidationHelper) buildFederatedCredentialsMapForIdentity(identity *models.Identity, namespace string, issuers []string, reason string) clientFedCredMap {
	if identity == nil || identity.Azure == nil || identity.Azure.ClientID == nil {
		return nil
	}

	clientId := *identity.Azure.ClientID
	fedCreds := clientFedCredMap{clientId: []FederatedCredential{}}

	for _, issuer := range issuers {
		fedCreds[clientId] = append(fedCreds[clientId], FederatedCredential{
			FederatedCredential: workloadidentity.FederatedCredential{
				Issuer:    issuer,
				Subject:   fmt.Sprintf("system:serviceaccount:%s:%s", namespace, *identity.Azure.ServiceAccountName),
				Audiences: []string{"api://AzureADTokenExchange"},
			},
			Reason: reason,
		})
	}

	return fedCreds
}

func (v *federatedCredentialsValidationHelper) getActiveDeploymentsForApplication(ctx context.Context, appName string) ([]models.Deployment, error) {
	app, err := v.apiClient.Application.GetApplication(&application.GetApplicationParams{Context: ctx, AppName: appName}, nil)
	if err != nil {
		return nil, err
	}

	var deployments []models.Deployment

	for _, env := range app.Payload.Environments {
		if env.ActiveDeployment != nil {
			activeDeployment, err := v.apiClient.Deployment.GetDeployment(&deployment.GetDeploymentParams{Context: ctx, AppName: appName, DeploymentName: *env.ActiveDeployment.Name}, nil)
			if err != nil {
				return nil, err
			}
			deployments = append(deployments, *activeDeployment.Payload)
		}
	}

	return deployments, nil
}

func (v *federatedCredentialsValidationHelper) log(msg string) {
	if v.logger != nil {
		v.logger(msg)
	}
}

func generateDeleteFederatedCredentialAzureCLICommand(sp workloadidentity.ServicePrincipal, deleteFedCred FederatedCredential) (string, error) {
	var command string
	switch sp.Type {
	case workloadidentity.ManagedIdentity:
		command = fmt.Sprintf("az identity federated-credential delete --name %s --identity-name %s --resource-group %s --subscription %s", deleteFedCred.Name, sp.DisplayName, sp.ResourceGroup, sp.SubscriptionId)
	case workloadidentity.AppRegistration:
		command = fmt.Sprintf("az ad app federated-credential delete --id %s --federated-credential-id %s", sp.ClientID, deleteFedCred.Name)
	}

	if command == "" {
		return "", fmt.Errorf("unable to generate delete federated credential command for principal %s (%s)", sp.DisplayName, sp.Type)
	}

	return command, nil
}

func generateCreateFederatedCredentialAzureCLICommand(sp workloadidentity.ServicePrincipal, newFedCred FederatedCredential) (string, error) {
	parsedFedCred, err := generateFederatedCredential(newFedCred, sp.FederatedCredentials)
	if err != nil {
		return "", err
	}

	type adAppParam struct {
		Name      string   `json:"name"`
		Issuer    string   `json:"issuer"`
		Subject   string   `json:"subject"`
		Audiences []string `json:"audiences"`
	}

	var command string
	switch sp.Type {
	case workloadidentity.ManagedIdentity:
		command = fmt.Sprintf("az identity federated-credential create --name %s --identity-name %s --resource-group %s --subscription %s --issuer %s --subject %s --audiences %s", parsedFedCred.Name, sp.DisplayName, sp.ResourceGroup, sp.SubscriptionId, parsedFedCred.Issuer, parsedFedCred.Subject, parsedFedCred.Audiences[0])
	case workloadidentity.AppRegistration:

		paramsBytes, err := json.Marshal(adAppParam(*parsedFedCred))
		if err != nil {
			return "", err
		}
		command = fmt.Sprintf("az ad app federated-credential create --id %s --parameters %s", sp.ClientID, fmt.Sprintf("%q", string(paramsBytes)))
	}

	if command == "" {
		return "", fmt.Errorf("unable to generate create federated credential command for principal %s (%s)", sp.DisplayName, sp.Type)
	}

	return command, nil
}

func generateFederatedCredential(federatedCredential FederatedCredential, existingFedCreds []workloadidentity.FederatedCredential) (*workloadidentity.FederatedCredential, error) {
	credentialName, err := validateOrGenerateUniqueFederatedCredentialName(federatedCredential, existingFedCreds)
	if err != nil {
		return nil, err
	}

	return &workloadidentity.FederatedCredential{
		Name:      credentialName,
		Issuer:    federatedCredential.Issuer,
		Subject:   federatedCredential.Subject,
		Audiences: federatedCredential.Audiences,
	}, nil
}

func validateOrGenerateUniqueFederatedCredentialName(fedCred FederatedCredential, existingFedCreds []workloadidentity.FederatedCredential) (string, error) {
	nameUnused := func(name string) bool {
		return !slice.Any(existingFedCreds, federatedCredentialsNameEqualsPredicate(name))
	}

	if existingName := strings.TrimSpace(fedCred.Name); len(existingName) > 0 {
		if !nameUnused(existingName) {
			return "", errors.New("federated credential name already in use")
		}
		return existingName, nil
	}

	for counter := 0; ; counter++ {
		if candidate := generateFederatedCredentialName(fedCred.Subject, fedCred.Issuer, counter); nameUnused(candidate) {
			return candidate, nil
		}
	}
}

func federatedCredentialsNameEqualsPredicate(name string) func(fedCred workloadidentity.FederatedCredential) bool {
	nameLower := strings.ToLower(name)
	return func(fedCred workloadidentity.FederatedCredential) bool {
		return strings.ToLower(fedCred.Name) == nameLower
	}
}

func generateFederatedCredentialName(subject, issuer string, counter int) string {
	const (
		issuerHashLength                 = 10
		federatedCredentialNameMaxLength = 120
	)

	issuerHashFull := sha1.Sum([]byte(strings.TrimSpace(strings.ToLower(issuer))))
	issuerHash := hex.EncodeToString(issuerHashFull[:])[:issuerHashLength]

	subjectPart := sanitizeFederatedCredential(subject)
	if subjectPart == "" {
		subjectPart = "fedcred"
	}

	counterSuffix := ""
	if counter > 0 {
		counterSuffix = fmt.Sprintf("-%d", counter)
	}

	reserved := 1 + len(issuerHash) + len(counterSuffix)
	maxSubjectLength := max(federatedCredentialNameMaxLength-reserved, 1)
	if len(subjectPart) > maxSubjectLength {
		subjectPart = strings.Trim(subjectPart[:maxSubjectLength], "-")
	}

	return subjectPart + "-" + issuerHash + counterSuffix
}

func sanitizeFederatedCredential(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	if value == "" {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(value))
	lastWasDash := false

	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			builder.WriteRune(r)
			lastWasDash = false
			continue
		}

		if !lastWasDash {
			builder.WriteByte('-')
			lastWasDash = true
		}
	}

	return strings.Trim(builder.String(), "-")
}

func findMissingFedCreds(expected, actual []FederatedCredential) (missing []FederatedCredential) {
	for _, expectedFedCred := range expected {
		exists := slices.ContainsFunc(actual, func(actualFedCred FederatedCredential) bool {
			s1 := slices.Clone(expectedFedCred.Audiences)
			s2 := slices.Clone(actualFedCred.Audiences)
			slices.Sort(s1)
			slices.Sort(s2)

			return expectedFedCred.Issuer == actualFedCred.Issuer && expectedFedCred.Subject == actualFedCred.Subject && slices.Equal(s1, s2)
		})
		if !exists {
			missing = append(missing, expectedFedCred)
		}
	}

	return
}

type tokenCredentialAdapter struct {
	auth *auth.Auth
}

func (a *tokenCredentialAdapter) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	if a.auth == nil {
		return azcore.AccessToken{}, fmt.Errorf("auth not set")
	}

	t, err := a.auth.GetAccessToken(ctx, options.Scopes)
	if err != nil {
		return azcore.AccessToken{}, err
	}

	return azcore.AccessToken{Token: t.Token, ExpiresOn: t.ExpiresOn}, nil
}

func init() {
	validateCmd.AddCommand(validateWorkloadIdentityCmd)
	validateWorkloadIdentityCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	_ = validateWorkloadIdentityCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(validateWorkloadIdentityCmd)
}
