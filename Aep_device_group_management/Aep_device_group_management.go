package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数MasterKey: 类型String, 参数可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateDeviceGroup(appKey string, appSecret string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_device_group_management/deviceGroup"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615001622"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数MasterKey: 类型String, 参数可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateDeviceGroup(appKey string, appSecret string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_device_group_management/deviceGroup"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615001615"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数productId: 类型long, 参数不可以为空
//  描述:产品Id
//参数deviceGroupId: 类型long, 参数不可以为空
//  描述:分组Id
//参数MasterKey: 类型String, 参数可以为空
//  描述:
func DeleteDeviceGroup(appKey string, appSecret string, productId string, deviceGroupId string, MasterKey string) (*http.Response, error) {
	path := "/aep_device_group_management/deviceGroup"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["deviceGroupId"] = deviceGroupId

	version := "20190615001601"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数productId: 类型long, 参数不可以为空
//  描述:
//参数searchValue: 类型String, 参数可以为空
//  描述:分组名称，分组ID
//参数pageNow: 类型long, 参数不可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数不可以为空
//  描述:每页记录数
//参数MasterKey: 类型String, 参数可以为空
//  描述:
func QueryDeviceGroupList(appKey string, appSecret string, productId string, pageNow string, pageSize string, searchValue string, MasterKey string) (*http.Response, error) {
	path := "/aep_device_group_management/deviceGroups"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190615001555"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数可以为空
//  描述:
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数searchValue: 类型String, 参数可以为空
//  描述:可查询：设备ID，设备名称，设备编号或者IMEI号
//参数pageNow: 类型long, 参数不可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数不可以为空
//  描述:每页条数
//参数deviceGroupId: 类型long, 参数可以为空
//  描述:群组ID：1.有值则查询该群组已关联的设备信息列表。2.为空则查询该产品下未关联的设备信息列表
func QueryGroupDeviceList(appKey string, appSecret string, productId string, pageNow string, pageSize string, MasterKey string, searchValue string, deviceGroupId string) (*http.Response, error) {
	path := "/aep_device_group_management/groupDeviceList"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize
	param["deviceGroupId"] = deviceGroupId

	version := "20190615001540"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateDeviceGroupRelation(appKey string, appSecret string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_device_group_management/deviceGroupRelation"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615001526"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}
