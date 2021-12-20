package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// IBMPIStorageCapacityClient ..
type IBMPIStorageCapacityClient struct {
	IBMPIClient
}

// NewIBMPIStorageCapacityClient ...
func NewIBMPIStorageCapacityClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIStorageCapacityClient {
	return &IBMPIStorageCapacityClient{
		*NewIBMPIClient(ctx, sess, cloudInstanceID),
	}
}

//Storage capacity for all available storage pools in a region
func (f *IBMPIStorageCapacityClient) GetAllStoragePoolsCapacity() (*models.StoragePoolsCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityPoolsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityPoolsGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools")
	}
	return resp.Payload, nil
}

// Storage capacity for a storage pool in a region
func (f *IBMPIStorageCapacityClient) GetStoragePoolCapacity(storagePool string) (*models.StoragePoolCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityPoolsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithStoragePoolName(storagePool)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityPoolsGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s: %w", storagePool, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s", storagePool)
	}
	return resp.Payload, nil
}

// Storage capacity for a storage type in a region
func (f *IBMPIStorageCapacityClient) GetStorageTypeCapacity(storageType string) (*models.StorageTypeCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityTypesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithStorageTypeName(storageType)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityTypesGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s: %w", storageType, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s", storageType)
	}
	return resp.Payload, nil
}

// Storage capacity for all available storage types in a region
func (f *IBMPIStorageCapacityClient) GetAllStorageTypesCapacity() (*models.StorageTypesCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityTypesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityTypesGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types")
	}
	return resp.Payload, nil
}
