package whitelist_proxy

import (
	"fmt"
	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/services/transferbill"
	"github.com/line-lee/whitelist-proxy/confidential"
	"testing"
)

func TestTransferBill(t *testing.T) {
	request := TransferBillRequest{
		PrivateKeyStr:       confidential.PrivateKeyStr,
		MchID:               confidential.MchID,
		CertificateSerialNo: confidential.CertificateSerialNo,
		MchAPIv3Key:         confidential.MchAPIv3Key,
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
