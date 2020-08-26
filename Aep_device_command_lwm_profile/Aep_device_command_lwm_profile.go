package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey在该设备所属产品的概况中可以查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateCommandLwm2mProfile(appKey string, appSecret string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_device_command_lwm_profile/commandLwm2mProfile"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20191231141545"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
