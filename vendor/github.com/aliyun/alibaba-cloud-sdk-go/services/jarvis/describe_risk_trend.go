package jarvis

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

// DescribeRiskTrend invokes the jarvis.DescribeRiskTrend API synchronously
// api document: https://help.aliyun.com/api/jarvis/describerisktrend.html
func (client *Client) DescribeRiskTrend(request *DescribeRiskTrendRequest) (response *DescribeRiskTrendResponse, err error) {
	response = CreateDescribeRiskTrendResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskTrendWithChan invokes the jarvis.DescribeRiskTrend API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describerisktrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskTrendWithChan(request *DescribeRiskTrendRequest) (<-chan *DescribeRiskTrendResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskTrend(request)
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

// DescribeRiskTrendWithCallback invokes the jarvis.DescribeRiskTrend API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describerisktrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskTrendWithCallback(request *DescribeRiskTrendRequest, callback func(response *DescribeRiskTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskTrendResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskTrend(request)
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

// DescribeRiskTrendRequest is the request struct for api DescribeRiskTrend
type DescribeRiskTrendRequest struct {
	*requests.RpcRequest
	SourceIp      string `position:"Query" name:"SourceIp"`
	QueryProduct  string `position:"Query" name:"QueryProduct"`
	Lang          string `position:"Query" name:"Lang"`
	Peroid        string `position:"Query" name:"Peroid"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	QueryRegionId string `position:"Query" name:"QueryRegionId"`
}

// DescribeRiskTrendResponse is the response struct for api DescribeRiskTrend
type DescribeRiskTrendResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalCount string     `json:"TotalCount" xml:"TotalCount"`
	Module     string     `json:"Module" xml:"Module"`
	DataList   []DataItem `json:"DataList" xml:"DataList"`
}

// CreateDescribeRiskTrendRequest creates a request to invoke DescribeRiskTrend API
func CreateDescribeRiskTrendRequest() (request *DescribeRiskTrendRequest) {
	request = &DescribeRiskTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeRiskTrend", "", "")
	return
}

// CreateDescribeRiskTrendResponse creates a response to parse from DescribeRiskTrend response
func CreateDescribeRiskTrendResponse() (response *DescribeRiskTrendResponse) {
	response = &DescribeRiskTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
