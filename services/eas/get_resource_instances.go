package eas

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

// GetResourceInstances invokes the eas.GetResourceInstances API synchronously
// api document: https://help.aliyun.com/api/eas/getresourceinstances.html
func (client *Client) GetResourceInstances(request *GetResourceInstancesRequest) (response *GetResourceInstancesResponse, err error) {
	response = CreateGetResourceInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceInstancesWithChan invokes the eas.GetResourceInstances API asynchronously
// api document: https://help.aliyun.com/api/eas/getresourceinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourceInstancesWithChan(request *GetResourceInstancesRequest) (<-chan *GetResourceInstancesResponse, <-chan error) {
	responseChan := make(chan *GetResourceInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceInstances(request)
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

// GetResourceInstancesWithCallback invokes the eas.GetResourceInstances API asynchronously
// api document: https://help.aliyun.com/api/eas/getresourceinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourceInstancesWithCallback(request *GetResourceInstancesRequest, callback func(response *GetResourceInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceInstancesResponse
		var err error
		defer close(result)
		response, err = client.GetResourceInstances(request)
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

// GetResourceInstancesRequest is the request struct for api GetResourceInstances
type GetResourceInstancesRequest struct {
	*requests.RoaRequest
	ClusterId    string `position:"Path" name:"cluster_id"`
	ResourceName string `position:"Path" name:"resource_name"`
}

// GetResourceInstancesResponse is the response struct for api GetResourceInstances
type GetResourceInstancesResponse struct {
	*responses.BaseResponse
}

// CreateGetResourceInstancesRequest creates a request to invoke GetResourceInstances API
func CreateGetResourceInstancesRequest() (request *GetResourceInstancesRequest) {
	request = &GetResourceInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "GetResourceInstances", "/api/resources/[cluster_id]/[resource_name]/instances", "", "")
	request.Method = requests.GET
	return
}

// CreateGetResourceInstancesResponse creates a response to parse from GetResourceInstances response
func CreateGetResourceInstancesResponse() (response *GetResourceInstancesResponse) {
	response = &GetResourceInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
