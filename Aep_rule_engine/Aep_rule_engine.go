package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func SaasCreateRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/createRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200111000503"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数ruleId: 类型String, 参数可以为空
//  描述:
//参数productId: 类型String, 参数可以为空
//  描述:
//参数pageNow: 类型long, 参数可以为空
//  描述:
//参数pageSize: 类型long, 参数可以为空
//  描述:
func SaasQueryRule(appKey string, appSecret string, ruleId string, productId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/queryRule"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["ruleId"] = ruleId
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200111000633"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func SaasUpdateRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/updateRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200111000540"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func SaasDeleteRuleEngine(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/deleteRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200111000611"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
