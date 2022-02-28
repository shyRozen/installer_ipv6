package ddosbgp

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

// DescribeDdosEvent invokes the ddosbgp.DescribeDdosEvent API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeddosevent.html
func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (response *DescribeDdosEventResponse, err error) {
	response = CreateDescribeDdosEventResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDdosEventWithChan invokes the ddosbgp.DescribeDdosEvent API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeddosevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDdosEventWithChan(request *DescribeDdosEventRequest) (<-chan *DescribeDdosEventResponse, <-chan error) {
	responseChan := make(chan *DescribeDdosEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDdosEvent(request)
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

// DescribeDdosEventWithCallback invokes the ddosbgp.DescribeDdosEvent API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeddosevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDdosEventWithCallback(request *DescribeDdosEventRequest, callback func(response *DescribeDdosEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDdosEventResponse
		var err error
		defer close(result)
		response, err = client.DescribeDdosEvent(request)
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

// DescribeDdosEventRequest is the request struct for api DescribeDdosEvent
type DescribeDdosEventRequest struct {
	*requests.RpcRequest
	StartTime        requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId  string           `position:"Query" name:"ResourceGroupId"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	ResourceRegionId string           `position:"Query" name:"ResourceRegionId"`
	Ip               string           `position:"Query" name:"Ip"`
	EndTime          requests.Integer `position:"Query" name:"EndTime"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
	PageNo           requests.Integer `position:"Query" name:"PageNo"`
}

// DescribeDdosEventResponse is the response struct for api DescribeDdosEvent
type DescribeDdosEventResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Total     int64   `json:"Total" xml:"Total"`
	Events    []Event `json:"Events" xml:"Events"`
}

// CreateDescribeDdosEventRequest creates a request to invoke DescribeDdosEvent API
func CreateDescribeDdosEventRequest() (request *DescribeDdosEventRequest) {
	request = &DescribeDdosEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeDdosEvent", "ddosbgp", "openAPI")
	return
}

// CreateDescribeDdosEventResponse creates a response to parse from DescribeDdosEvent response
func CreateDescribeDdosEventResponse() (response *DescribeDdosEventResponse) {
	response = &DescribeDdosEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
