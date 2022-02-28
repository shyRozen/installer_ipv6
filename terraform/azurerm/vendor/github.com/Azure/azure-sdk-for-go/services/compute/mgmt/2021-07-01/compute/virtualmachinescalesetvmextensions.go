package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualMachineScaleSetVMExtensionsClient is the compute Client
type VirtualMachineScaleSetVMExtensionsClient struct {
	BaseClient
}

// NewVirtualMachineScaleSetVMExtensionsClient creates an instance of the VirtualMachineScaleSetVMExtensionsClient
// client.
func NewVirtualMachineScaleSetVMExtensionsClient(subscriptionID string) VirtualMachineScaleSetVMExtensionsClient {
	return NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI creates an instance of the
// VirtualMachineScaleSetVMExtensionsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetVMExtensionsClient {
	return VirtualMachineScaleSetVMExtensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update the VMSS VM extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// VMExtensionName - the name of the virtual machine extension.
// extensionParameters - parameters supplied to the Create Virtual Machine Extension operation.
func (client VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension) (result VirtualMachineScaleSetVMExtensionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, VMExtensionName, extensionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	extensionParameters.Name = nil
	extensionParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithJSON(extensionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdateSender(req *http.Request) (future VirtualMachineScaleSetVMExtensionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMExtensionsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualMachineScaleSetVMExtension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the VMSS VM extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// VMExtensionName - the name of the virtual machine extension.
func (client VirtualMachineScaleSetVMExtensionsClient) Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string) (result VirtualMachineScaleSetVMExtensionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMExtensionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, VMExtensionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineScaleSetVMExtensionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMExtensionsClient) DeleteSender(req *http.Request) (future VirtualMachineScaleSetVMExtensionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMExtensionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get the VMSS VM extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// VMExtensionName - the name of the virtual machine extension.
// expand - the expand expression to apply on the operation.
func (client VirtualMachineScaleSetVMExtensionsClient) Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, expand string) (result VirtualMachineScaleSetVMExtension, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMExtensionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, VMExtensionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineScaleSetVMExtensionsClient) GetPreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMExtensionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMExtensionsClient) GetResponder(resp *http.Response) (result VirtualMachineScaleSetVMExtension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the operation to get all extensions of an instance in Virtual Machine Scaleset.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// expand - the expand expression to apply on the operation.
func (client VirtualMachineScaleSetVMExtensionsClient) List(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand string) (result VirtualMachineScaleSetVMExtensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMExtensionsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineScaleSetVMExtensionsClient) ListPreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMExtensionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMExtensionsClient) ListResponder(resp *http.Response) (result VirtualMachineScaleSetVMExtensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update the operation to update the VMSS VM extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMScaleSetName - the name of the VM scale set.
// instanceID - the instance ID of the virtual machine.
// VMExtensionName - the name of the virtual machine extension.
// extensionParameters - parameters supplied to the Update Virtual Machine Extension operation.
func (client VirtualMachineScaleSetVMExtensionsClient) Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate) (result VirtualMachineScaleSetVMExtensionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineScaleSetVMExtensionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, VMScaleSetName, instanceID, VMExtensionName, extensionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMExtensionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualMachineScaleSetVMExtensionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, VMExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	extensionParameters.Name = nil
	extensionParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithJSON(extensionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMExtensionsClient) UpdateSender(req *http.Request) (future VirtualMachineScaleSetVMExtensionsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMExtensionsClient) UpdateResponder(resp *http.Response) (result VirtualMachineScaleSetVMExtension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
