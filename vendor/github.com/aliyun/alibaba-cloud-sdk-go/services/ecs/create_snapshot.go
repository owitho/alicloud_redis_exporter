package ecs

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

// CreateSnapshot invokes the ecs.CreateSnapshot API synchronously
// api document: https://help.aliyun.com/api/ecs/createsnapshot.html
func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
	response = CreateCreateSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSnapshotWithChan invokes the ecs.CreateSnapshot API asynchronously
// api document: https://help.aliyun.com/api/ecs/createsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSnapshotWithChan(request *CreateSnapshotRequest) (<-chan *CreateSnapshotResponse, <-chan error) {
	responseChan := make(chan *CreateSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSnapshot(request)
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

// CreateSnapshotWithCallback invokes the ecs.CreateSnapshot API asynchronously
// api document: https://help.aliyun.com/api/ecs/createsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSnapshotWithCallback(request *CreateSnapshotRequest, callback func(response *CreateSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSnapshotResponse
		var err error
		defer close(result)
		response, err = client.CreateSnapshot(request)
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

// CreateSnapshotRequest is the request struct for api CreateSnapshot
type CreateSnapshotRequest struct {
	*requests.RpcRequest
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	SnapshotName         string           `position:"Query" name:"SnapshotName"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
	DiskId               string           `position:"Query" name:"DiskId"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
}

// CreateSnapshotResponse is the response struct for api CreateSnapshot
type CreateSnapshotResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SnapshotId string `json:"SnapshotId" xml:"SnapshotId"`
}

// CreateCreateSnapshotRequest creates a request to invoke CreateSnapshot API
func CreateCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateSnapshot", "ecs", "openAPI")
	return
}

// CreateCreateSnapshotResponse creates a response to parse from CreateSnapshot response
func CreateCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
