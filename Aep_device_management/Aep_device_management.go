package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数searchValue: 类型String, 参数可以为空
//  描述:T-link协议可选填:设备名称，设备编号，设备Id
//  MQTT协议可选填:设备名称，设备编号，设备Id
//  LWM2M协议可选填:设备名称，设备Id ,IMEI号
//  TUP协议可选填:设备名称，设备Id ,IMEI号
//  TCP协议可选填:设备名称，设备编号，设备Id
//  HTTP协议可选填:设备名称，设备编号，设备Id
//  JT/T808协议可选填:设备名称，设备编号，设备Id
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数,最大100
func QueryDeviceList(appKey string, appSecret string, MasterKey string, productId string, searchValue string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_device_management/devices"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190507012134"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数deviceId: 类型String, 参数不可以为空
//  描述:
//参数productId: 类型long, 参数不可以为空
//  描述:
func QueryDevice(appKey string, appSecret string, MasterKey string, deviceId string, productId string) (*http.Response, error) {
	path := "/aep_device_management/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["deviceId"] = deviceId
	param["productId"] = productId

	version := "20181031202139"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数deviceIds: 类型String, 参数不可以为空
//  描述:可以删除多个设备（最多支持200个设备）。多个设备id，中间以逗号 "," 隔开 。样例：05979394b88a45b0842de729c03d99af,06106b8e1d5a458399326e003bcf05b4
func DeleteDevice(appKey string, appSecret string, MasterKey string, productId string, deviceIds string) (*http.Response, error) {
	path := "/aep_device_management/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["deviceIds"] = deviceIds

	version := "20181031202131"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数deviceId: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateDevice(appKey string, appSecret string, MasterKey string, deviceId string, body string) (*http.Response, error) {
	path := "/aep_device_management/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["deviceId"] = deviceId

	version := "20181031202122"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateDevice(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_management/device"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20181031202117"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func BindDevice(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_management/bindDevice"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20191024140057"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UnbindDevice(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_management/unbindDevice"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20191024140103"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数imei: 类型String, 参数不可以为空
//  描述:
func QueryProductInfoByImei(appKey string, appSecret string, imei string) (*http.Response, error) {
	path := "/aep_device_management/device/getProductInfoFormApiByImei"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["imei"] = imei

	version := "20191213161859"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
