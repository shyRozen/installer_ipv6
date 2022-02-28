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

// GetUserOutputStatisticInfo invokes the emr.GetUserOutputStatisticInfo API synchronously
func (client *Client) GetUserOutputStatisticInfo(request *GetUserOutputStatisticInfoRequest) (response *GetUserOutputStatisticInfoResponse, err error) {
	response = CreateGetUserOutputStatisticInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserOutputStatisticInfoWithChan invokes the emr.GetUserOutputStatisticInfo API asynchronously
func (client *Client) GetUserOutputStatisticInfoWithChan(request *GetUserOutputStatisticInfoRequest) (<-chan *GetUserOutputStatisticInfoResponse, <-chan error) {
	responseChan := make(chan *GetUserOutputStatisticInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserOutputStatisticInfo(request)
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

// GetUserOutputStatisticInfoWithCallback invokes the emr.GetUserOutputStatisticInfo API asynchronously
func (client *Client) GetUserOutputStatisticInfoWithCallback(request *GetUserOutputStatisticInfoRequest, callback func(response *GetUserOutputStatisticInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserOutputStatisticInfoResponse
		var err error
		defer close(result)
		response, err = client.GetUserOutputStatisticInfo(request)
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

// GetUserOutputStatisticInfoRequest is the request struct for api GetUserOutputStatisticInfo
type GetUserOutputStatisticInfoRequest struct {
	*requests.RpcRequest
	FromDatetime    string           `position:"Query" name:"FromDatetime"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ToDatetime      string           `position:"Query" name:"ToDatetime"`
}

// GetUserOutputStatisticInfoResponse is the response struct for api GetUserOutputStatisticInfo
type GetUserOutputStatisticInfoResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	UserOutputList UserOutputList `json:"UserOutputList" xml:"UserOutputList"`
}

// CreateGetUserOutputStatisticInfoRequest creates a request to invoke GetUserOutputStatisticInfo API
func CreateGetUserOutputStatisticInfoRequest() (request *GetUserOutputStatisticInfoRequest) {
	request = &GetUserOutputStatisticInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "GetUserOutputStatisticInfo", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserOutputStatisticInfoResponse creates a response to parse from GetUserOutputStatisticInfo response
func CreateGetUserOutputStatisticInfoResponse() (response *GetUserOutputStatisticInfoResponse) {
	response = &GetUserOutputStatisticInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
