package alidns

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

// DescribeGtmInstance invokes the alidns.DescribeGtmInstance API synchronously
// api document: https://help.aliyun.com/api/alidns/describegtminstance.html
func (client *Client) DescribeGtmInstance(request *DescribeGtmInstanceRequest) (response *DescribeGtmInstanceResponse, err error) {
	response = CreateDescribeGtmInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmInstanceWithChan invokes the alidns.DescribeGtmInstance API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtminstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmInstanceWithChan(request *DescribeGtmInstanceRequest) (<-chan *DescribeGtmInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmInstance(request)
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

// DescribeGtmInstanceWithCallback invokes the alidns.DescribeGtmInstance API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtminstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmInstanceWithCallback(request *DescribeGtmInstanceRequest, callback func(response *DescribeGtmInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmInstance(request)
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

// DescribeGtmInstanceRequest is the request struct for api DescribeGtmInstance
type DescribeGtmInstanceRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeGtmInstanceResponse is the response struct for api DescribeGtmInstance
type DescribeGtmInstanceResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	InstanceId      string `json:"InstanceId" xml:"InstanceId"`
	InstanceName    string `json:"InstanceName" xml:"InstanceName"`
	VersionCode     string `json:"VersionCode" xml:"VersionCode"`
	ExpireTime      string `json:"ExpireTime" xml:"ExpireTime"`
	ExpireTimestamp int64  `json:"ExpireTimestamp" xml:"ExpireTimestamp"`
	Cname           string `json:"Cname" xml:"Cname"`
	UserDomainName  string `json:"UserDomainName" xml:"UserDomainName"`
	Ttl             int    `json:"Ttl" xml:"Ttl"`
	LbaStrategy     string `json:"LbaStrategy" xml:"LbaStrategy"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	CreateTimestamp int64  `json:"CreateTimestamp" xml:"CreateTimestamp"`
	AlertGroup      string `json:"AlertGroup" xml:"AlertGroup"`
	CnameMode       string `json:"CnameMode" xml:"CnameMode"`
}

// CreateDescribeGtmInstanceRequest creates a request to invoke DescribeGtmInstance API
func CreateDescribeGtmInstanceRequest() (request *DescribeGtmInstanceRequest) {
	request = &DescribeGtmInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmInstance", "Alidns", "openAPI")
	return
}

// CreateDescribeGtmInstanceResponse creates a response to parse from DescribeGtmInstance response
func CreateDescribeGtmInstanceResponse() (response *DescribeGtmInstanceResponse) {
	response = &DescribeGtmInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
