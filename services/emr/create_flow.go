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

// CreateFlow invokes the emr.CreateFlow API synchronously
// api document: https://help.aliyun.com/api/emr/createflow.html
func (client *Client) CreateFlow(request *CreateFlowRequest) (response *CreateFlowResponse, err error) {
	response = CreateCreateFlowResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFlowWithChan invokes the emr.CreateFlow API asynchronously
// api document: https://help.aliyun.com/api/emr/createflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFlowWithChan(request *CreateFlowRequest) (<-chan *CreateFlowResponse, <-chan error) {
	responseChan := make(chan *CreateFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFlow(request)
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

// CreateFlowWithCallback invokes the emr.CreateFlow API asynchronously
// api document: https://help.aliyun.com/api/emr/createflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFlowWithCallback(request *CreateFlowRequest, callback func(response *CreateFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFlowResponse
		var err error
		defer close(result)
		response, err = client.CreateFlow(request)
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

// CreateFlowRequest is the request struct for api CreateFlow
type CreateFlowRequest struct {
	*requests.RpcRequest
	CronExpr                string           `position:"Query" name:"CronExpr"`
	Description             string           `position:"Query" name:"Description"`
	AlertUserGroupBizId     string           `position:"Query" name:"AlertUserGroupBizId"`
	HostName                string           `position:"Query" name:"HostName"`
	CreateCluster           requests.Boolean `position:"Query" name:"CreateCluster"`
	EndSchedule             requests.Integer `position:"Query" name:"EndSchedule"`
	AlertConf               string           `position:"Query" name:"AlertConf"`
	ProjectId               string           `position:"Query" name:"ProjectId"`
	ParentFlowList          string           `position:"Query" name:"ParentFlowList"`
	AlertDingDingGroupBizId string           `position:"Query" name:"AlertDingDingGroupBizId"`
	StartSchedule           requests.Integer `position:"Query" name:"StartSchedule"`
	ClusterId               string           `position:"Query" name:"ClusterId"`
	Application             string           `position:"Query" name:"Application"`
	Name                    string           `position:"Query" name:"Name"`
	ParentCategory          string           `position:"Query" name:"ParentCategory"`
}

// CreateFlowResponse is the response struct for api CreateFlow
type CreateFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateCreateFlowRequest creates a request to invoke CreateFlow API
func CreateCreateFlowRequest() (request *CreateFlowRequest) {
	request = &CreateFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateFlow", "emr", "openAPI")
	return
}

// CreateCreateFlowResponse creates a response to parse from CreateFlow response
func CreateCreateFlowResponse() (response *CreateFlowResponse) {
	response = &CreateFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
