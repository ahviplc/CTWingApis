package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数subId: 类型long, 参数不可以为空
//  描述:订阅记录id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:产品MasterKey
func GetSubscription(appKey string, appSecret string, subId string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_subscribe_north/subscription"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["subId"] = subId
	param["productId"] = productId

	version := "20181031202033"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数productId: 类型long, 参数不可以为空
//  描述:产品ID
//参数pageNow: 类型long, 参数不可以为空
//  描述:当前页
//参数pageSize: 类型long, 参数不可以为空
//  描述:每页条数
//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数subType: 类型long, 参数可以为空
//  描述:订阅类型
//参数searchValue: 类型String, 参数可以为空
//  描述:检索deviceId,模糊匹配
func GetSubscriptionsList(appKey string, appSecret string, productId string, pageNow string, pageSize string, MasterKey string, subType string, searchValue string) (*http.Response, error) {
	path := "/aep_subscribe_north/subscribes"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize
	param["subType"] = subType
	param["searchValue"] = searchValue

	version := "20181031202027"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数subId: 类型String, 参数不可以为空
//  描述:订阅记录id
//参数deviceId: 类型String, 参数可以为空
//  描述:设备id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数subLevel: 类型long, 参数不可以为空
//  描述:订阅级别
//参数MasterKey: 类型String, 参数不可以为空
//  描述:产品MasterKey
func DeleteSubscription(appKey string, appSecret string, subId string, productId string, subLevel string, MasterKey string, deviceId string) (*http.Response, error) {
	path := "/aep_subscribe_north/subscription"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["subId"] = subId
	param["deviceId"] = deviceId
	param["productId"] = productId
	param["subLevel"] = subLevel

	version := "20181031202023"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateSubscription(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_subscribe_north/subscription"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20181031202018"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
