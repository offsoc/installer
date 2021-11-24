package edas

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

// ScaleOutApplication invokes the edas.ScaleOutApplication API synchronously
func (client *Client) ScaleOutApplication(request *ScaleOutApplicationRequest) (response *ScaleOutApplicationResponse, err error) {
	response = CreateScaleOutApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// ScaleOutApplicationWithChan invokes the edas.ScaleOutApplication API asynchronously
func (client *Client) ScaleOutApplicationWithChan(request *ScaleOutApplicationRequest) (<-chan *ScaleOutApplicationResponse, <-chan error) {
	responseChan := make(chan *ScaleOutApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScaleOutApplication(request)
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

// ScaleOutApplicationWithCallback invokes the edas.ScaleOutApplication API asynchronously
func (client *Client) ScaleOutApplicationWithCallback(request *ScaleOutApplicationRequest, callback func(response *ScaleOutApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScaleOutApplicationResponse
		var err error
		defer close(result)
		response, err = client.ScaleOutApplication(request)
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

// ScaleOutApplicationRequest is the request struct for api ScaleOutApplication
type ScaleOutApplicationRequest struct {
	*requests.RoaRequest
	EcuInfo     string `position:"Query" name:"EcuInfo"`
	DeployGroup string `position:"Query" name:"DeployGroup"`
	AppId       string `position:"Query" name:"AppId"`
}

// ScaleOutApplicationResponse is the response struct for api ScaleOutApplication
type ScaleOutApplicationResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateScaleOutApplicationRequest creates a request to invoke ScaleOutApplication API
func CreateScaleOutApplicationRequest() (request *ScaleOutApplicationRequest) {
	request = &ScaleOutApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ScaleOutApplication", "/pop/v5/changeorder/co_scale_out", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateScaleOutApplicationResponse creates a response to parse from ScaleOutApplication response
func CreateScaleOutApplicationResponse() (response *ScaleOutApplicationResponse) {
	response = &ScaleOutApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}