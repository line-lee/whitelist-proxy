package whitelist_proxy

import (
	"encoding/json"
	"github.com/tsuna/gohbase/hrpc"
)

type HbaseGetParam struct {
	Table string `json:"table" form:"table"`
	Key   string `json:"key" form:"key"`
}

type HbaseGetResult struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Result *hrpc.Result `json:"result"`
}

func HbaseGet(request HbaseGetParam) *HbaseGetResult {
	body := sdkProxy("/hbase/get", request)
	resp := new(HbaseGetResult)
	_ = json.Unmarshal(body, resp)
	return resp
}

type HbasePutParam struct {
	Table  string                       `json:"table" form:"table"`
	Key    string                       `json:"key" form:"key"`
	Values map[string]map[string]string `json:"values" form:"values"`
}

type HbasePutResult struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Result *hrpc.Result `json:"result"`
}

func HbasePut(request HbasePutParam) *HbasePutResult {
	body := sdkProxy("/hbase/put", request)
	resp := new(HbasePutResult)
	_ = json.Unmarshal(body, resp)
	return resp
}

type HbaseAppendParam struct {
	Table  string                       `json:"table" form:"table"`
	Key    string                       `json:"key" form:"key"`
	Values map[string]map[string]string `json:"values" form:"values"`
}

type HbaseAppendResult struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Result *hrpc.Result `json:"result"`
}

func HbaseAppend(request HbaseAppendParam) *HbaseAppendResult {
	body := sdkProxy("/hbase/append", request)
	resp := new(HbaseAppendResult)
	_ = json.Unmarshal(body, resp)
	return resp
}

type HbaseDeleteParam struct {
	Table string `json:"table" form:"table"`
	Key   string `json:"key" form:"key"`
}

type HbaseDeleteResult struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Result *hrpc.Result `json:"result"`
}

func HbaseDelete(request HbaseDeleteParam) *HbaseDeleteResult {
	body := sdkProxy("/hbase/delete", request)
	resp := new(HbaseDeleteResult)
	_ = json.Unmarshal(body, resp)
	return resp
}
