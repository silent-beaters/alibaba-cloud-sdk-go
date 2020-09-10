package cloudcallcenter

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

// ModifyVnSilenceTimeoutConfig invokes the cloudcallcenter.ModifyVnSilenceTimeoutConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnsilencetimeoutconfig.html
func (client *Client) ModifyVnSilenceTimeoutConfig(request *ModifyVnSilenceTimeoutConfigRequest) (response *ModifyVnSilenceTimeoutConfigResponse, err error) {
	response = CreateModifyVnSilenceTimeoutConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVnSilenceTimeoutConfigWithChan invokes the cloudcallcenter.ModifyVnSilenceTimeoutConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnsilencetimeoutconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnSilenceTimeoutConfigWithChan(request *ModifyVnSilenceTimeoutConfigRequest) (<-chan *ModifyVnSilenceTimeoutConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyVnSilenceTimeoutConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVnSilenceTimeoutConfig(request)
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

// ModifyVnSilenceTimeoutConfigWithCallback invokes the cloudcallcenter.ModifyVnSilenceTimeoutConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnsilencetimeoutconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnSilenceTimeoutConfigWithCallback(request *ModifyVnSilenceTimeoutConfigRequest, callback func(response *ModifyVnSilenceTimeoutConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVnSilenceTimeoutConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyVnSilenceTimeoutConfig(request)
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

// ModifyVnSilenceTimeoutConfigRequest is the request struct for api ModifyVnSilenceTimeoutConfig
type ModifyVnSilenceTimeoutConfigRequest struct {
	*requests.RpcRequest
	FinalAction       string           `position:"Query" name:"FinalAction"`
	FinalPrompt       string           `position:"Query" name:"FinalPrompt"`
	Threshold         requests.Integer `position:"Query" name:"Threshold"`
	IntentTrigger     string           `position:"Query" name:"IntentTrigger"`
	Timeout           requests.Integer `position:"Query" name:"Timeout"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	SourceType        string           `position:"Query" name:"SourceType"`
	FinalActionParams string           `position:"Query" name:"FinalActionParams"`
	Prompt            string           `position:"Query" name:"Prompt"`
}

// ModifyVnSilenceTimeoutConfigResponse is the response struct for api ModifyVnSilenceTimeoutConfig
type ModifyVnSilenceTimeoutConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVnSilenceTimeoutConfigRequest creates a request to invoke ModifyVnSilenceTimeoutConfig API
func CreateModifyVnSilenceTimeoutConfigRequest() (request *ModifyVnSilenceTimeoutConfigRequest) {
	request = &ModifyVnSilenceTimeoutConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyVnSilenceTimeoutConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyVnSilenceTimeoutConfigResponse creates a response to parse from ModifyVnSilenceTimeoutConfig response
func CreateModifyVnSilenceTimeoutConfigResponse() (response *ModifyVnSilenceTimeoutConfigResponse) {
	response = &ModifyVnSilenceTimeoutConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}