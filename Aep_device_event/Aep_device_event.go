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
//  描述:可填值：设备Id
//参数eventType: 类型long, 参数可以为空
//  描述:LWM2M,MQTT,TUP协议可选填:
//  1：信息
//  2：警告
//  3：故障
//  T-link协议可选填:
//  0:普通事件
//  1：告警事件(普通级)
//  2：告警事件(重大级)
//  3：告警事件(严重级)
//参数startTime: 类型String, 参数可以为空
//  描述:精确到毫秒的时间戳
//参数endTime: 类型String, 参数可以为空
//  描述:精确到毫秒的时间戳
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数
//参数pageSize: 类型long, 参数可以为空
//  描述:每页记录数
func QueryEventList(appKey string, appSecret string, MasterKey string, productId string, searchValue string, eventType string, startTime string, endTime string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_device_event/events"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["searchValue"] = searchValue
	param["eventType"] = eventType
	param["startTime"] = startTime
	param["endTime"] = endTime
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190507012824"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
