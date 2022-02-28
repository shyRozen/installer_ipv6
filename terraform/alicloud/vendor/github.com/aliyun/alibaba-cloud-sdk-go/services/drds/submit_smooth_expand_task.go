package drds

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

// SubmitSmoothExpandTask invokes the drds.SubmitSmoothExpandTask API synchronously
func (client *Client) SubmitSmoothExpandTask(request *SubmitSmoothExpandTaskRequest) (response *SubmitSmoothExpandTaskResponse, err error) {
	response = CreateSubmitSmoothExpandTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitSmoothExpandTaskWithChan invokes the drds.SubmitSmoothExpandTask API asynchronously
func (client *Client) SubmitSmoothExpandTaskWithChan(request *SubmitSmoothExpandTaskRequest) (<-chan *SubmitSmoothExpandTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitSmoothExpandTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSmoothExpandTask(request)
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

// SubmitSmoothExpandTaskWithCallback invokes the drds.SubmitSmoothExpandTask API asynchronously
func (client *Client) SubmitSmoothExpandTaskWithCallback(request *SubmitSmoothExpandTaskRequest, callback func(response *SubmitSmoothExpandTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSmoothExpandTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitSmoothExpandTask(request)
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

// SubmitSmoothExpandTaskRequest is the request struct for api SubmitSmoothExpandTask
type SubmitSmoothExpandTaskRequest struct {
	*requests.RpcRequest
	DrdsInstanceId       string                                     `position:"Query" name:"DrdsInstanceId"`
	DbInstanceIsCreating requests.Boolean                           `position:"Query" name:"DbInstanceIsCreating"`
	RdsSuperInstances    *[]SubmitSmoothExpandTaskRdsSuperInstances `position:"Query" name:"RdsSuperInstances"  type:"Repeated"`
	DbName               string                                     `position:"Query" name:"DbName"`
	TransferTaskInfos    *[]SubmitSmoothExpandTaskTransferTaskInfos `position:"Query" name:"TransferTaskInfos"  type:"Repeated"`
}

// SubmitSmoothExpandTaskRdsSuperInstances is a repeated param struct in SubmitSmoothExpandTaskRequest
type SubmitSmoothExpandTaskRdsSuperInstances struct {
	Password    string `name:"Password"`
	AccountName string `name:"AccountName"`
	RdsName     string `name:"RdsName"`
}

// SubmitSmoothExpandTaskTransferTaskInfos is a repeated param struct in SubmitSmoothExpandTaskRequest
type SubmitSmoothExpandTaskTransferTaskInfos struct {
	DbName          string `name:"DbName"`
	SrcInstanceName string `name:"SrcInstanceName"`
	InstanceType    string `name:"InstanceType"`
	DstInstanceName string `name:"DstInstanceName"`
}

// SubmitSmoothExpandTaskResponse is the response struct for api SubmitSmoothExpandTask
type SubmitSmoothExpandTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSubmitSmoothExpandTaskRequest creates a request to invoke SubmitSmoothExpandTask API
func CreateSubmitSmoothExpandTaskRequest() (request *SubmitSmoothExpandTaskRequest) {
	request = &SubmitSmoothExpandTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SubmitSmoothExpandTask", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitSmoothExpandTaskResponse creates a response to parse from SubmitSmoothExpandTask response
func CreateSubmitSmoothExpandTaskResponse() (response *SubmitSmoothExpandTaskResponse) {
	response = &SubmitSmoothExpandTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
