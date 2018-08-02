package drds

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

// ModifyDrdsInstanceDescription invokes the drds.ModifyDrdsInstanceDescription API synchronously
// api document: https://help.aliyun.com/api/drds/modifydrdsinstancedescription.html
func (client *Client) ModifyDrdsInstanceDescription(request *ModifyDrdsInstanceDescriptionRequest) (response *ModifyDrdsInstanceDescriptionResponse, err error) {
	response = CreateModifyDrdsInstanceDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDrdsInstanceDescriptionWithChan invokes the drds.ModifyDrdsInstanceDescription API asynchronously
// api document: https://help.aliyun.com/api/drds/modifydrdsinstancedescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDrdsInstanceDescriptionWithChan(request *ModifyDrdsInstanceDescriptionRequest) (<-chan *ModifyDrdsInstanceDescriptionResponse, <-chan error) {
	responseChan := make(chan *ModifyDrdsInstanceDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDrdsInstanceDescription(request)
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

// ModifyDrdsInstanceDescriptionWithCallback invokes the drds.ModifyDrdsInstanceDescription API asynchronously
// api document: https://help.aliyun.com/api/drds/modifydrdsinstancedescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDrdsInstanceDescriptionWithCallback(request *ModifyDrdsInstanceDescriptionRequest, callback func(response *ModifyDrdsInstanceDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDrdsInstanceDescriptionResponse
		var err error
		defer close(result)
		response, err = client.ModifyDrdsInstanceDescription(request)
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

// ModifyDrdsInstanceDescriptionRequest is the request struct for api ModifyDrdsInstanceDescription
type ModifyDrdsInstanceDescriptionRequest struct {
	*requests.RpcRequest
	Description    string `position:"Query" name:"Description"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// ModifyDrdsInstanceDescriptionResponse is the response struct for api ModifyDrdsInstanceDescription
type ModifyDrdsInstanceDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyDrdsInstanceDescriptionRequest creates a request to invoke ModifyDrdsInstanceDescription API
func CreateModifyDrdsInstanceDescriptionRequest() (request *ModifyDrdsInstanceDescriptionRequest) {
	request = &ModifyDrdsInstanceDescriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsInstanceDescription", "", "")
	return
}

// CreateModifyDrdsInstanceDescriptionResponse creates a response to parse from ModifyDrdsInstanceDescription response
func CreateModifyDrdsInstanceDescriptionResponse() (response *ModifyDrdsInstanceDescriptionResponse) {
	response = &ModifyDrdsInstanceDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
