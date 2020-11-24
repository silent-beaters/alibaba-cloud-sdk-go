package hitsdb

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

// ExploreHiTSDBInstanceDataMetricList invokes the hitsdb.ExploreHiTSDBInstanceDataMetricList API synchronously
func (client *Client) ExploreHiTSDBInstanceDataMetricList(request *ExploreHiTSDBInstanceDataMetricListRequest) (response *ExploreHiTSDBInstanceDataMetricListResponse, err error) {
	response = CreateExploreHiTSDBInstanceDataMetricListResponse()
	err = client.DoAction(request, response)
	return
}

// ExploreHiTSDBInstanceDataMetricListWithChan invokes the hitsdb.ExploreHiTSDBInstanceDataMetricList API asynchronously
func (client *Client) ExploreHiTSDBInstanceDataMetricListWithChan(request *ExploreHiTSDBInstanceDataMetricListRequest) (<-chan *ExploreHiTSDBInstanceDataMetricListResponse, <-chan error) {
	responseChan := make(chan *ExploreHiTSDBInstanceDataMetricListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExploreHiTSDBInstanceDataMetricList(request)
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

// ExploreHiTSDBInstanceDataMetricListWithCallback invokes the hitsdb.ExploreHiTSDBInstanceDataMetricList API asynchronously
func (client *Client) ExploreHiTSDBInstanceDataMetricListWithCallback(request *ExploreHiTSDBInstanceDataMetricListRequest, callback func(response *ExploreHiTSDBInstanceDataMetricListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExploreHiTSDBInstanceDataMetricListResponse
		var err error
		defer close(result)
		response, err = client.ExploreHiTSDBInstanceDataMetricList(request)
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

// ExploreHiTSDBInstanceDataMetricListRequest is the request struct for api ExploreHiTSDBInstanceDataMetricList
type ExploreHiTSDBInstanceDataMetricListRequest struct {
	*requests.RpcRequest
	ReverseVpcIp         string           `position:"Query" name:"ReverseVpcIp"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Prefix               string           `position:"Query" name:"Prefix"`
	ReverseVpcPort       requests.Integer `position:"Query" name:"ReverseVpcPort"`
	PassWord             string           `position:"Query" name:"PassWord"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Max                  requests.Integer `position:"Query" name:"Max"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// ExploreHiTSDBInstanceDataMetricListResponse is the response struct for api ExploreHiTSDBInstanceDataMetricList
type ExploreHiTSDBInstanceDataMetricListResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	MetricList []Data `json:"MetricList" xml:"MetricList"`
}

// CreateExploreHiTSDBInstanceDataMetricListRequest creates a request to invoke ExploreHiTSDBInstanceDataMetricList API
func CreateExploreHiTSDBInstanceDataMetricListRequest() (request *ExploreHiTSDBInstanceDataMetricListRequest) {
	request = &ExploreHiTSDBInstanceDataMetricListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2017-06-01", "ExploreHiTSDBInstanceDataMetricList", "tsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExploreHiTSDBInstanceDataMetricListResponse creates a response to parse from ExploreHiTSDBInstanceDataMetricList response
func CreateExploreHiTSDBInstanceDataMetricListResponse() (response *ExploreHiTSDBInstanceDataMetricListResponse) {
	response = &ExploreHiTSDBInstanceDataMetricListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
