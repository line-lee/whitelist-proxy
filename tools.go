package whitelist_proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	Success = 1
	Error   = 100
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 特殊代理，如微信支付，hbase统一调用方法
func sdkProxy(path string, param any) []byte {
	// 1. 参数转 JSON
	jsonData, err := json.Marshal(param)
	if err != nil {
		respBytes, _ := json.Marshal(response{Code: Error, Msg: fmt.Sprintf("SDK代理，序列化参数:%v，抛出错误：%v", param, err)})
		return respBytes
	}
	// 2. 创建请求，body 传入 JSON
	req, err := http.NewRequest("POST", "http://47.93.183.173:30771"+path, bytes.NewReader(jsonData))
	if err != nil {
		respBytes, _ := json.Marshal(response{Code: Error, Msg: fmt.Sprintf("SDK代理，NewRequest，抛出错误：%v", err)})
		return respBytes
	}
	// 3. 设置 Content-Type
	req.Header.Set("Content-Type", "application/json")
	// 4. 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		respBytes, _ := json.Marshal(response{Code: Error, Msg: fmt.Sprintf("SDK代理，DefaultClient.Do，抛出错误：%v", err)})
		return respBytes
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	// 5. 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		respBytes, _ := json.Marshal(response{Code: Error, Msg: fmt.Sprintf("SDK代理，response Body io.ReadAll，抛出错误：%v", err)})
		return respBytes
	}
	log.Printf("SDK代理结果: %s\n", string(body))
	return body
}
