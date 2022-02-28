package vpc

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

// ModifyExpressCloudConnectionAttribute invokes the vpc.ModifyExpressCloudConnectionAttribute API synchronously
func (client *Client) ModifyExpressCloudConnectionAttribute(request *ModifyExpressCloudConnectionAttributeRequest) (response *ModifyExpressCloudConnectionAttributeResponse, err error) {
	response = CreateModifyExpressCloudConnectionAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyExpressCloudConnectionAttributeWithChan invokes the vpc.ModifyExpressCloudConnectionAttribute API asynchronously
func (client *Client) ModifyExpressCloudConnectionAttributeWithChan(request *ModifyExpressCloudConnectionAttributeRequest) (<-chan *ModifyExpressCloudConnectionAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyExpressCloudConnectionAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyExpressCloudConnectionAttribute(request)
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

// ModifyExpressCloudConnectionAttributeWithCallback invokes the vpc.ModifyExpressCloudConnectionAttribute API asynchronously
func (client *Client) ModifyExpressCloudConnectionAttributeWithCallback(request *ModifyExpressCloudConnectionAttributeRequest, callback func(response *ModifyExpressCloudConnectionAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyExpressCloudConnectionAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyExpressCloudConnectionAttribute(request)
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

// ModifyExpressCloudConnectionAttributeRequest is the request struct for api ModifyExpressCloudConnectionAttribute
type ModifyExpressCloudConnectionAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	EccId                string           `position:"Query" name:"EccId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	CeIp                 string           `position:"Query" name:"CeIp"`
	BgpAs                string           `position:"Query" name:"BgpAs"`
	PeIp                 string           `position:"Query" name:"PeIp"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyExpressCloudConnectionAttributeResponse is the response struct for api ModifyExpressCloudConnectionAttribute
type ModifyExpressCloudConnectionAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyExpressCloudConnectionAttributeRequest creates a request to invoke ModifyExpressCloudConnectionAttribute API
func CreateModifyExpressCloudConnectionAttributeRequest() (request *ModifyExpressCloudConnectionAttributeRequest) {
	request = &ModifyExpressCloudConnectionAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyExpressCloudConnectionAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyExpressCloudConnectionAttributeResponse creates a response to parse from ModifyExpressCloudConnectionAttribute response
func CreateModifyExpressCloudConnectionAttributeResponse() (response *ModifyExpressCloudConnectionAttributeResponse) {
	response = &ModifyExpressCloudConnectionAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
