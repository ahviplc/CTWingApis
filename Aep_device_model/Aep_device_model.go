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
//  描述:可填值：属性名称，属性标识符
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数
func QueryPropertyList(appKey string, appSecret string, MasterKey string, productId string, searchValue string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_device_model/dm/app/model/properties"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190712223330"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数searchValue: 类型String, 参数可以为空
//  描述:可填： 服务Id, 服务名称,服务标识符
//参数serviceType: 类型long, 参数可以为空
//  描述:服务类型
//  1. 数据上报
//  2. 事件上报
//  3.数据获取
//  4.参数查询
//  5.参数配置
//  6.指令下发
//  7.指令下发响应
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数
func QueryServiceList(appKey string, appSecret string, MasterKey string, productId string, searchValue string, serviceType string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_device_model/dm/app/model/services"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["serviceType"] = serviceType
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190712224233"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
