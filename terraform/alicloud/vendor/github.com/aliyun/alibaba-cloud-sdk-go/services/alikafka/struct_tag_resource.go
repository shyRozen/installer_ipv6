package alikafka

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

// TagResource is a nested struct in alikafka response
type TagResource struct {
	TagKey       string `json:"TagKey" xml:"TagKey"`
	TagValue     string `json:"TagValue" xml:"TagValue"`
	ResourceType string `json:"ResourceType" xml:"ResourceType"`
	ResourceId   string `json:"ResourceId" xml:"ResourceId"`
}
