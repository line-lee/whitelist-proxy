package whitelist_proxy

import (
	"encoding/json"
	"fmt"
	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/services/transferbill"
	"log"
	"os"
	"testing"
)

type WechatPayConfig struct {
	PrivateKeyStr       string
	MchID               string
	CertificateSerialNo string
	MchAPIv3Key         string
}

var wpc *WechatPayConfig

func GetWechatPayConfig() *WechatPayConfig {
	if wpc != nil {
		return wpc
	}
	data, err := os.ReadFile("E:\\github\\whitelist-proxy-wechatpay-config.json")
	if err != nil {
		log.Fatalln("读取微信支付配置文件失败:", err)
	}
	// 反序列化到结构体对象
	err = json.Unmarshal(data, &wpc)
	if err != nil {
		log.Fatalln("微信支付配置反序列化失败:", err)
	}
	return wpc
}

func TestTransferBill(t *testing.T) {
	wechatPayConfig := GetWechatPayConfig()
	request := TransferBillRequest{
		PrivateKeyStr:       wechatPayConfig.PrivateKeyStr,
		MchID:               wechatPayConfig.MchID,
		CertificateSerialNo: wechatPayConfig.CertificateSerialNo,
		MchAPIv3Key:         wechatPayConfig.MchAPIv3Key,
		Request: transferbill.CreateTransferBillRequest{
			Appid:           core.String("wx2db503e0c502dd4b"),
			OutBillNo:       core.String("xk202606031620no23000001"),
			TransferSceneId: core.String("1000"),
			Openid:          core.String("ooTcG5SvKNXydNizzjdL7dFAo13Y"),
			TransferAmount:  core.Int64(13),
			TransferRemark:  core.String("测试转账功能"),
			TransferSceneReportInfos: []transferbill.TransferSceneReportInfo{
				{
					InfoType:    core.String("活动名称"),
					InfoContent: core.String("提现"),
				},
				{
					InfoType:    core.String("奖励说明"),
					InfoContent: core.String(fmt.Sprintf("共计%.2f元", float64(13)/float64(100))),
				},
			},
		},
	}
	// 3. 调用方法
	resp := TransferBill(request)
	// 4. 验证结果
	if resp.Code != Success {
		t.Errorf("TransferBill failed, code: %d, msg: %s", resp.Code, resp.Msg)
	}
	if resp.Response != nil {
		t.Logf("TransferBill success: %+v", resp.Response)
	}
}
