package emr

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

// DescribeEmrMainVersion invokes the emr.DescribeEmrMainVersion API synchronously
func (client *Client) DescribeEmrMainVersion(request *DescribeEmrMainVersionRequest) (response *DescribeEmrMainVersionResponse, err error) {
	response = CreateDescribeEmrMainVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEmrMainVersionWithChan invokes the emr.DescribeEmrMainVersion API asynchronously
func (client *Client) DescribeEmrMainVersionWithChan(request *DescribeEmrMainVersionRequest) (<-chan *DescribeEmrMainVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeEmrMainVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEmrMainVersion(request)
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

// DescribeEmrMainVersionWithCallback invokes the emr.DescribeEmrMainVersion API asynchronously
func (client *Client) DescribeEmrMainVersionWithCallback(request *DescribeEmrMainVersionRequest, callback func(response *DescribeEmrMainVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEmrMainVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeEmrMainVersion(request)
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

// DescribeEmrMainVersionRequest is the request struct for api DescribeEmrMainVersion
type DescribeEmrMainVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EmrVersion      string           `position:"Query" name:"EmrVersion"`
}

// DescribeEmrMainVersionResponse is the response struct for api DescribeEmrMainVersion
type DescribeEmrMainVersionResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	EmrMainVersion EmrMainVersion `json:"EmrMainVersion" xml:"EmrMainVersion"`
}

// CreateDescribeEmrMainVersionRequest creates a request to invoke DescribeEmrMainVersion API
func CreateDescribeEmrMainVersionRequest() (request *DescribeEmrMainVersionRequest) {
	request = &DescribeEmrMainVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeEmrMainVersion", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEmrMainVersionResponse creates a response to parse from DescribeEmrMainVersion response
func CreateDescribeEmrMainVersionResponse() (response *DescribeEmrMainVersionResponse) {
	response = &DescribeEmrMainVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
