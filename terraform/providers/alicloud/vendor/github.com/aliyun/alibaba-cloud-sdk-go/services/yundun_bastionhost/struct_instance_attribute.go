package yundun_bastionhost

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

// InstanceAttribute is a nested struct in yundun_bastionhost response
type InstanceAttribute struct {
	RegionId               string   `json:"RegionId" xml:"RegionId"`
	InstanceId             string   `json:"InstanceId" xml:"InstanceId"`
	InstanceStatus         string   `json:"InstanceStatus" xml:"InstanceStatus"`
	VpcId                  string   `json:"VpcId" xml:"VpcId"`
	VswitchId              string   `json:"VswitchId" xml:"VswitchId"`
	IntranetIp             string   `json:"IntranetIp" xml:"IntranetIp"`
	InternetIp             string   `json:"InternetIp" xml:"InternetIp"`
	IntranetEndpoint       string   `json:"IntranetEndpoint" xml:"IntranetEndpoint"`
	InternetEndpoint       string   `json:"InternetEndpoint" xml:"InternetEndpoint"`
	NetworkType            string   `json:"NetworkType" xml:"NetworkType"`
	EcsStatus              string   `json:"EcsStatus" xml:"EcsStatus"`
	ImageVersionName       string   `json:"ImageVersionName" xml:"ImageVersionName"`
	Renewable              bool     `json:"Renewable" xml:"Renewable"`
	PlanUpgradeable        bool     `json:"PlanUpgradeable" xml:"PlanUpgradeable"`
	PlanUpgradeStatus      int      `json:"PlanUpgradeStatus" xml:"PlanUpgradeStatus"`
	Upgradeable            bool     `json:"Upgradeable" xml:"Upgradeable"`
	UpgradeStatus          int      `json:"UpgradeStatus" xml:"UpgradeStatus"`
	AccessType             int      `json:"AccessType" xml:"AccessType"`
	PublicAccessControl    int      `json:"PublicAccessControl" xml:"PublicAccessControl"`
	PublicNetworkAccess    bool     `json:"PublicNetworkAccess" xml:"PublicNetworkAccess"`
	EniInstanceId          string   `json:"EniInstanceId" xml:"EniInstanceId"`
	StartTime              int64    `json:"StartTime" xml:"StartTime"`
	ExpireTime             int64    `json:"ExpireTime" xml:"ExpireTime"`
	Description            string   `json:"Description" xml:"Description"`
	LicenseCode            string   `json:"LicenseCode" xml:"LicenseCode"`
	SeriesCode             string   `json:"SeriesCode" xml:"SeriesCode"`
	PublicWhiteList        []string `json:"PublicWhiteList" xml:"PublicWhiteList"`
	PrivateWhiteList       []string `json:"PrivateWhiteList" xml:"PrivateWhiteList"`
	SecurityGroupIds       []string `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	ReferredSecurityGroups []string `json:"ReferredSecurityGroups" xml:"ReferredSecurityGroups"`
}