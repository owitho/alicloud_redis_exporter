package rds

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

// DescribeReplicaConflictInfo invokes the rds.DescribeReplicaConflictInfo API synchronously
// api document: https://help.aliyun.com/api/rds/describereplicaconflictinfo.html
func (client *Client) DescribeReplicaConflictInfo(request *DescribeReplicaConflictInfoRequest) (response *DescribeReplicaConflictInfoResponse, err error) {
	response = CreateDescribeReplicaConflictInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReplicaConflictInfoWithChan invokes the rds.DescribeReplicaConflictInfo API asynchronously
// api document: https://help.aliyun.com/api/rds/describereplicaconflictinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReplicaConflictInfoWithChan(request *DescribeReplicaConflictInfoRequest) (<-chan *DescribeReplicaConflictInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeReplicaConflictInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReplicaConflictInfo(request)
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

// DescribeReplicaConflictInfoWithCallback invokes the rds.DescribeReplicaConflictInfo API asynchronously
// api document: https://help.aliyun.com/api/rds/describereplicaconflictinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReplicaConflictInfoWithCallback(request *DescribeReplicaConflictInfoRequest, callback func(response *DescribeReplicaConflictInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReplicaConflictInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeReplicaConflictInfo(request)
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

// DescribeReplicaConflictInfoRequest is the request struct for api DescribeReplicaConflictInfo
type DescribeReplicaConflictInfoRequest struct {
	*requests.RpcRequest
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeReplicaConflictInfoResponse is the response struct for api DescribeReplicaConflictInfo
type DescribeReplicaConflictInfoResponse struct {
	*responses.BaseResponse
	RequestId        string      `json:"RequestId" xml:"RequestId"`
	ReplicaId        string      `json:"ReplicaId" xml:"ReplicaId"`
	PagNumber        int         `json:"PagNumber" xml:"PagNumber"`
	PageRecordCount  int         `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeReplicaConflictInfoRequest creates a request to invoke DescribeReplicaConflictInfo API
func CreateDescribeReplicaConflictInfoRequest() (request *DescribeReplicaConflictInfoRequest) {
	request = &DescribeReplicaConflictInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicaConflictInfo", "rds", "openAPI")
	return
}

// CreateDescribeReplicaConflictInfoResponse creates a response to parse from DescribeReplicaConflictInfo response
func CreateDescribeReplicaConflictInfoResponse() (response *DescribeReplicaConflictInfoResponse) {
	response = &DescribeReplicaConflictInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
