package aepapi

import (
	"CTWingApis/core"
	"net/http"
)

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，在产品概况中查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func OperationalSoftwareUpgradeTask(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/operational"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20200529233236"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数taskStatus: 类型long, 参数可以为空
//  描述:子任务状态：
//  0.待升级，1.查询设备版本号，2.新版本通知，3.请求升级包，4.设备上报下载状态，5.执行升级，6.通知升级结果
//参数searchValue: 类型String, 参数可以为空
//  描述:查询值，设备ID或设备编号(IMEI)或设备名称模糊查询
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页码
//参数pageSize: 类型long, 参数可以为空
//  描述:每页显示数
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，可在产品概况中查看
func QuerySoftwareUpgradeSubtasks(appKey string, appSecret string, id string, productId string, MasterKey string, taskStatus string, searchValue string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/details"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId
	param["taskStatus"] = taskStatus
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200529233212"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey,产品概况中查看
func QuerySoftwareUpgradeTask(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20200529233136"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，产品概况可以查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateSoftwareUpgradeTask(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20200529233123"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，在产品概况中查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func ModifySoftwareUpgradeTask(appKey string, appSecret string, id string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20200529233103"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，在产品概况中查看
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func ControlSoftwareUpgradeTask(appKey string, appSecret string, id string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/control"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20200529233046"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数id: 类型long, 参数不可以为空
//  描述:主任务id
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数updateBy: 类型String, 参数可以为空
//  描述:修改人
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，在产品概况中查看
func DeleteSoftwareUpgradeTask(appKey string, appSecret string, id string, productId string, MasterKey string, updateBy string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/task"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId
	param["updateBy"] = updateBy

	version := "20200529233037"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数id: 类型long, 参数可以为空
//  描述:主任务id,isSelectDevice为1时必填，为2不必填
//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数isSelectDevice: 类型String, 参数不可以为空
//  描述:查询类型（1.查询加入升级设备，2.查询可加入升级设备）
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页，默认1
//参数pageSize: 类型long, 参数可以为空
//  描述:每页显示数，默认20
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，产品概况中查看
//参数deviceIdSearch: 类型String, 参数可以为空
//  描述:根据设备id精确查询
//参数deviceNameSearch: 类型String, 参数可以为空
//  描述:根据设备名称精确查询
//参数imeiSearch: 类型String, 参数可以为空
//  描述:根据imei号精确查询，仅支持LWM2M协议
//参数deviceGroupIdSearch: 类型String, 参数可以为空
//  描述:根据群组id精确查询
func QuerySoftwareUpradeDeviceList(appKey string, appSecret string, productId string, isSelectDevice string, MasterKey string, id string, pageNow string, pageSize string, deviceIdSearch string, deviceNameSearch string, imeiSearch string, deviceGroupIdSearch string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/devices"

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
	param["deviceGroupIdSearch"] = deviceGroupIdSearch

	version := "20200529233027"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数id: 类型long, 参数不可以为空
//  描述:
//参数productId: 类型long, 参数不可以为空
//  描述:
//参数MasterKey: 类型String, 参数不可以为空
//  描述:
func QuerySoftwareUpgradeDetail(appKey string, appSecret string, id string, productId string, MasterKey string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/detail"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["id"] = id
	param["productId"] = productId

	version := "20200529233010"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数productId: 类型long, 参数不可以为空
//  描述:产品id
//参数pageNow: 类型long, 参数可以为空
//  描述:当前页数，默认1
//参数pageSize: 类型long, 参数可以为空
//  描述:每页显示数，默认20
//参数MasterKey: 类型String, 参数不可以为空
//  描述:MasterKey，产品概况中查看
//参数searchValue: 类型String, 参数可以为空
//  描述:查询条件，支持主任务名称模糊查询
func QuerySoftwareUpgradeTaskList(appKey string, appSecret string, productId string, MasterKey string, pageNow string, pageSize string, searchValue string) (*http.Response, error) {
	path := "/aep_software_upgrade_management/tasks"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize
	param["searchValue"] = searchValue

	version := "20200529232940"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
