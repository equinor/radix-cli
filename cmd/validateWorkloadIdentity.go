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

	radixapiclient "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/client/configuration"
	"github.com/equinor/radix-cli/generated/radixapi/client/deployment"
	"github.com/equinor/radix-cli/generated/radixapi/client/platform"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/auth"
	"github.com/equinor/radix-cli/pkg/auth/adapter"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/flagvalues"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/workloadidentity"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/fatih/color"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/spf13/cobra"
)

const (
	flagExcludeObsoleteFederatedCredentials = "exclude-obsolete"
	flagAzureCLICommandFormat               = "azure-cli-command-format"

	azureCLICommandFormatPosix             = "posix"
	azureCLICommandFormatWindowsCmd        = "windows-cmd"
	azureCLICommandFormatWindowsPowershell = "windows-powershell"
)

var (
	validAzureCliCommandFormats = []string{azureCLICommandFormatPosix, azureCLICommandFormatWindowsCmd, azureCLICommandFormatWindowsPowershell}
)

var validateWorkloadIdentityCmd = &cobra.Command{
	Use:   "workload-identity",
	Short: "Validate radixconfig.yaml",
	Long: `Validate workload identity configuration for one application or all applications.

The command compares expected federated credentials with existing credentials in Azure and prints Azure CLI commands to create missing credentials and delete potentially obsolete credentials.

Take care when reviewing obsolete federated credentials: the obsolete list is best-effort and must not be trusted 100%. Existing federated credentials can belong to another Radix cluster, even if they look obsolete for the currently selected context.

Warning: Federated credential subjects used for horizontal scaling triggers (subject system:serviceaccount:keda:keda-operator) can be shared across applications for the same service principal. A credential that appears obsolete for one application can still be required by another application in the same Radix cluster.

Note: After running generated Azure CLI commands, Azure can take up to one minute to process federated credential changes. If you run validation immediately after, the result can be stale.`,
	Example: `  # Validate workload identity for all applications in current context
  rx validate workload-identity

  # Validate workload identity for one application
  rx validate workload-identity --application my-app

  # Print result as JSON
  rx validate workload-identity --application my-app --output json

  # Exclude potentially obsolete federated credentials from output
  rx validate workload-identity --application my-app --exclude-obsolete

  # Generate Azure CLI commands formatted for Powershell on Windows
  rx validate workload-identity --application my-app --azure-cli-command-format windows-powershell`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		outputFormat, _ := cmd.Flags().GetString(flagnames.Output)
		azureCLICommandFormat, _ := cmd.Flags().GetString(flagAzureCLICommandFormat)
		if !slices.Contains(validAzureCliCommandFormats, azureCLICommandFormat) {
			return fmt.Errorf("invalid value for --%s: %q", flagAzureCLICommandFormat, azureCLICommandFormat)
		}

		validationPrinter := func(validations []workloadIdentityValidation) error {
			return validationTextPrinter(validations, azureCLICommandFormat)
		}
		if outputFormat == flagvalues.OutputFormatJson {
			validationPrinter = validationJsonPrinter
		}

		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

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

		radixAuth, err := auth.New()
		if err != nil {
			return err
		}

		servicePrincipalHelper, err := workloadidentity.NewAzureServicePrincipalHelper(adapter.NewAzureTokenCredentialAdapter(radixAuth))
		if err != nil {
			return fmt.Errorf("error initializing service principal client: %w", err)
		}

		validator := workloadIdentityValidationHelper{
			radixApiClient:         apiClient,
			servicePrincipalHelper: servicePrincipalHelper,
			logger:                 func(msg string) { fmt.Fprintln(os.Stderr, msg) },
		}

		validationResult, err := validator.ValidateWorkloadIdentities(cmd.Context(), appNames)
		if err != nil {
			return err
		}

		if excludeObsolete, _ := cmd.Flags().GetBool(flagExcludeObsoleteFederatedCredentials); excludeObsolete {
			for i := range validationResult {
				validationResult[i].ObsoleteFederatedCredentials = nil
			}
		}

		if err := validationPrinter(validationResult); err != nil {
			return err
		}

		return nil
	},
}

func validationTextPrinter(validations []workloadIdentityValidation, azureCLICommandFormat string) error {
	if len(validations) == 0 {
		fmt.Fprintln(os.Stdout, "No workload identity configurations found for selected application(s).")
		return nil
	}

	var lineBreak bool
	commentPrinter := color.RGB(128, 128, 128)

	for _, validation := range validations {
		if len(validation.MissingFederatedCredentials) > 0 {
			if lineBreak {
				fmt.Fprintln(os.Stdout)
			}

			commentPrinter.Fprintf(os.Stdout, "# Azure CLI commands to create missing federated credentials for %s %s (client ID: %s)\n", validation.ServicePrincipal.Type, validation.ServicePrincipal.DisplayName, validation.ServicePrincipal.ClientID)
		}

		for _, missing := range validation.MissingFederatedCredentials {
			command, err := generateCreateFederatedCredentialAzureCLICommand(validation.ServicePrincipal, missing, azureCLICommandFormat)
			if err != nil {
				return err
			}

			for _, reason := range missing.Reason {
				commentPrinter.Fprintf(os.Stdout, "# %s\n", reason)
			}

			color.Green(command)
			lineBreak = true
		}

	}

	for _, validation := range validations {
		if len(validation.ObsoleteFederatedCredentials) > 0 {
			if lineBreak {
				fmt.Fprintln(os.Stdout)
			}

			commentPrinter.Fprintf(os.Stdout, "# Azure CLI commands to delete obsolete federated credentials for %s %s (client ID: %s)\n", validation.ServicePrincipal.Type, validation.ServicePrincipal.DisplayName, validation.ServicePrincipal.ClientID)
		}

		for _, obsolete := range validation.ObsoleteFederatedCredentials {
			command, err := generateDeleteFederatedCredentialAzureCLICommand(validation.ServicePrincipal, obsolete)
			if err != nil {
				return err
			}

			for _, reason := range obsolete.Reason {
				commentPrinter.Fprintf(os.Stdout, "# %s\n", reason)
			}

			color.Red(command)
			lineBreak = true
		}
	}

	return nil
}

func validationJsonPrinter(validations []workloadIdentityValidation) error {
	if len(validations) == 0 {
		fmt.Fprintln(os.Stdout, "No workload identity configurations found for selected application(s).")
		return nil
	}

	printPayload(validations)
	return nil
}

func generateDeleteFederatedCredentialAzureCLICommand(sp workloadidentity.ServicePrincipal, deleteFedCred federatedCredential) (string, error) {
	switch sp.Type {
	case workloadidentity.ManagedIdentity:
		return fmt.Sprintf("az identity federated-credential delete --name %s --identity-name %s --resource-group %s --subscription %s", deleteFedCred.Name, sp.DisplayName, sp.ResourceGroup, sp.SubscriptionID), nil
	case workloadidentity.AppRegistration:
		return fmt.Sprintf("az ad app federated-credential delete --id %s --federated-credential-id %s", sp.ClientID, deleteFedCred.Name), nil
	}

	return "", fmt.Errorf("unable to generate delete federated credential command for principal %s with unknow type %s", sp.DisplayName, sp.Type)
}

func generateCreateFederatedCredentialAzureCLICommand(sp workloadidentity.ServicePrincipal, createFedCred federatedCredential, azureCLICommandFormat string) (string, error) {
	switch sp.Type {
	case workloadidentity.ManagedIdentity:
		return fmt.Sprintf("az identity federated-credential create --name %s --identity-name %s --resource-group %s --subscription %s --issuer %s --subject %s --audiences %s", createFedCred.Name, sp.DisplayName, sp.ResourceGroup, sp.SubscriptionID, createFedCred.Issuer, createFedCred.Subject, createFedCred.Audiences[0]), nil
	case workloadidentity.AppRegistration:
		type addParam struct {
			Name      string   `json:"name"`
			Issuer    string   `json:"issuer"`
			Subject   string   `json:"subject"`
			Audiences []string `json:"audiences"`
		}
		param := addParam{
			Name:      createFedCred.Name,
			Issuer:    createFedCred.Issuer,
			Subject:   createFedCred.Subject,
			Audiences: createFedCred.Audiences,
		}
		paramBytes, err := json.Marshal(param)
		if err != nil {
			return "", err
		}

		formattedParams, err := formatAzureCLIJSONParameters(string(paramBytes), azureCLICommandFormat)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("az ad app federated-credential create --id %s --parameters %s", sp.ClientID, formattedParams), nil
	}

	return "", fmt.Errorf("unable to generate create federated credential command for principal %s with unknow type %s", sp.DisplayName, sp.Type)
}

func formatAzureCLIJSONParameters(parametersJSON string, azureCLICommandFormat string) (string, error) {
	switch azureCLICommandFormat {
	case azureCLICommandFormatPosix:
		escapedParams := strings.ReplaceAll(parametersJSON, `'`, `'\''`)
		return fmt.Sprintf("'%s'", escapedParams), nil
	case azureCLICommandFormatWindowsCmd:
		escapedParams := strings.ReplaceAll(parametersJSON, `"`, `\"`)
		return fmt.Sprintf("\"%s\"", escapedParams), nil
	case azureCLICommandFormatWindowsPowershell:
		escapedParams := strings.ReplaceAll(parametersJSON, `"`, `"""`)
		escapedParams = strings.ReplaceAll(escapedParams, `'`, `''`)
		return fmt.Sprintf("'%s'", escapedParams), nil
	default:
		return "", fmt.Errorf("unsupported Azure CLI command format %q", azureCLICommandFormat)
	}
}

type federatedCredential struct {
	workloadidentity.FederatedCredential
	Reason []string `json:"reason"`
}

type workloadIdentityValidation struct {
	ServicePrincipal             workloadidentity.ServicePrincipal `json:"servicePrincipal"`
	MissingFederatedCredentials  []federatedCredential             `json:"missingFederatedCredentials"`
	ObsoleteFederatedCredentials []federatedCredential             `json:"obsoleteFederatedCredentials,omitempty"`
}

type clientFedCredMap map[string][]federatedCredential

func (target clientFedCredMap) MergeFrom(source clientFedCredMap) {
	for clientId, fedCred := range source {
		if _, ok := target[clientId]; !ok {
			target[clientId] = make([]federatedCredential, 0, len(fedCred))
		}

		target[clientId] = append(target[clientId], fedCred...)
	}
}

type workloadIdentityValidationHelper struct {
	radixApiClient         *radixapiclient.Radixapi
	servicePrincipalHelper *workloadidentity.AzureServicePrincipalHelper
	logger                 func(msg string)
}

func (v *workloadIdentityValidationHelper) ValidateWorkloadIdentities(ctx context.Context, appNames []string) ([]workloadIdentityValidation, error) {
	appDeploymentsMap := map[string][]models.Deployment{}
	affectedNamespaces := []string{"keda"}

	for _, appName := range appNames {
		v.logMessage(fmt.Sprintf("Reading deployments for application %s", appName))
		appDeployments, err := v.getActiveDeploymentsForApplication(ctx, appName)
		if err != nil {
			return nil, err
		}

		appDeploymentsMap[appName] = appDeployments
		affectedNamespaces = append(affectedNamespaces, slice.Map(appDeployments, func(d models.Deployment) string { return *d.Namespace })...)
	}

	cfg, err := v.radixApiClient.Configuration.GetConfiguration(&configuration.GetConfigurationParams{Context: ctx}, nil)
	if err != nil {
		return nil, err
	}

	fedCreds := clientFedCredMap{}
	for appName, appDeployments := range appDeploymentsMap {
		for _, deployment := range appDeployments {
			fedCreds.MergeFrom(v.buildFederatedCredentialsMapForDeployment(appName, deployment, cfg.Payload.ClusterOidcIssuers))
		}
	}

	return v.buildFederatedCredentialValidation(ctx, fedCreds, affectedNamespaces)
}

func (v *workloadIdentityValidationHelper) buildFederatedCredentialValidation(ctx context.Context, fedCredMap clientFedCredMap, affectedNamespaces []string) ([]workloadIdentityValidation, error) {
	validations := []workloadIdentityValidation{}

	v.logMessage("")

	for clientId, expectedFedCreds := range fedCredMap {
		expectedFedCreds = compactFederatedCredentials(expectedFedCreds)

		v.logMessage(fmt.Sprintf("Analyzing service principal with client id %s", clientId))

		sp, err := v.servicePrincipalHelper.GetServicePrincipal(ctx, clientId)
		if err != nil {
			return nil, err
		}
		v.logMessage(fmt.Sprintf("Found %s with name %s ", sp.Type, sp.DisplayName))

		missingFedCreds, obsoleteFedCreds := compareServicePrincipalFederatedCredentials(sp, expectedFedCreds, affectedNamespaces)

		for fedCredIndex := range missingFedCreds {
			fedCredName, err := validateOrGenerateUniqueFederatedCredentialName(missingFedCreds[fedCredIndex], sp.FederatedCredentials)
			if err != nil {
				return nil, err
			}
			missingFedCreds[fedCredIndex].Name = fedCredName
		}

		printColor := color.FgGreen
		if len(missingFedCreds) > 0 || len(obsoleteFedCreds) > 0 {
			printColor = color.FgYellow
		}
		v.logMessage(color.Set(printColor).Sprintf("Federated credentials expected: %v, missing: %v, obsolete: %v\n", len(expectedFedCreds), len(missingFedCreds), len(obsoleteFedCreds)))

		validations = append(validations, workloadIdentityValidation{
			ServicePrincipal:             *sp,
			MissingFederatedCredentials:  missingFedCreds,
			ObsoleteFederatedCredentials: obsoleteFedCreds,
		})
	}

	return validations, nil
}

func compareServicePrincipalFederatedCredentials(sp *workloadidentity.ServicePrincipal, expectedFedCreds []federatedCredential, affectedNamespaces []string) (missing, obsolete []federatedCredential) {
	expectedFedCreds = compactFederatedCredentials(expectedFedCreds)

	existingFedCreds := slice.Map(sp.FederatedCredentials, func(c workloadidentity.FederatedCredential) federatedCredential {
		return federatedCredential{FederatedCredential: c}
	})

	// Find feddcreds that exists in expected fedcreds but not in existing fedcreds
	missingFedCreds := findMissingFedCreds(expectedFedCreds, existingFedCreds)

	// Find fedcreds that exists in existing fedcreds but not in expected fedcreds
	obsoleteFedCreds := findMissingFedCreds(existingFedCreds, expectedFedCreds)

	// Filter obsolete fedcreds and only include fedcreds which namespace exists in the list of affected namespaces
	// We do not care about obsolete fedcreds that does not belong to other apps
	filteredObsoleteFedCreds := slice.FindAll(obsoleteFedCreds, func(fc federatedCredential) bool {
		namespace, _, ok := classifyKubernetesFederatedCredential(fc.FederatedCredential)
		return ok && slices.Contains(affectedNamespaces, namespace)
	})

	return missingFedCreds, filteredObsoleteFedCreds
}

func compactFederatedCredentials(fedCreds []federatedCredential) []federatedCredential {
	var compactFedCred []federatedCredential

	for _, fedCred := range fedCreds {
		predicate := createFederatedCredentialEqualsPredicate(fedCred)
		if i := slices.IndexFunc(compactFedCred, predicate); i == -1 {
			compactFedCred = append(compactFedCred, fedCred)
		} else {
			compactFedCred[i].Reason = append(compactFedCred[i].Reason, fedCred.Reason...)
		}
	}

	return compactFedCred
}

func (v *workloadIdentityValidationHelper) buildFederatedCredentialsMapForDeployment(appName string, deployment models.Deployment, issuers []string) clientFedCredMap {
	fedCreds := clientFedCredMap{}

	for _, component := range deployment.Components {
		reason := fmt.Sprintf("Federated credential for %s %s in application %s, environment %s", strings.ToLower(*component.Type), *component.Name, appName, *deployment.Environment)
		if fedCred := v.buildFederatedCredentialsMapForIdentity(component.Identity, issuers, reason); fedCred != nil {
			fedCreds.MergeFrom(fedCred)
		}

		if component.Oauth2 != nil {
			reason := fmt.Sprintf("Federated credential for %s %s OAuth2 service in application %s, environment %s", strings.ToLower(*component.Type), *component.Name, appName, *deployment.Environment)
			if fedCred := v.buildFederatedCredentialsMapForIdentity(component.Oauth2.Identity, issuers, reason); fedCred != nil {
				fedCreds.MergeFrom(fedCred)
			}
		}

		if horizontalScaling := component.HorizontalScalingSummary; horizontalScaling != nil {
			for _, trigger := range horizontalScaling.Triggers {
				reason := fmt.Sprintf("Federated credential for %s %s horizontal scaling %s in application %s, environment %s", strings.ToLower(*component.Type), *component.Name, trigger.Name, appName, *deployment.Environment)
				if fedCred := v.buildFederatedCredentialsMapForIdentity(trigger.Identity, issuers, reason); fedCred != nil {
					fedCreds.MergeFrom(fedCred)
				}
			}
		}
	}

	return fedCreds
}

func (v *workloadIdentityValidationHelper) buildFederatedCredentialsMapForIdentity(identity *models.Identity, issuers []string, reason string) clientFedCredMap {
	if identity == nil || identity.Azure == nil || identity.Azure.ClientID == nil {
		return nil
	}

	clientId := *identity.Azure.ClientID
	fedCreds := clientFedCredMap{clientId: []federatedCredential{}}

	for _, issuer := range issuers {
		fedCred := federatedCredential{
			FederatedCredential: workloadidentity.FederatedCredential{
				Issuer:    issuer,
				Subject:   fmt.Sprintf("system:serviceaccount:%s:%s", *identity.Azure.Namespace, *identity.Azure.ServiceAccountName),
				Audiences: []string{"api://AzureADTokenExchange"},
			},
		}
		if len(reason) > 0 {
			fedCred.Reason = append(fedCred.Reason, reason)
		}
		fedCreds[clientId] = append(fedCreds[clientId], fedCred)
	}

	return fedCreds
}

func (v *workloadIdentityValidationHelper) getActiveDeploymentsForApplication(ctx context.Context, appName string) ([]models.Deployment, error) {
	app, err := v.radixApiClient.Application.GetApplication(&application.GetApplicationParams{Context: ctx, AppName: appName}, nil)
	if err != nil {
		return nil, err
	}

	var deployments []models.Deployment

	for _, env := range app.Payload.Environments {
		if env.ActiveDeployment != nil {
			activeDeployment, err := v.radixApiClient.Deployment.GetDeployment(&deployment.GetDeploymentParams{Context: ctx, AppName: appName, DeploymentName: *env.ActiveDeployment.Name}, nil)
			if err != nil {
				return nil, err
			}
			deployments = append(deployments, *activeDeployment.Payload)
		}
	}

	return deployments, nil
}

func (v *workloadIdentityValidationHelper) logMessage(msg string) {
	if v.logger != nil {
		v.logger(msg)
	}
}

func classifyKubernetesFederatedCredential(fedCred workloadidentity.FederatedCredential) (namespace, serviceAccount string, ok bool) {
	subjectParts := strings.Split(fedCred.Subject, ":")
	if len(subjectParts) != 4 || subjectParts[0] != "system" || subjectParts[1] != "serviceaccount" {
		return "", "", false
	}

	return subjectParts[2], subjectParts[3], true
}

func validateOrGenerateUniqueFederatedCredentialName(fedCred federatedCredential, existingFedCreds []workloadidentity.FederatedCredential) (string, error) {
	nameUnused := func(name string) bool {
		return !slice.Any(existingFedCreds, createFederatedCredentialsNameEqualsPredicate(name))
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

// returns a list of federatedCredential that exists in expected but not in actual
func findMissingFedCreds(expected, actual []federatedCredential) []federatedCredential {
	var missing []federatedCredential

	for _, expectedFedCred := range expected {
		predicate := createFederatedCredentialEqualsPredicate(expectedFedCred)

		if exists := slices.ContainsFunc(actual, predicate); !exists {
			missing = append(missing, expectedFedCred)
		}
	}

	return missing
}

func createFederatedCredentialsNameEqualsPredicate(name string) func(fedCred workloadidentity.FederatedCredential) bool {
	nameLower := strings.ToLower(name)
	return func(fedCred workloadidentity.FederatedCredential) bool {
		return strings.ToLower(fedCred.Name) == nameLower
	}
}

func createFederatedCredentialEqualsPredicate(fedCred federatedCredential) func(federatedCredential) bool {
	return func(compareWith federatedCredential) bool {
		fedCredAudSorted := slices.Clone(fedCred.Audiences)
		compareWithAudSorted := slices.Clone(compareWith.Audiences)
		slices.Sort(fedCredAudSorted)
		slices.Sort(compareWithAudSorted)

		return fedCred.Issuer == compareWith.Issuer && fedCred.Subject == compareWith.Subject && slices.Equal(fedCredAudSorted, compareWithAudSorted)
	}
}

func init() {
	validateCmd.AddCommand(validateWorkloadIdentityCmd)
	validateWorkloadIdentityCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	validateWorkloadIdentityCmd.Flags().StringP(flagnames.Output, "o", "text", "(Optional) Output format. Valud options are json or text")
	validateWorkloadIdentityCmd.Flags().Bool(flagExcludeObsoleteFederatedCredentials, false, "Exclude potential obsolete federated credentials from output")
	validateWorkloadIdentityCmd.Flags().String(flagAzureCLICommandFormat, azureCLICommandFormatDefault, fmt.Sprintf("Azure CLI command format. Valid options: %s", strings.Join(validAzureCliCommandFormats, ", ")))

	_ = validateWorkloadIdentityCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = validateWorkloadIdentityCmd.RegisterFlagCompletionFunc(flagnames.Output, completion.Output)
	_ = validateWorkloadIdentityCmd.RegisterFlagCompletionFunc(flagAzureCLICommandFormat, azureCLICommandFormatCompletions)
	setContextSpecificPersistentFlags(validateWorkloadIdentityCmd)
}

func azureCLICommandFormatCompletions(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	formats := slice.FindAll(validAzureCliCommandFormats, func(format string) bool {
		return strings.HasPrefix(format, toComplete)
	})
	return formats, cobra.ShellCompDirectiveNoFileComp
}
