package cr

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

// DeleteCollection invokes the cr.DeleteCollection API synchronously
// api document: https://help.aliyun.com/api/cr/deletecollection.html
func (client *Client) DeleteCollection(request *DeleteCollectionRequest) (response *DeleteCollectionResponse, err error) {
	response = CreateDeleteCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCollectionWithChan invokes the cr.DeleteCollection API asynchronously
// api document: https://help.aliyun.com/api/cr/deletecollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCollectionWithChan(request *DeleteCollectionRequest) (<-chan *DeleteCollectionResponse, <-chan error) {
	responseChan := make(chan *DeleteCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCollection(request)
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

// DeleteCollectionWithCallback invokes the cr.DeleteCollection API asynchronously
// api document: https://help.aliyun.com/api/cr/deletecollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCollectionWithCallback(request *DeleteCollectionRequest, callback func(response *DeleteCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCollectionResponse
		var err error
		defer close(result)
		response, err = client.DeleteCollection(request)
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

// DeleteCollectionRequest is the request struct for api DeleteCollection
type DeleteCollectionRequest struct {
	*requests.RoaRequest
	CollectionId requests.Integer `position:"Path" name:"CollectionId"`
}

// DeleteCollectionResponse is the response struct for api DeleteCollection
type DeleteCollectionResponse struct {
	*responses.BaseResponse
}

// CreateDeleteCollectionRequest creates a request to invoke DeleteCollection API
func CreateDeleteCollectionRequest() (request *DeleteCollectionRequest) {
	request = &DeleteCollectionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "DeleteCollection", "/collections/[CollectionId]", "cr", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteCollectionResponse creates a response to parse from DeleteCollection response
func CreateDeleteCollectionResponse() (response *DeleteCollectionResponse) {
	response = &DeleteCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
