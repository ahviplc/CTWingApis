package aepsdkcore

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	//"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const baseUrl = "https://ag-api.ctwing.cn"
const timeUrl = "https://ag-api.ctwing.cn/echo"

var offset = getTimeOffset(timeUrl)

// 获取时间偏移量(毫秒)
func getTimeOffset(url string) int64 {
	start := time.Now().UnixNano() / 1e6
	resp, _ := SendHttpRequest(url, nil, "", "GET", "")
	end := time.Now().UnixNano() / 1e6
	if resp != nil {
		time, _ := strconv.ParseInt(resp.Header.Get("x-ag-timestamp"), 10, 64)
		return time - (end+start)/2
	} else {
		return 0
	}
}

/// <summary>
/// 签名算法
/// </summary>
/// <param name="param">api接口参数</param>
/// <param name="timestamp">时间戳，毫秒级</param>
/// <param name="application">App Key</param>
/// <param name="secret">App secret</param>
/// <param name="body">body</param>
/// <returns></returns>
func signature(key string, application string, timestamp int64, param map[string]string, body string) string {
	temp := "application:" + application + "\n"
	temp += "timestamp:" + strconv.FormatInt(timestamp, 10) + "\n"

	var SortString []string
	for k := range param {
		SortString = append(SortString, k)
	}
	sort.Strings(SortString) //会根据字母的顺序进行排序。
	for _, k := range SortString {
		v, ok := param[k]
		if ok {
			temp += k + ":" + v + "\n"
		} else {
			temp += k + ":\n"
		}
	}

	// 得到需要签名的字符串
	if len([]rune(body)) > 0 {
		temp += body + "\n"
	}
	//fmt.Println("Sign string: ", temp)

	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(temp))
	result := mac.Sum(nil)

	base64Result := base64.StdEncoding.EncodeToString(result)
	//fmt.Println(base64Result)
	return base64Result
}

/// <summary>
/// 发送api请求到aep
/// </summary>
/// <param name="path">api接口路径</param>
/// <param name="headers">请求head</param>
/// <param name="param">参数</param>
/// <param name="body">body，如果为get等没有body请求，填null</param>
/// <param name="version">api接口版本，在文档中查询</param>
/// <param name="application">App Key</param>
/// <param name="key">App Secret</param>
/// <param name="method">请求的类型,GET、POST、PUT、DELETE</param>
/// <returns></returns>
func SendAepHttpRequest(path string, headers map[string]string, param map[string]string, body string, version string,
	application string, key string, method string) (*http.Response, error) {
	paramString := ""

	paramsTmp := make(map[string]string)

	if param != nil {
		for key, value := range param {
			paramString += key + "=" + value + "&"
			paramsTmp[key] = value
		}
	}

	if offset == 0 {
		offset = getTimeOffset(timeUrl)
	}

	timestamp := time.Now().UnixNano()/1e6 + offset

	for key, value := range headers {
		paramsTmp[key] = value
	}

	headersTmp := make(map[string]string)

	headersTmp["application"] = application
	headersTmp["timestamp"] = strconv.FormatInt(timestamp, 10)
	headersTmp["version"] = version
	headersTmp["signature"] = signature(key, application, timestamp, paramsTmp, body)

	for key, value := range headers {
		paramsTmp[key] = value
		headersTmp[key] = value
	}

	url := baseUrl + path
	if len([]rune(paramString)) > 0 {
		url += "?" + paramString
	}

	return SendHttpRequest(url, headersTmp, "application/json; charset=UTF-8", method, body)
}

/// <summary>
/// 处理http请求
/// </summary>
/// <param name="url">请求的url地址</param>
/// <param name="headers">协议标头</param>
/// <param name="contentType">请求的内容类型</param>
/// <param name="method">请求的类型,GET、POST、PUT、DELETE</param>
/// <param name="body">请求的数据流</param>
/// <returns></returns>
func SendHttpRequest(url string, headers map[string]string, contentType string, method string, body string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header[key] = []string{value}
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
