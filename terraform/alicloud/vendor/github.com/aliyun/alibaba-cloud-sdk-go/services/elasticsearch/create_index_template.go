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

// CreateIndexTemplate invokes the elasticsearch.CreateIndexTemplate API synchronously
func (client *Client) CreateIndexTemplate(request *CreateIndexTemplateRequest) (response *CreateIndexTemplateResponse, err error) {
	response = CreateCreateIndexTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIndexTemplateWithChan invokes the elasticsearch.CreateIndexTemplate API asynchronously
func (client *Client) CreateIndexTemplateWithChan(request *CreateIndexTemplateRequest) (<-chan *CreateIndexTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateIndexTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIndexTemplate(request)
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

// CreateIndexTemplateWithCallback invokes the elasticsearch.CreateIndexTemplate API asynchronously
func (client *Client) CreateIndexTemplateWithCallback(request *CreateIndexTemplateRequest, callback func(response *CreateIndexTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIndexTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateIndexTemplate(request)
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

// CreateIndexTemplateRequest is the request struct for api CreateIndexTemplate
type CreateIndexTemplateRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

// CreateIndexTemplateResponse is the response struct for api CreateIndexTemplate
type CreateIndexTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateCreateIndexTemplateRequest creates a request to invoke CreateIndexTemplate API
func CreateCreateIndexTemplateRequest() (request *CreateIndexTemplateRequest) {
	request = &CreateIndexTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "CreateIndexTemplate", "/openapi/instances/[InstanceId]/index-templates", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateIndexTemplateResponse creates a response to parse from CreateIndexTemplate response
func CreateCreateIndexTemplateResponse() (response *CreateIndexTemplateResponse) {
	response = &CreateIndexTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
