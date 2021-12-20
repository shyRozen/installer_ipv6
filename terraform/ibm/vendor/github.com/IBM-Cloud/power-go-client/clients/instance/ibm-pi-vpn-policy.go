package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_v_p_n_policies"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIVpnPolicyClient ...
type IBMPIVpnPolicyClient struct {
	IBMPIClient
}

// IBMPIVpnPolicyClient ...
func NewIBMPIVpnPolicyClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIVpnPolicyClient {
	return &IBMPIVpnPolicyClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

// IKE Policies

// Get information about a IKE Policy
func (f *IBMPIVpnPolicyClient) GetIKEPolicy(id string) (*models.IKEPolicy, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithIkePolicyID(id)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNPolicyOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get ike policy for policy id %s", id)
	}
	return resp.Payload, nil
}

// Create a IKE Policy
func (f *IBMPIVpnPolicyClient) CreateIKEPolicy(body *models.IKEPolicyCreate) (*models.IKEPolicy, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNPolicyOperationFailed, f.cloudInstanceID, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Policy")
}

// Update a IKE Policy
func (f *IBMPIVpnPolicyClient) UpdateIKEPolicy(id string, body *models.IKEPolicyUpdate) (*models.IKEPolicy, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithIkePolicyID(id).
		WithBody(body)
	putok, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesPut(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNPolicyOperationFailed, id, err)
	}
	if putok != nil && putok.Payload != nil {
		return putok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Policy")
}

// Get all IKE Policies
func (f *IBMPIVpnPolicyClient) GetAllIKEPolicies() (*models.IKEPolicies, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ike policies: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all ike policies")
	}
	return resp.Payload, nil
}

// Delete a IKE Policy
func (f *IBMPIVpnPolicyClient) DeleteIKEPolicy(id string) error {
	params := p_cloud_v_p_n_policies.NewPcloudIkepoliciesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithIkePolicyID(id)
	_, err := f.session.Power.PCloudVPNPolicies.PcloudIkepoliciesDelete(params, f.authInfo)
	if err != nil {
		return fmt.Errorf(errors.DeleteVPNPolicyOperationFailed, id, err)
	}
	return nil
}

// IPSec Policies

// Get information about a IPSec Policy
func (f *IBMPIVpnPolicyClient) GetIPSecPolicy(id string) (*models.IPSecPolicy, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithIpsecPolicyID(id)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.GetVPNPolicyOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get ipsec policy for policy id %s", id)
	}
	return resp.Payload, nil
}

// Create a IPSec Policy
func (f *IBMPIVpnPolicyClient) CreateIPSecPolicy(body *models.IPSecPolicyCreate) (*models.IPSecPolicy, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVPNPolicyOperationFailed, f.cloudInstanceID, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create VPN Policy")
}

// Update a IPSec Policy
func (f *IBMPIVpnPolicyClient) UpdateIPSecPolicy(id string, body *models.IPSecPolicyUpdate) (*models.IPSecPolicy, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithIpsecPolicyID(id).
		WithBody(body)
	putok, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesPut(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVPNPolicyOperationFailed, id, err)
	}
	if putok != nil && putok.Payload != nil {
		return putok.Payload, nil
	}
	return nil, fmt.Errorf("failed to Update VPN Policy")
}

// Get all IPSec Policies
func (f *IBMPIVpnPolicyClient) GetAllIPSecPolicies() (*models.IPSecPolicies, error) {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ipsec policies: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all ipsec policies")
	}
	return resp.Payload, nil
}

// Delete a IPSec Policy
func (f *IBMPIVpnPolicyClient) DeleteIPSecPolicy(id string) error {
	params := p_cloud_v_p_n_policies.NewPcloudIpsecpoliciesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithIpsecPolicyID(id)
	_, err := f.session.Power.PCloudVPNPolicies.PcloudIpsecpoliciesDelete(params, f.authInfo)
	if err != nil {
		return fmt.Errorf(errors.DeleteVPNPolicyOperationFailed, id, err)
	}
	return nil
}
