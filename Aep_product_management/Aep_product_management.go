package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数productId: 类型long, 参数不可以为空
//  描述:
func QueryProduct(appKey string, appSecret string, productId string) (*http.Response, error) {
	path := "/aep_product_management/product"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["productId"] = productId

	version := "20181031202055"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数searchValue: 类型String, 参数可以为空
//  描述:产品id或者产品名称
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数，最大40
func QueryProductList(appKey string, appSecret string, searchValue string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_product_management/products"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190507004824"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数productId: 类型long, 参数不可以为空
//  描述:
func DeleteProduct(appKey string, appSecret string, MasterKey string, productId string) (*http.Response, error) {
	path := "/aep_product_management/product"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId

	version := "20181031202029"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateProduct(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_product_management/product"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20191018204154"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func UpdateProduct(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_product_management/product"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20191018204806"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}
