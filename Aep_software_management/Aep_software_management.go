package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数id: 类型long, 参数不可以为空
//  描述:升级包id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，在产品概况中可以查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateSoftware(appKey string, appSecret string, id string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_software_management/software"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20200529232851"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数id: 类型long, 参数不可以为空
//  描述:升级包id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey,在产品概况中可以查询
func DeleteSoftware(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_software_management/software"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20200529232809"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数id: 类型long, 参数不可以为空
//  描述:升级包ID
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，可在产品概况中查看
func QuerySoftware(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_software_management/software"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20200529232806"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数searchValue: 类型String, 参数可以为空
//  描述:查询条件，可以根据升级包名称模糊查询
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，可以在产品概况中查看
func QuerySoftwareList(appKey string, appSecret string, productId string, MasterKey string, searchValue string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_software_management/softwares"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200529232801"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
