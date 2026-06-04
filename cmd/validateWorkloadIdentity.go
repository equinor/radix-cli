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
	"slices"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/client/configuration"
	"github.com/equinor/radix-cli/generated/radixapi/client/deployment"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/auth"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/workloadidentity"
	"github.com/equinor/radix-common/utils/slice"
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
		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		cfg, err := apiClient.Configuration.GetConfiguration(&configuration.GetConfigurationParams{Context: cmd.Context()}, nil)
		if err != nil {
			return err
		}
		issuers := cfg.Payload.ClusterOidcIssuers

		app, err := apiClient.Application.GetApplication(&application.GetApplicationParams{AppName: "oauth-demo", Context: cmd.Context()}, nil)
		if err != nil {
			return err
		}

		type appEnv struct {
			app, env string
		}

		deployments := map[appEnv]string{}
		for _, env := range app.Payload.Environments {
			if deployment := env.ActiveDeployment; deployment != nil {
				deployments[appEnv{app: *app.Payload.Name, env: *env.Name}] = *deployment.Name
			}

		}
		clientFedCreds := map[string][]workloadidentity.FederatedCredential{}

		addClientFedCred := func(identity *models.Identity, namespace string, issuers []string) {
			if identity == nil || identity.Azure == nil || identity.Azure.ClientID == nil {
				return
			}

			clientId := *identity.Azure.ClientID

			if _, ok := clientFedCreds[clientId]; !ok {
				clientFedCreds[clientId] = []workloadidentity.FederatedCredential{}
			}

			for _, issuer := range issuers {
				clientFedCreds[clientId] = append(clientFedCreds[clientId], workloadidentity.FederatedCredential{
					Issuer:    issuer,
					Subject:   fmt.Sprintf("system:serviceaccount:%s:%s", namespace, *identity.Azure.ServiceAccountName),
					Audiences: []string{"api://AzureADTokenExchange"},
				})
			}

		}

		for appEnvInfo, deploymentName := range deployments {
			deployment, err := apiClient.Deployment.GetDeployment(&deployment.GetDeploymentParams{AppName: appEnvInfo.app, DeploymentName: deploymentName, Context: cmd.Context()}, nil)
			if err != nil {
				return err
			}

			for _, component := range deployment.Payload.Components {
				addClientFedCred(component.Identity, *deployment.Payload.Namespace, issuers)

				if component.Oauth2 != nil {
					addClientFedCred(component.Oauth2.Identity, *deployment.Payload.Namespace, issuers)
				}
			}

		}

		radixAuth, err := auth.New()
		if err != nil {
			return err
		}
		graphHelper, err := workloadidentity.NewAzureServicePrincipalService(&tokenCredentialAdapter{auth: radixAuth})
		if err != nil {
			return fmt.Errorf("Error initializing client: %w", err)
		}

		var createCommands, deleteCommands []string

		for clientId, expectedFedCreds := range clientFedCreds {
			fmt.Printf("Analysing service principal with client id %v: ", clientId)
			sp, err := graphHelper.GetServicePrincipal(cmd.Context(), clientId)
			if err != nil {
				return err
			}
			existingWorkloadIdentityFedCreds := slice.FindAll(sp.FederatedCredentials, func(fedCred workloadidentity.FederatedCredential) bool {
				subjectParts := strings.Split(fedCred.Subject, ":")
				if len(subjectParts) != 4 {
					return false
				}

				return subjectParts[0] == "system" && subjectParts[1] == "serviceaccount"
			})
			fmt.Printf("%s (%s). ", sp.DisplayName, sp.Type)
			missingFedCreds := findMissingFedCreds(expectedFedCreds, existingWorkloadIdentityFedCreds)
			extraFedCreds := findMissingFedCreds(existingWorkloadIdentityFedCreds, expectedFedCreds)
			fmt.Printf("(total: %v, missing: %v, extra: %v)\n", len(expectedFedCreds), len(missingFedCreds), len(extraFedCreds))

			for _, missingFedCred := range missingFedCreds {
				command, err := generateCreateFederatedCredentialAzureCLICommand(*sp, missingFedCred)
				if err != nil {
					return err
				}
				createCommands = append(createCommands, command)
			}

			for _, extraFedCred := range extraFedCreds {
				command, err := generateDeleteFederatedCredentialAzureCLICommand(*sp, extraFedCred)
				if err != nil {
					return err
				}
				deleteCommands = append(deleteCommands, command)
			}
		}

		fmt.Println()
		fmt.Println(strings.Join(createCommands, "\n"))
		fmt.Println()
		fmt.Println(strings.Join(deleteCommands, "\n"))

		return nil
	},
}

func generateDeleteFederatedCredentialAzureCLICommand(sp workloadidentity.ServicePrincipal, deleteFedCred workloadidentity.FederatedCredential) (string, error) {
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

func generateCreateFederatedCredentialAzureCLICommand(sp workloadidentity.ServicePrincipal, newFedCred workloadidentity.FederatedCredential) (string, error) {
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

func generateFederatedCredential(federatedCredential workloadidentity.FederatedCredential, existingFedCreds []workloadidentity.FederatedCredential) (*workloadidentity.FederatedCredential, error) {
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

func validateOrGenerateUniqueFederatedCredentialName(fedCred workloadidentity.FederatedCredential, existingFedCreds []workloadidentity.FederatedCredential) (string, error) {
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

func findMissingFedCreds(expected, actual []workloadidentity.FederatedCredential) (missing []workloadidentity.FederatedCredential) {
	for _, expectedFedCred := range expected {
		exists := slices.ContainsFunc(actual, func(actualFedCred workloadidentity.FederatedCredential) bool {
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
