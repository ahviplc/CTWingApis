package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数id: 类型long, 参数不可以为空
//  描述:固件id
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateFirmware(appKey string, appSecret string, id string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_firmware_management/firmware"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20190615001705"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数searchValue: 类型String, 参数可以为空
//  描述:查询条件，可以根据固件名称模糊查询
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数
//参数MasterKey: 类型String, 参数可以为空
//  描述:
func QueryFirmwareList(appKey string, appSecret string, productId string, searchValue string, pageNow string, pageSize string, MasterKey string) (*http.Response, error) {
	path := "/aep_firmware_management/firmwares"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190615001608"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:固件id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
func QueryFirmware(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_firmware_management/firmware"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20190618151645"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:固件id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数updateBy: 类型String, 参数可以为空
//  描述:修改人
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
func DeleteFirmware(appKey string, appSecret string, id string, productId string, updateBy string, MasterKey string) (*http.Response, error) {
	path := "/aep_firmware_management/firmware"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId
	param["updateBy"] = updateBy
	param["MasterKey"] = MasterKey

	version := "20190615001534"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}
