package whitelist_proxy

import (
	"io"
	"net/http"
	"testing"
)

func TestProxyRequest(t *testing.T) {
	// 1. 创建代理请求
	req, err := NewRequestWithProxy(
		http.MethodGet,
		"https://httpbin.org/get?name=test",
		nil,
		true,
	)
	if err != nil {
		t.Fatalf("创建请求失败: %v", err)
	}

	// 2. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 3. 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("读取响应失败: %v", err)
	}

	// 4. 验证结果
	t.Logf("状态码: %d", resp.StatusCode)
	t.Logf("响应体:\n%s", string(body))

	if resp.StatusCode != http.StatusOK {
		t.Errorf("期望状态码 200，实际 %d", resp.StatusCode)
	}
}
