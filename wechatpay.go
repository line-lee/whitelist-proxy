package whitelist_proxy

import (
	"encoding/json"
	"github.com/line-lee/wechat-pay/services/transferbatch"
	"github.com/line-lee/wechat-pay/services/transferbill"
)

type TransferBatchesRequest struct {
	PrivateKeyStr       string                                     `json:"private_key_str" form:"private_key_str"`
	MchID               string                                     `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string                                     `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string                                     `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	Request             transferbatch.InitiateBatchTransferRequest `json:"request" form:"request"`
}

type TransferBatchesResponse struct {
	Code     int                                          `json:"code"`
	Msg      string                                       `json:"msg"`
	Response *transferbatch.InitiateBatchTransferResponse `json:"response"`
}

func TransferBatches(request TransferBatchesRequest) *TransferBatchesResponse {
	body := sdkProxy("/wechatpay/transfer/batches", request)
	resp := new(TransferBatchesResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}

type GetTransferBatchByNoRequest struct {
	PrivateKeyStr       string                                    `json:"private_key_str" form:"private_key_str"`
	MchID               string                                    `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string                                    `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string                                    `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	Request             transferbatch.GetTransferBatchByNoRequest `json:"request" form:"request"`
}

type GetTransferBatchByNoResponse struct {
	Code     int                                `json:"code"`
	Msg      string                             `json:"msg"`
	Response *transferbatch.TransferBatchEntity `json:"response"`
}

func GetTransferBatchByNo(request GetTransferBatchByNoRequest) *GetTransferBatchByNoResponse {
	body := sdkProxy("/wechatpay/transfer/batches/no", request)
	resp := new(GetTransferBatchByNoResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}

type GetTransferDetailByNoRequest struct {
	PrivateKeyStr       string                                     `json:"private_key_str" form:"private_key_str"`
	MchID               string                                     `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string                                     `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string                                     `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	Request             transferbatch.GetTransferDetailByNoRequest `json:"request" form:"request"`
}

type GetTransferDetailByNoResponse struct {
	Code     int                                 `json:"code"`
	Msg      string                              `json:"msg"`
	Response *transferbatch.TransferDetailEntity `json:"response"`
}

func GetTransferDetailByNo(request GetTransferDetailByNoRequest) *GetTransferDetailByNoResponse {
	body := sdkProxy("/wechatpay/transfer/batches/detail/no", request)
	resp := new(GetTransferDetailByNoResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}

type GetTransferDetailByOutNoRequest struct {
	PrivateKeyStr       string                                        `json:"private_key_str" form:"private_key_str"`
	MchID               string                                        `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string                                        `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string                                        `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	Request             transferbatch.GetTransferDetailByOutNoRequest `json:"request" form:"request"`
}

type GetTransferDetailByOutNoResponse struct {
	Code     int                                 `json:"code"`
	Msg      string                              `json:"msg"`
	Response *transferbatch.TransferDetailEntity `json:"response"`
}

func GetTransferDetailByOutNo(request GetTransferDetailByOutNoRequest) *GetTransferDetailByOutNoResponse {
	body := sdkProxy("/wechatpay/transfer/batches/detail/out_no", request)
	resp := new(GetTransferDetailByOutNoResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}

type TransferBillRequest struct {
	PrivateKeyStr       string                                 `json:"private_key_str" form:"private_key_str"`
	MchID               string                                 `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string                                 `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string                                 `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	Request             transferbill.CreateTransferBillRequest `json:"request" form:"request"`
}

type TransferBillResponse struct {
	Code     int                                      `json:"code"`
	Msg      string                                   `json:"msg"`
	Response *transferbill.CreateTransferBillResponse `json:"response"`
}

func TransferBill(request TransferBillRequest) *TransferBillResponse {
	body := sdkProxy("/wechatpay/transfer/bill", request)
	resp := new(TransferBillResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}

type GetTransferBillByNoRequest struct {
	PrivateKeyStr       string                                   `json:"private_key_str" form:"private_key_str"`
	MchID               string                                   `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string                                   `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string                                   `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	Request             transferbill.GetlTransferBillByNoRequest `json:"request" form:"request"`
}

type GetTransferBillByNoResponse struct {
	Code     int                                        `json:"code"`
	Msg      string                                     `json:"msg"`
	Response *transferbill.GetlTransferBillByNoResponse `json:"response"`
}

func GetTransferBillByNo(request GetTransferBillByNoRequest) *GetTransferBillByNoResponse {
	body := sdkProxy("/wechatpay/transfer/bill/no", request)
	resp := new(GetTransferBillByNoResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}

type TransferBillNotifyRequest struct {
	PrivateKeyStr       string            `json:"private_key_str" form:"private_key_str"`
	MchID               string            `json:"mch_id" form:"mch_id"`
	CertificateSerialNo string            `json:"certificate_serial_no" form:"certificate_serial_no"`
	MchAPIv3Key         string            `json:"mch_api_v3_key" form:"mch_api_v3_key"`
	NotifyHeader        map[string]string `json:"notify_header" form:"notify_header"`
	NotifyBody          string            `json:"notify_body" form:"notify_body"`
}

type TransferBillNotifyResponse struct {
	Code     int                         `json:"code"`
	Msg      string                      `json:"msg"`
	Response *TransferBillNotifyResource `json:"response"`
}

type TransferBillNotifyResource struct {
	//【商户单号】商户系统内部的商家单号，在商户系统内部唯一
	OutBillNo string `json:"out_bill_no"`
	//【商家转账订单号】微信单号，微信商家转账系统返回的唯一标识
	TransferBillNo string `json:"transfer_bill_no"`
	//【单据状态】商家转账订单状态
	//ACCEPTED：单据已受理
	//PROCESSING：单据处理中，转账结果尚未明确，如一直处于此状态，建议检查账户余额是否足够
	//WAIT_USER_CONFIRM：待收款用户确认，可拉起微信收款确认页面进行收款确认
	//TRANSFERING：转账中，转账结果尚未明确，可拉起微信收款确认页面再次重试确认收款
	//SUCCESS： 转账成功
	//FAIL： 转账失败
	//CANCELING： 撤销中
	//CANCELLED： 已撤销
	State string `json:"state"`
	//【商户号】微信支付分配的商户号
	MchId string `json:"mch_id"`
	//【转账金额】转账总金额，单位为“分”
	TransferAmount int64 `json:"transfer_amount"`
	//【收款用户OpenID】用户在商户appid下的唯一标识
	Openid string `json:"openid"`
	//【失败原因】单已失败或者已退资金时，会返回订单失败原因
	FailReason string `json:"fail_reason"`
	//【单据创建时间】遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示北京时间2015年05月20日13点29分35秒。
	CreateTime string `json:"create_time"`
	//【最后一次状态变更时间】遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示北京时间2015年05月20日13点29分35秒。
	UpdateTime string `json:"update_time"`
}

func TransferBillNotify(request TransferBillNotifyRequest) *TransferBillNotifyResponse {
	body := sdkProxy("/wechatpay/transfer/bill/notify", request)
	resp := new(TransferBillNotifyResponse)
	_ = json.Unmarshal(body, resp)
	return resp
}
