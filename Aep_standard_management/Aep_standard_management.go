package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数standardVersion: 类型String, 参数可以为空
//  描述:标准物模型版本号
//参数thirdType: 类型long, 参数不可以为空
//  描述:三级分类Id
func QueryStandardModel(appKey string, appSecret string, thirdType string, standardVersion string) (*http.Response, error) {
	path := "/aep_standard_management/standardModel"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["standardVersion"] = standardVersion
	param["thirdType"] = thirdType

	version := "20190713033424"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
