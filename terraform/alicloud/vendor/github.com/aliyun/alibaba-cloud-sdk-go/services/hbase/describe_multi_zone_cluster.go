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

// DescribeMultiZoneCluster invokes the hbase.DescribeMultiZoneCluster API synchronously
func (client *Client) DescribeMultiZoneCluster(request *DescribeMultiZoneClusterRequest) (response *DescribeMultiZoneClusterResponse, err error) {
	response = CreateDescribeMultiZoneClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMultiZoneClusterWithChan invokes the hbase.DescribeMultiZoneCluster API asynchronously
func (client *Client) DescribeMultiZoneClusterWithChan(request *DescribeMultiZoneClusterRequest) (<-chan *DescribeMultiZoneClusterResponse, <-chan error) {
	responseChan := make(chan *DescribeMultiZoneClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMultiZoneCluster(request)
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

// DescribeMultiZoneClusterWithCallback invokes the hbase.DescribeMultiZoneCluster API asynchronously
func (client *Client) DescribeMultiZoneClusterWithCallback(request *DescribeMultiZoneClusterRequest, callback func(response *DescribeMultiZoneClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMultiZoneClusterResponse
		var err error
		defer close(result)
		response, err = client.DescribeMultiZoneCluster(request)
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

// DescribeMultiZoneClusterRequest is the request struct for api DescribeMultiZoneCluster
type DescribeMultiZoneClusterRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeMultiZoneClusterResponse is the response struct for api DescribeMultiZoneCluster
type DescribeMultiZoneClusterResponse struct {
	*responses.BaseResponse
	AutoRenewal             bool                           `json:"AutoRenewal" xml:"AutoRenewal"`
	ClusterId               string                         `json:"ClusterId" xml:"ClusterId"`
	ClusterName             string                         `json:"ClusterName" xml:"ClusterName"`
	CreatedTime             string                         `json:"CreatedTime" xml:"CreatedTime"`
	CreatedTimeUTC          string                         `json:"CreatedTimeUTC" xml:"CreatedTimeUTC"`
	Duration                int                            `json:"Duration" xml:"Duration"`
	Engine                  string                         `json:"Engine" xml:"Engine"`
	ExpireTime              string                         `json:"ExpireTime" xml:"ExpireTime"`
	ExpireTimeUTC           string                         `json:"ExpireTimeUTC" xml:"ExpireTimeUTC"`
	InstanceId              string                         `json:"InstanceId" xml:"InstanceId"`
	InstanceName            string                         `json:"InstanceName" xml:"InstanceName"`
	IsDeletionProtection    bool                           `json:"IsDeletionProtection" xml:"IsDeletionProtection"`
	MaintainEndTime         string                         `json:"MaintainEndTime" xml:"MaintainEndTime"`
	MaintainStartTime       string                         `json:"MaintainStartTime" xml:"MaintainStartTime"`
	MajorVersion            string                         `json:"MajorVersion" xml:"MajorVersion"`
	MasterDiskSize          int                            `json:"MasterDiskSize" xml:"MasterDiskSize"`
	MasterDiskType          string                         `json:"MasterDiskType" xml:"MasterDiskType"`
	MasterInstanceType      string                         `json:"MasterInstanceType" xml:"MasterInstanceType"`
	MasterNodeCount         int                            `json:"MasterNodeCount" xml:"MasterNodeCount"`
	CoreDiskCount           string                         `json:"CoreDiskCount" xml:"CoreDiskCount"`
	CoreDiskSize            int                            `json:"CoreDiskSize" xml:"CoreDiskSize"`
	CoreDiskType            string                         `json:"CoreDiskType" xml:"CoreDiskType"`
	CoreInstanceType        string                         `json:"CoreInstanceType" xml:"CoreInstanceType"`
	CoreNodeCount           int                            `json:"CoreNodeCount" xml:"CoreNodeCount"`
	LogDiskCount            string                         `json:"LogDiskCount" xml:"LogDiskCount"`
	LogDiskSize             int                            `json:"LogDiskSize" xml:"LogDiskSize"`
	LogDiskType             string                         `json:"LogDiskType" xml:"LogDiskType"`
	LogInstanceType         string                         `json:"LogInstanceType" xml:"LogInstanceType"`
	LogNodeCount            int                            `json:"LogNodeCount" xml:"LogNodeCount"`
	ModuleId                int                            `json:"ModuleId" xml:"ModuleId"`
	ModuleStackVersion      string                         `json:"ModuleStackVersion" xml:"ModuleStackVersion"`
	NetworkType             string                         `json:"NetworkType" xml:"NetworkType"`
	ParentId                string                         `json:"ParentId" xml:"ParentId"`
	PayType                 string                         `json:"PayType" xml:"PayType"`
	RegionId                string                         `json:"RegionId" xml:"RegionId"`
	RequestId               string                         `json:"RequestId" xml:"RequestId"`
	Status                  string                         `json:"Status" xml:"Status"`
	MultiZoneCombination    string                         `json:"MultiZoneCombination" xml:"MultiZoneCombination"`
	PrimaryZoneId           string                         `json:"PrimaryZoneId" xml:"PrimaryZoneId"`
	PrimaryVSwitchIds       string                         `json:"PrimaryVSwitchIds" xml:"PrimaryVSwitchIds"`
	StandbyZoneId           string                         `json:"StandbyZoneId" xml:"StandbyZoneId"`
	StandbyVSwitchIds       string                         `json:"StandbyVSwitchIds" xml:"StandbyVSwitchIds"`
	ArbiterZoneId           string                         `json:"ArbiterZoneId" xml:"ArbiterZoneId"`
	ArbiterVSwitchIds       string                         `json:"ArbiterVSwitchIds" xml:"ArbiterVSwitchIds"`
	VpcId                   string                         `json:"VpcId" xml:"VpcId"`
	ResourceGroupId         string                         `json:"ResourceGroupId" xml:"ResourceGroupId"`
	EncryptionType          string                         `json:"EncryptionType" xml:"EncryptionType"`
	EncryptionKey           string                         `json:"EncryptionKey" xml:"EncryptionKey"`
	MultiZoneInstanceModels MultiZoneInstanceModels        `json:"MultiZoneInstanceModels" xml:"MultiZoneInstanceModels"`
	Tags                    TagsInDescribeMultiZoneCluster `json:"Tags" xml:"Tags"`
}

// CreateDescribeMultiZoneClusterRequest creates a request to invoke DescribeMultiZoneCluster API
func CreateDescribeMultiZoneClusterRequest() (request *DescribeMultiZoneClusterRequest) {
	request = &DescribeMultiZoneClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeMultiZoneCluster", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMultiZoneClusterResponse creates a response to parse from DescribeMultiZoneCluster response
func CreateDescribeMultiZoneClusterResponse() (response *DescribeMultiZoneClusterResponse) {
	response = &DescribeMultiZoneClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
