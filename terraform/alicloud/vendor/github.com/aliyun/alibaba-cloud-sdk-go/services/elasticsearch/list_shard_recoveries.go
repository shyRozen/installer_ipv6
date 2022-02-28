package elasticsearch

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

// ListShardRecoveries invokes the elasticsearch.ListShardRecoveries API synchronously
func (client *Client) ListShardRecoveries(request *ListShardRecoveriesRequest) (response *ListShardRecoveriesResponse, err error) {
	response = CreateListShardRecoveriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListShardRecoveriesWithChan invokes the elasticsearch.ListShardRecoveries API asynchronously
func (client *Client) ListShardRecoveriesWithChan(request *ListShardRecoveriesRequest) (<-chan *ListShardRecoveriesResponse, <-chan error) {
	responseChan := make(chan *ListShardRecoveriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListShardRecoveries(request)
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

// ListShardRecoveriesWithCallback invokes the elasticsearch.ListShardRecoveries API asynchronously
func (client *Client) ListShardRecoveriesWithCallback(request *ListShardRecoveriesRequest, callback func(response *ListShardRecoveriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListShardRecoveriesResponse
		var err error
		defer close(result)
		response, err = client.ListShardRecoveries(request)
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

// ListShardRecoveriesRequest is the request struct for api ListShardRecoveries
type ListShardRecoveriesRequest struct {
	*requests.RoaRequest
	ActiveOnly requests.Boolean `position:"Query" name:"activeOnly"`
	InstanceId string           `position:"Path" name:"InstanceId"`
}

// ListShardRecoveriesResponse is the response struct for api ListShardRecoveries
type ListShardRecoveriesResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListShardRecoveriesRequest creates a request to invoke ListShardRecoveries API
func CreateListShardRecoveriesRequest() (request *ListShardRecoveriesRequest) {
	request = &ListShardRecoveriesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListShardRecoveries", "/openapi/instances/[InstanceId]/cat-recovery", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListShardRecoveriesResponse creates a response to parse from ListShardRecoveries response
func CreateListShardRecoveriesResponse() (response *ListShardRecoveriesResponse) {
	response = &ListShardRecoveriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
