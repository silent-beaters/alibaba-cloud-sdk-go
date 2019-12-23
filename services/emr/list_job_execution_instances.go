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

// ListJobExecutionInstances invokes the emr.ListJobExecutionInstances API synchronously
// api document: https://help.aliyun.com/api/emr/listjobexecutioninstances.html
func (client *Client) ListJobExecutionInstances(request *ListJobExecutionInstancesRequest) (response *ListJobExecutionInstancesResponse, err error) {
	response = CreateListJobExecutionInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobExecutionInstancesWithChan invokes the emr.ListJobExecutionInstances API asynchronously
// api document: https://help.aliyun.com/api/emr/listjobexecutioninstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobExecutionInstancesWithChan(request *ListJobExecutionInstancesRequest) (<-chan *ListJobExecutionInstancesResponse, <-chan error) {
	responseChan := make(chan *ListJobExecutionInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobExecutionInstances(request)
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

// ListJobExecutionInstancesWithCallback invokes the emr.ListJobExecutionInstances API asynchronously
// api document: https://help.aliyun.com/api/emr/listjobexecutioninstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobExecutionInstancesWithCallback(request *ListJobExecutionInstancesRequest, callback func(response *ListJobExecutionInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobExecutionInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListJobExecutionInstances(request)
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

// ListJobExecutionInstancesRequest is the request struct for api ListJobExecutionInstances
type ListJobExecutionInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IsDesc                  requests.Boolean `position:"Query" name:"IsDesc"`
	PageNumber              requests.Integer `position:"Query" name:"PageNumber"`
	ExecutionPlanInstanceId string           `position:"Query" name:"ExecutionPlanInstanceId"`
	PageSize                requests.Integer `position:"Query" name:"PageSize"`
}

// ListJobExecutionInstancesResponse is the response struct for api ListJobExecutionInstances
type ListJobExecutionInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	JobInstances JobInstances `json:"JobInstances" xml:"JobInstances"`
}

// CreateListJobExecutionInstancesRequest creates a request to invoke ListJobExecutionInstances API
func CreateListJobExecutionInstancesRequest() (request *ListJobExecutionInstancesRequest) {
	request = &ListJobExecutionInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListJobExecutionInstances", "emr", "openAPI")
	return
}

// CreateListJobExecutionInstancesResponse creates a response to parse from ListJobExecutionInstances response
func CreateListJobExecutionInstancesResponse() (response *ListJobExecutionInstancesResponse) {
	response = &ListJobExecutionInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
