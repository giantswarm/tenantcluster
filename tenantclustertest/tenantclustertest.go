package tenantclustertest

import (
	"context"

	"k8s.io/client-go/rest"
)

// Config represents the configuration used to create a new tenant cluster
// service.
type Config struct {
}

// TenantCluster provides functionality for connecting to tenant clusters.
type TenantClusterTest struct {
}

// New creates a new tenant cluster service.
func New(config Config) *TenantClusterTest {
	t := &TenantClusterTest{}
	return t
}

func (t *TenantClusterTest) NewRestConfig(ctx context.Context, clusterID, apiDomain string) (*rest.Config, error) {
	return nil, nil
}
