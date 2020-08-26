package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数sdk_type: 类型String, 参数可以为空
//  描述:SDK语言类型，默认为Java(字典项: sdk_type)
//参数file_name: 类型String, 参数不可以为空
//  描述:SDK描述，用以标识其中的biz包
//参数application_id: 类型String, 参数不可以为空
//  描述:应用编码，下载的SDK会根据该编码收集所有有权限的API打包
//参数api_version: 类型String, 参数可以为空
//  描述:API版本信息 TODO
func SdkDownload(appKey string, appSecret string, file_name string, application_id string, sdk_type string, api_version string) (*http.Response, error) {
	path := "/usr/sdk/download"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["sdk_type"] = sdk_type
	param["file_name"] = file_name
	param["application_id"] = application_id
	param["api_version"] = api_version

	version := "20180000000000"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
