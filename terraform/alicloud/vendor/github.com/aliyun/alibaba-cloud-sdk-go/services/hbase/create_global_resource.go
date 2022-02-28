package hbase

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

// CreateGlobalResource invokes the hbase.CreateGlobalResource API synchronously
func (client *Client) CreateGlobalResource(request *CreateGlobalResourceRequest) (response *CreateGlobalResourceResponse, err error) {
	response = CreateCreateGlobalResourceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGlobalResourceWithChan invokes the hbase.CreateGlobalResource API asynchronously
func (client *Client) CreateGlobalResourceWithChan(request *CreateGlobalResourceRequest) (<-chan *CreateGlobalResourceResponse, <-chan error) {
	responseChan := make(chan *CreateGlobalResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGlobalResource(request)
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

// CreateGlobalResourceWithCallback invokes the hbase.CreateGlobalResource API asynchronously
func (client *Client) CreateGlobalResourceWithCallback(request *CreateGlobalResourceRequest, callback func(response *CreateGlobalResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGlobalResourceResponse
		var err error
		defer close(result)
		response, err = client.CreateGlobalResource(request)
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

// CreateGlobalResourceRequest is the request struct for api CreateGlobalResource
type CreateGlobalResourceRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	ClusterId    string `position:"Query" name:"ClusterId"`
	ResourceType string `position:"Query" name:"ResourceType"`
	ResourceName string `position:"Query" name:"ResourceName"`
}

// CreateGlobalResourceResponse is the response struct for api CreateGlobalResource
type CreateGlobalResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateGlobalResourceRequest creates a request to invoke CreateGlobalResource API
func CreateCreateGlobalResourceRequest() (request *CreateGlobalResourceRequest) {
	request = &CreateGlobalResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "CreateGlobalResource", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGlobalResourceResponse creates a response to parse from CreateGlobalResource response
func CreateCreateGlobalResourceResponse() (response *CreateGlobalResourceResponse) {
	response = &CreateGlobalResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
