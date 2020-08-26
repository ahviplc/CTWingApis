package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数deviceId: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateDevice(appKey string, appSecret string, MasterKey string, deviceId string, body string) (*http.Response, error) {
	path := "/aep_modbus_device_management/modbus/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["deviceId"] = deviceId

	version := "20200404012440"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateDevice(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_modbus_device_management/modbus/device"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200404012437"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数deviceId: 类型String, 参数不可以为空
//  描述:
//参数productId: 类型long, 参数不可以为空
//  描述:
func QueryDevice(appKey string, appSecret string, MasterKey string, deviceId string, productId string) (*http.Response, error) {
	path := "/aep_modbus_device_management/modbus/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["deviceId"] = deviceId
	param["productId"] = productId

	version := "20200404012435"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数searchValue: 类型String, 参数可以为空
//  描述:设备名称，设备编号，设备Id
//参数pageNow: 类型long, 参数可以为空
//  描述:
//参数pageSize: 类型long, 参数可以为空
//  描述:
func QueryDeviceList(appKey string, appSecret string, MasterKey string, productId string, searchValue string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_modbus_device_management/modbus/devices"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200404012428"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数deviceIds: 类型String, 参数不可以为空
//  描述:
func DeleteDevice(appKey string, appSecret string, MasterKey string, productId string, deviceIds string) (*http.Response, error) {
	path := "/aep_modbus_device_management/modbus/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["deviceIds"] = deviceIds

	version := "20200404012425"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}
