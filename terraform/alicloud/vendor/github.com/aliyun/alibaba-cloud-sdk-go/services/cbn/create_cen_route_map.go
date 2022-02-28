package cbn

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

// CreateCenRouteMap invokes the cbn.CreateCenRouteMap API synchronously
func (client *Client) CreateCenRouteMap(request *CreateCenRouteMapRequest) (response *CreateCenRouteMapResponse, err error) {
	response = CreateCreateCenRouteMapResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCenRouteMapWithChan invokes the cbn.CreateCenRouteMap API asynchronously
func (client *Client) CreateCenRouteMapWithChan(request *CreateCenRouteMapRequest) (<-chan *CreateCenRouteMapResponse, <-chan error) {
	responseChan := make(chan *CreateCenRouteMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCenRouteMap(request)
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

// CreateCenRouteMapWithCallback invokes the cbn.CreateCenRouteMap API asynchronously
func (client *Client) CreateCenRouteMapWithCallback(request *CreateCenRouteMapRequest, callback func(response *CreateCenRouteMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCenRouteMapResponse
		var err error
		defer close(result)
		response, err = client.CreateCenRouteMap(request)
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

// CreateCenRouteMapRequest is the request struct for api CreateCenRouteMap
type CreateCenRouteMapRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                    requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CommunityMatchMode                 string           `position:"Query" name:"CommunityMatchMode"`
	MapResult                          string           `position:"Query" name:"MapResult"`
	DestinationRegionIds               *[]string        `position:"Query" name:"DestinationRegionIds"  type:"Repeated"`
	NextPriority                       requests.Integer `position:"Query" name:"NextPriority"`
	DestinationCidrBlocks              *[]string        `position:"Query" name:"DestinationCidrBlocks"  type:"Repeated"`
	SystemPolicy                       requests.Boolean `position:"Query" name:"SystemPolicy"`
	OriginalRouteTableIds              *[]string        `position:"Query" name:"OriginalRouteTableIds"  type:"Repeated"`
	TransitRouterRouteTableId          string           `position:"Query" name:"TransitRouterRouteTableId"`
	SourceInstanceIds                  *[]string        `position:"Query" name:"SourceInstanceIds"  type:"Repeated"`
	SourceRegionIds                    *[]string        `position:"Query" name:"SourceRegionIds"  type:"Repeated"`
	GatewayZoneId                      string           `position:"Query" name:"GatewayZoneId"`
	MatchAsns                          *[]string        `position:"Query" name:"MatchAsns"  type:"Repeated"`
	Preference                         requests.Integer `position:"Query" name:"Preference"`
	OwnerId                            requests.Integer `position:"Query" name:"OwnerId"`
	Priority                           requests.Integer `position:"Query" name:"Priority"`
	DestinationChildInstanceTypes      *[]string        `position:"Query" name:"DestinationChildInstanceTypes"  type:"Repeated"`
	SourceRouteTableIds                *[]string        `position:"Query" name:"SourceRouteTableIds"  type:"Repeated"`
	SourceChildInstanceTypes           *[]string        `position:"Query" name:"SourceChildInstanceTypes"  type:"Repeated"`
	CommunityOperateMode               string           `position:"Query" name:"CommunityOperateMode"`
	OperateCommunitySet                *[]string        `position:"Query" name:"OperateCommunitySet"  type:"Repeated"`
	RouteTypes                         *[]string        `position:"Query" name:"RouteTypes"  type:"Repeated"`
	MatchAddressType                   string           `position:"Query" name:"MatchAddressType"`
	CidrMatchMode                      string           `position:"Query" name:"CidrMatchMode"`
	CenId                              string           `position:"Query" name:"CenId"`
	Description                        string           `position:"Query" name:"Description"`
	SourceInstanceIdsReverseMatch      requests.Boolean `position:"Query" name:"SourceInstanceIdsReverseMatch"`
	DestinationRouteTableIds           *[]string        `position:"Query" name:"DestinationRouteTableIds"  type:"Repeated"`
	SourceZoneIds                      *[]string        `position:"Query" name:"SourceZoneIds"  type:"Repeated"`
	TransmitDirection                  string           `position:"Query" name:"TransmitDirection"`
	DestinationInstanceIds             *[]string        `position:"Query" name:"DestinationInstanceIds"  type:"Repeated"`
	ResourceOwnerAccount               string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                       string           `position:"Query" name:"OwnerAccount"`
	DestinationInstanceIdsReverseMatch requests.Boolean `position:"Query" name:"DestinationInstanceIdsReverseMatch"`
	PrependAsPath                      *[]string        `position:"Query" name:"PrependAsPath"  type:"Repeated"`
	AsPathMatchMode                    string           `position:"Query" name:"AsPathMatchMode"`
	MatchCommunitySet                  *[]string        `position:"Query" name:"MatchCommunitySet"  type:"Repeated"`
	CenRegionId                        string           `position:"Query" name:"CenRegionId"`
}

// CreateCenRouteMapResponse is the response struct for api CreateCenRouteMap
type CreateCenRouteMapResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	RouteMapId string `json:"RouteMapId" xml:"RouteMapId"`
}

// CreateCreateCenRouteMapRequest creates a request to invoke CreateCenRouteMap API
func CreateCreateCenRouteMapRequest() (request *CreateCenRouteMapRequest) {
	request = &CreateCenRouteMapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateCenRouteMap", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCenRouteMapResponse creates a response to parse from CreateCenRouteMap response
func CreateCreateCenRouteMapResponse() (response *CreateCenRouteMapResponse) {
	response = &CreateCenRouteMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
