package polardb

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

// ModifyDBClusterTDE invokes the polardb.ModifyDBClusterTDE API synchronously
func (client *Client) ModifyDBClusterTDE(request *ModifyDBClusterTDERequest) (response *ModifyDBClusterTDEResponse, err error) {
	response = CreateModifyDBClusterTDEResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterTDEWithChan invokes the polardb.ModifyDBClusterTDE API asynchronously
func (client *Client) ModifyDBClusterTDEWithChan(request *ModifyDBClusterTDERequest) (<-chan *ModifyDBClusterTDEResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterTDEResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterTDE(request)
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

// ModifyDBClusterTDEWithCallback invokes the polardb.ModifyDBClusterTDE API asynchronously
func (client *Client) ModifyDBClusterTDEWithCallback(request *ModifyDBClusterTDERequest, callback func(response *ModifyDBClusterTDEResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterTDEResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterTDE(request)
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

// ModifyDBClusterTDERequest is the request struct for api ModifyDBClusterTDE
type ModifyDBClusterTDERequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EncryptionKey        string           `position:"Query" name:"EncryptionKey"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RoleArn              string           `position:"Query" name:"RoleArn"`
	EncryptNewTables     string           `position:"Query" name:"EncryptNewTables"`
	TDEStatus            string           `position:"Query" name:"TDEStatus"`
}

// ModifyDBClusterTDEResponse is the response struct for api ModifyDBClusterTDE
type ModifyDBClusterTDEResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterTDERequest creates a request to invoke ModifyDBClusterTDE API
func CreateModifyDBClusterTDERequest() (request *ModifyDBClusterTDERequest) {
	request = &ModifyDBClusterTDERequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterTDE", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterTDEResponse creates a response to parse from ModifyDBClusterTDE response
func CreateModifyDBClusterTDEResponse() (response *ModifyDBClusterTDEResponse) {
	response = &ModifyDBClusterTDEResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
