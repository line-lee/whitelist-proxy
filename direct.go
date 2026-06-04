package whitelist_proxy

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func NewRequestWithProxy(method, address string, body io.Reader, isProxy bool) (*http.Request, error) {
	if !isProxy {
		log.Printf("不经过代理服务器，本地访问：%s | %s\n", method, address)
		return http.NewRequest(method, address, body)
	}
	// 解析原始URL提取域名:例如：https://rvakva.xiaokayun.cn/v1/openCompany
	u, err := url.Parse(address)
	if err != nil {
		log.Printf("经过代理服务器，访问路径解析失败：%s | %s\n", method, address)
		return nil, err
	}
	// 提取 scheme + host 作为代理目标域名
	domain := u.Scheme + "://" + u.Host
	// 构建代理地址：http://47.93.183.173:30771 + 原始路径
	proxyURL := "http://47.93.183.173:30771/direct" + u.Path
	if u.RawQuery != "" {
		proxyURL += "?" + u.RawQuery
	}
	request, err := http.NewRequest(method, proxyURL, body)
	if err != nil {
		log.Printf("经过代理服务器，构建request失败：%s | %s\n", method, address)
		return nil, err
	}
	// 设置 DestinationDomain 头，Nginx 会根据它代理到目标
	request.Header.Set("DestinationDomain", domain)
	return request, err
}
