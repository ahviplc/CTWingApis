package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型String, 参数不可以为空
//  描述:
//参数searchValue: 类型String, 参数可以为空
//  描述:
//参数deviceId: 类型String, 参数可以为空
//  描述:
//参数status: 类型String, 参数可以为空
//  描述:状态可选填： 1：指令已保存 2：指令已发送 3：指令已送达 4：指令已完成 6：指令已取消 999：指令失败
//参数startTime: 类型String, 参数可以为空
//  描述:
//参数endTime: 类型String, 参数可以为空
//  描述:
//参数pageNow: 类型String, 参数可以为空
//  描述:
//参数pageSize: 类型String, 参数可以为空
//  描述:
func QueryCommandList(appKey string, appSecret string, MasterKey string, productId string, searchValue string, deviceId string, status string, startTime string, endTime string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_command_modbus/modbus/commands"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["deviceId"] = deviceId
	param["status"] = status
	param["startTime"] = startTime
	param["endTime"] = endTime
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200404012458"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数commandId: 类型String, 参数不可以为空
//  描述:
//参数productId: 类型long, 参数不可以为空
//  描述:
func QueryCommand(appKey string, appSecret string, MasterKey string, commandId string, productId string) (*http.Response, error) {
	path := "/aep_command_modbus/modbus/command"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["commandId"] = commandId
	param["productId"] = productId

	version := "20200404012456"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CancelCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_command_modbus/modbus/cancelCommand"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20200404012453"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_command_modbus/modbus/command"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20200404012449"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
