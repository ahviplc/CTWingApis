package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
func QueryRemoteUpgradeDetail(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/detail"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20190615001517"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
func QueryRemoteUpgradeTask(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20190615001509"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func ControlRemoteUpgradeTask(appKey string, appSecret string, id string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/control"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20190615001456"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数id: 类型String, 参数可以为空
//  描述:主任务id,isSelectDevice为1时必填，为2不必填
//参数productId: 类型String, 参数不可以为空
//  描述:产品id
//参数isSelectDevice: 类型String, 参数不可以为空
//  描述:查询类型（1.查询加入升级设备，2.查询可加入升级设备）
//参数pageNow: 类型String, 参数可以为空
//  描述:当前页，默认1
//参数pageSize: 类型String, 参数可以为空
//  描述:每页显示数，默认20
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
//参数deviceIdSearch: 类型String, 参数可以为空
//  描述:根据设备id精确查询
//参数deviceNameSearch: 类型String, 参数可以为空
//  描述:根据设备名称精确查询
//参数imeiSearch: 类型String, 参数可以为空
//  描述:根据imei号精确查询，仅支持LWM2M协议
//参数deviceNoSearch: 类型String, 参数可以为空
//  描述:根据设备编号精确查询，仅支持T_Link协议
//参数deviceGroupIdSearch: 类型String, 参数可以为空
//  描述:根据群组id精确查询
func QueryRemoteUpradeDeviceList(appKey string, appSecret string, productId string, isSelectDevice string, id string, pageNow string, pageSize string, MasterKey string, deviceIdSearch string, deviceNameSearch string, imeiSearch string, deviceNoSearch string, deviceGroupIdSearch string) (*http.Response, error) {
	path := "/aep_upgrade_management/devices"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId
	param["isSelectDevice"] = isSelectDevice
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize
	param["deviceIdSearch"] = deviceIdSearch
	param["deviceNameSearch"] = deviceNameSearch
	param["imeiSearch"] = imeiSearch
	param["deviceNoSearch"] = deviceNoSearch
	param["deviceGroupIdSearch"] = deviceGroupIdSearch

	version := "20190615001451"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数updateBy: 类型String, 参数可以为空
//  描述:修改人
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
func DeleteRemoteUpgradeTask(appKey string, appSecret string, id string, productId string, updateBy string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId
	param["updateBy"] = updateBy

	version := "20190615001444"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数，默认1
//参数pageSize: 类型long, 参数可以为空
//  描述:每页显示数，默认20
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
//参数searchValue: 类型String, 参数可以为空
//  描述:查询条件，支持主任务名称模糊查询
func QueryRemoteUpgradeTaskList(appKey string, appSecret string, productId string, pageNow string, pageSize string, MasterKey string, searchValue string) (*http.Response, error) {
	path := "/aep_upgrade_management/tasks"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize
	param["searchValue"] = searchValue

	version := "20190615001440"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数MasterKey: 类型String, 参数可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func ModifyRemoteUpgradeTask(appKey string, appSecret string, id string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20190615001433"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateRemoteUpgradeTask(appKey string, appSecret string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615001416"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func OperationalRemoteUpgradeTask(appKey string, appSecret string, body string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/operational"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615001412"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数taskStatus: 类型long, 参数可以为空
//  描述:子任务状态
//  T-Link协议必填（1.待升级，2.升级中，3.升级成功，4.升级失败）
//  LWM2M协议选填（0:升级可行性判断,1:升级可行性判断失败,2:分派升级任务,3:分派升级任务失败,4:分派下载任务,5:分派下载任务失败,6:开始升级,7:升级失败,8:升级完成,9:取消当前设备的升级,10:取消当前设备升级成功,11:取消当前设备升级失败）
//参数searchValue: 类型String, 参数可以为空
//  描述:查询值，设备ID或设备编号(IMEI)或设备名称模糊查询
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页码
//参数pageSize: 类型long, 参数可以为空
//  描述:每页显示数
//参数MasterKey: 类型String, 参数可以为空
//  描述:MasterKey
func QueryRemoteUpgradeSubtasks(appKey string, appSecret string, id string, productId string, taskStatus string, searchValue string, pageNow string, pageSize string, MasterKey string) (*http.Response, error) {
	path := "/aep_upgrade_management/details"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId
	param["taskStatus"] = taskStatus
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20190615001406"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
