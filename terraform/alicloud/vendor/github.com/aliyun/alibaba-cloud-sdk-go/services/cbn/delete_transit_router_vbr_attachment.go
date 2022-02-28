package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteTransitRouterVbrAttachment invokes the cbn.DeleteTransitRouterVbrAttachment API synchronously
func (client *Client) DeleteTransitRouterVbrAttachment(request *DeleteTransitRouterVbrAttachmentRequest) (response *DeleteTransitRouterVbrAttachmentResponse, err error) {
	response = CreateDeleteTransitRouterVbrAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTransitRouterVbrAttachmentWithChan invokes the cbn.DeleteTransitRouterVbrAttachment API asynchronously
func (client *Client) DeleteTransitRouterVbrAttachmentWithChan(request *DeleteTransitRouterVbrAttachmentRequest) (<-chan *DeleteTransitRouterVbrAttachmentResponse, <-chan error) {
	responseChan := make(chan *DeleteTransitRouterVbrAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTransitRouterVbrAttachment(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteTransitRouterVbrAttachmentWithCallback invokes the cbn.DeleteTransitRouterVbrAttachment API asynchronously
func (client *Client) DeleteTransitRouterVbrAttachmentWithCallback(request *DeleteTransitRouterVbrAttachmentRequest, callback func(response *DeleteTransitRouterVbrAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTransitRouterVbrAttachmentResponse
		var err error
		defer close(result)
		response, err = client.DeleteTransitRouterVbrAttachment(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteTransitRouterVbrAttachmentRequest is the request struct for api DeleteTransitRouterVbrAttachment
type DeleteTransitRouterVbrAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType              string           `position:"Query" name:"ResourceType"`
	TransitRouterAttachmentId string           `position:"Query" name:"TransitRouterAttachmentId"`
}

// DeleteTransitRouterVbrAttachmentResponse is the response struct for api DeleteTransitRouterVbrAttachment
type DeleteTransitRouterVbrAttachmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTransitRouterVbrAttachmentRequest creates a request to invoke DeleteTransitRouterVbrAttachment API
func CreateDeleteTransitRouterVbrAttachmentRequest() (request *DeleteTransitRouterVbrAttachmentRequest) {
	request = &DeleteTransitRouterVbrAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteTransitRouterVbrAttachment", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTransitRouterVbrAttachmentResponse creates a response to parse from DeleteTransitRouterVbrAttachment response
func CreateDeleteTransitRouterVbrAttachmentResponse() (response *DeleteTransitRouterVbrAttachmentResponse) {
	response = &DeleteTransitRouterVbrAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
