package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPITenantClient ...
type IBMPITenantClient struct {
	IBMPIClient
}

// NewIBMPITenantClient ...
func NewIBMPITenantClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPITenantClient {
	return &IBMPITenantClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// Get ..
func (f *IBMPITenantClient) Get(tenantid string) (*models.Tenant, error) {
	params := p_cloud_tenants.NewPcloudTenantsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithTenantID(tenantid)
	resp, err := f.session.Power.PCloudTenants.PcloudTenantsGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant %s with error %w", tenantid, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get tenant %s", tenantid)
	}
	return resp.Payload, nil
}

// Get ..
func (f *IBMPITenantClient) GetSelfTenant() (*models.Tenant, error) {
	params := p_cloud_tenants.NewPcloudTenantsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithTenantID(f.session.UserAccount)
	resp, err := f.session.Power.PCloudTenants.PcloudTenantsGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get self tenant with error %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get self tenant")
	}
	return resp.Payload, nil
}
