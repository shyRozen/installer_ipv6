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

// AssociateHaVip invokes the vpc.AssociateHaVip API synchronously
func (client *Client) AssociateHaVip(request *AssociateHaVipRequest) (response *AssociateHaVipResponse, err error) {
	response = CreateAssociateHaVipResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateHaVipWithChan invokes the vpc.AssociateHaVip API asynchronously
func (client *Client) AssociateHaVipWithChan(request *AssociateHaVipRequest) (<-chan *AssociateHaVipResponse, <-chan error) {
	responseChan := make(chan *AssociateHaVipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateHaVip(request)
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

// AssociateHaVipWithCallback invokes the vpc.AssociateHaVip API asynchronously
func (client *Client) AssociateHaVipWithCallback(request *AssociateHaVipRequest, callback func(response *AssociateHaVipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateHaVipResponse
		var err error
		defer close(result)
		response, err = client.AssociateHaVip(request)
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

// AssociateHaVipRequest is the request struct for api AssociateHaVip
type AssociateHaVipRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	HaVipId              string           `position:"Query" name:"HaVipId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// AssociateHaVipResponse is the response struct for api AssociateHaVip
type AssociateHaVipResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateHaVipRequest creates a request to invoke AssociateHaVip API
func CreateAssociateHaVipRequest() (request *AssociateHaVipRequest) {
	request = &AssociateHaVipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateHaVip", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateHaVipResponse creates a response to parse from AssociateHaVip response
func CreateAssociateHaVipResponse() (response *AssociateHaVipResponse) {
	response = &AssociateHaVipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
