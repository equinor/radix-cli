package replicalog

import (
	"context"
	"errors"
	"strconv"
	"time"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/component"
	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/pkg/client/consumer"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/go-openapi/strfmt"
)

type ComponentItem struct {
	Component string
	Replica   string
	created   time.Time
}

func (c ComponentItem) Identifier() string {
	return c.Component + "/" + c.Replica
}
func (c ComponentItem) Created() time.Time {
	return c.created
}

func GetReplicasForComponent(apiClient *radixapi.Radixapi, appName, environmentName, componentName string, previousLog bool) GetReplicasFunc[ComponentItem] {
	return func() ([]ComponentItem, bool, error) {
		environmentParams := environment.NewGetEnvironmentParams()
		environmentParams.SetAppName(appName)
		environmentParams.SetEnvName(environmentName)
		environmentDetails, err := apiClient.Environment.GetEnvironment(environmentParams, nil)

		if err != nil {
			return nil, false, err
		}

		if environmentDetails == nil || environmentDetails.Payload.ActiveDeployment == nil {
			return nil, false, errors.New("active deployment was not found in environment")
		}

		var replicas []ComponentItem
		for _, comp := range environmentDetails.Payload.ActiveDeployment.Components {
			if comp.Name == nil || *comp.Name != componentName {
				continue
			}

			for _, replica := range comp.ReplicaList {
				replicas = append(replicas, ComponentItem{
					Component: *comp.Name,
					Replica:   *replica.Name,
					created:   time.Time(*replica.Created),
				})
			}
		}

		// if previous log is requested, return replicas only once
		return replicas, previousLog, err
	}
}

// getComponentReplicasForEnvironment returns all replicas for all components in an environment, if previousLog is true, only return replicas once
func GetComponentReplicasForEnvironment(apiClient *radixapi.Radixapi, appName, environmentName string, previousLog bool) GetReplicasFunc[ComponentItem] {
	return func() ([]ComponentItem, bool, error) {
		// Get active deployment
		environmentParams := environment.NewGetEnvironmentParams()
		environmentParams.SetAppName(appName)
		environmentParams.SetEnvName(environmentName)
		environmentDetails, err := apiClient.Environment.GetEnvironment(environmentParams, nil)

		if err != nil {
			return nil, false, err
		}

		if environmentDetails == nil || environmentDetails.Payload.ActiveDeployment == nil {
			return nil, false, errors.New("active deployment was not found in environment")
		}

		componentReplicas := make([]ComponentItem, 0, 50)
		for _, component := range environmentDetails.Payload.ActiveDeployment.Components {
			if component.Name != nil {
				for _, replica := range component.ReplicaList {
					componentReplicas = append(componentReplicas, ComponentItem{
						Component: *component.Name,
						Replica:   *replica.Name,
					})
				}
				//
			}
		}

		// If previous log is requested, return replicas only once
		return componentReplicas, previousLog, nil
	}
}

type OAuth2ComponentItem struct {
	Component string
	Env       string
	Replica   string
	AuxType   string
	created   time.Time
}

func (c OAuth2ComponentItem) Identifier() string {
	return c.Component + "/" + c.AuxType + "/" + c.Replica
}

func (c OAuth2ComponentItem) Created() time.Time {
	return c.created
}

func GetReplicasForComponentOAuth2(apiClient *radixapi.Radixapi, appName, environmentName, componentName, auxType string, previousLog bool) GetReplicasFunc[OAuth2ComponentItem] {
	return func() ([]OAuth2ComponentItem, bool, error) {
		environmentParams := environment.NewGetEnvironmentParams()
		environmentParams.SetAppName(appName)
		environmentParams.SetEnvName(environmentName)
		environmentDetails, err := apiClient.Environment.GetEnvironment(environmentParams, nil)

		if err != nil {
			return nil, false, err
		}

		if environmentDetails == nil || environmentDetails.Payload.ActiveDeployment == nil {
			return nil, false, errors.New("active deployment was not found in environment")
		}

		var replicas []OAuth2ComponentItem
		for _, comp := range environmentDetails.Payload.ActiveDeployment.Components {
			if comp.Name == nil || *comp.Name != componentName {
				continue
			}
			if comp.Oauth2 != nil {
				for _, deployment := range comp.Oauth2.Deployments {
					if deployment.Type != auxType {
						continue
					}
					for _, replica := range deployment.ReplicaList {
						replicas = append(replicas, OAuth2ComponentItem{
							Component: *comp.Name,
							Env:       environmentName,
							Replica:   *replica.Name,
							AuxType:   deployment.Type,
							created:   time.Time(*replica.Created),
						})
					}
				}
			}
			break
		}

		return replicas, previousLog, nil
	}
}

func GetOAuth2ComponentLog(apiClient *radixapi.Radixapi, appName string, previous bool) GetLogFunc[OAuth2ComponentItem] {
	previousStr := strconv.FormatBool(previous)

	return func(ctx context.Context, item OAuth2ComponentItem, since time.Time, callback consumer.EventCallback) error {
		logParameters := component.NewGetOAuthPodLogParamsWithContext(ctx)
		logParameters.WithAppName(appName)
		logParameters.WithEnvName(item.Env)
		logParameters.WithComponentName(item.Component)
		logParameters.WithPodName(item.Replica)
		logParameters.WithType(item.AuxType)
		logParameters.WithFollow(pointers.Ptr("true"))
		logParameters.WithSinceTime(pointers.Ptr(strfmt.DateTime(since)))
		logParameters.WithPrevious(&previousStr)

		_, err := apiClient.Component.GetOAuthPodLog(logParameters, nil, consumer.NewEventSourceClientOptions(callback))
		return err
	}
}

func GetComponentLog(apiClient *radixapi.Radixapi, appName string, previous bool) GetLogFunc[ComponentItem] {
	previousStr := strconv.FormatBool(previous)

	return func(ctx context.Context, item ComponentItem, since time.Time, callback consumer.EventCallback) error {
		logParameters := component.NewLogParamsWithContext(ctx)
		logParameters.WithAppName(appName)
		logParameters.WithDeploymentName("irrelevant")
		logParameters.WithComponentName(item.Component)
		logParameters.WithPodName(item.Replica)
		logParameters.WithFollow(pointers.Ptr("true"))
		logParameters.SetSinceTime(pointers.Ptr(strfmt.DateTime(since)))
		logParameters.WithPrevious(&previousStr)

		_, err := apiClient.Component.Log(logParameters, nil, consumer.NewEventSourceClientOptions(callback))
		return err
	}
}
