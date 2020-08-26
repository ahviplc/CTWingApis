package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func QueryPublicByPublicProductId(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_public_product_management/publicProducts"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20190507003930"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func QueryPublicByProductId(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_public_product_management/publicProductList"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20190507004139"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func InstantiateProduct(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_public_product_management/instantiateProduct"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200801233037"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
