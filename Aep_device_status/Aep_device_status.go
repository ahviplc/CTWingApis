package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func QueryDeviceStatus(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_device_status/deviceStatus"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20181031202028"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func QueryDeviceStatusList(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_device_status/deviceStatusList"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20181031202403"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func GetDeviceStatusHisInTotal(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_device_status/api/v1/getDeviceStatusHisInTotal"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20190928013529"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func GetDeviceStatusHisInPage(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_device_status/getDeviceStatusHisInPage"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20190928013337"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
