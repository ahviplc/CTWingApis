package main

import (
	aepapi "CTWingApis/Aep_device_command"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Golang操作结构体、Map转化为JSON - 温柔的风 - 博客园
// https://www.cnblogs.com/wt645631686/p/9649379.html

func f1() {
	// CreateCommand(appKey string, appSecret string, MasterKey string , body string) (*http.Response, error)
	appKey := "HhjQH0WKysa"
	appSecret := "###"
	MasterKey := "098fd4875ec54d3e957999b22b27ecfb"

	bodyTmp := make(map[string]interface{})
	bodyTmp["productId"] = "10090896"
	bodyTmp["deviceId"] = "c30e6732eb0f4004b2e9a32db2554fcd"
	bodyTmp["operator"] = "1"
	bodyTmp["ttl"] = 0

	contentTmp := make(map[string]interface{})
	contentTmp["payload"] = "6816"
	contentTmp["dataType"] = 1

	bodyTmp["content"] = contentTmp

	buf, err := json.MarshalIndent(bodyTmp, "", "    ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf body json = ", string(buf))

	response, err := aepapi.CreateCommand(appKey, appSecret, MasterKey, string(buf))
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	fmt.Println()
	bodyStr, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		fmt.Println("ioutil.ReadAll err:", err1)
		return
	}
	fmt.Println(string(bodyStr))

	defer response.Body.Close()
}

func f2() {

	bodyTmp := make(map[string]interface{})
	bodyTmp["productId"] = "10090896"
	bodyTmp["deviceId"] = "c30e6732eb0f4004b2e9a32db2554fcd"
	bodyTmp["operator"] = "1"
	bodyTmp["ttl"] = 0

	contentTmp := make(map[string]interface{})
	contentTmp["payload"] = "68000016" // 6800233001000100036000000186000000000000000009000360000001860000243516
	contentTmp["dataType"] = 1

	bodyTmp["content"] = contentTmp

	buf, err := json.MarshalIndent(bodyTmp, "", "    ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf body json = ", string(buf))

	// CreateCommand(appKey string, appSecret string, MasterKey string , body string) (*http.Response, error)
	resp, err := aepapi.CreateCommand("HhjQH0WKysa", "###", "098fd4875ec54d3e957999b22b27ecfb", string(buf))

	if err != nil {
		fmt.Println("CreateCommand err:", err)
		return
	}
	fmt.Println("code:", resp.StatusCode, "\nerr:", err)
	bodyStr, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("ioutil.ReadAll err:", err1)
		return
	}
	fmt.Println(string(bodyStr))

	defer resp.Body.Close()
}

// 执行main方法
func main() {
	//f1()
	f2()
}
