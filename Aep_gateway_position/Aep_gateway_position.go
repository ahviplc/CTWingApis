package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数cardId: 类型String, 参数不可以为空
//  描述:定位的物联网卡号
//参数posReqType: 类型long, 参数不可以为空
//  描述:返回的位置类型：1为初始位置；2为最新位置；3为最新or历史位置；7目前未用到，保留数字
func GetPosition(appKey string, appSecret string, cardId string, posReqType string) (*http.Response, error) {
	path := "/aep_gateway_position/api/getPosition"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["cardId"] = cardId
	param["posReqType"] = posReqType

	version := "20190301085737"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
