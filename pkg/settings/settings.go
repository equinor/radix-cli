package settings

import "time"

const (
	DeltaRefreshApplication = 3 * time.Second
	DeltaTimeout            = 30 * time.Second

	FromConfigOption       = "from-config"
	TokenEnvironmentOption = "token-environment"
	TokenStdinOption       = "token-stdin"
	ContextOption          = "context"
	ClusterOption          = "cluster"
	ApiEnvironmentOption   = "api-environment"
	VerboseOption          = "verbose"
)
