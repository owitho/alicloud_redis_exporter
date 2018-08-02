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

// RemoveReadOnlyAccount invokes the drds.RemoveReadOnlyAccount API synchronously
// api document: https://help.aliyun.com/api/drds/removereadonlyaccount.html
func (client *Client) RemoveReadOnlyAccount(request *RemoveReadOnlyAccountRequest) (response *RemoveReadOnlyAccountResponse, err error) {
	response = CreateRemoveReadOnlyAccountResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveReadOnlyAccountWithChan invokes the drds.RemoveReadOnlyAccount API asynchronously
// api document: https://help.aliyun.com/api/drds/removereadonlyaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveReadOnlyAccountWithChan(request *RemoveReadOnlyAccountRequest) (<-chan *RemoveReadOnlyAccountResponse, <-chan error) {
	responseChan := make(chan *RemoveReadOnlyAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveReadOnlyAccount(request)
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

// RemoveReadOnlyAccountWithCallback invokes the drds.RemoveReadOnlyAccount API asynchronously
// api document: https://help.aliyun.com/api/drds/removereadonlyaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveReadOnlyAccountWithCallback(request *RemoveReadOnlyAccountRequest, callback func(response *RemoveReadOnlyAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveReadOnlyAccountResponse
		var err error
		defer close(result)
		response, err = client.RemoveReadOnlyAccount(request)
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

// RemoveReadOnlyAccountRequest is the request struct for api RemoveReadOnlyAccount
type RemoveReadOnlyAccountRequest struct {
	*requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// RemoveReadOnlyAccountResponse is the response struct for api RemoveReadOnlyAccount
type RemoveReadOnlyAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRemoveReadOnlyAccountRequest creates a request to invoke RemoveReadOnlyAccount API
func CreateRemoveReadOnlyAccountRequest() (request *RemoveReadOnlyAccountRequest) {
	request = &RemoveReadOnlyAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "RemoveReadOnlyAccount", "", "")
	return
}

// CreateRemoveReadOnlyAccountResponse creates a response to parse from RemoveReadOnlyAccount response
func CreateRemoveReadOnlyAccountResponse() (response *RemoveReadOnlyAccountResponse) {
	response = &RemoveReadOnlyAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
