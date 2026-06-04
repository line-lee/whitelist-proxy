package whitelist_proxy

import (
	"encoding/json"
	"testing"
)

func TestHbaseGet(t *testing.T) {
	param := HbaseGetParam{
		Table: "gpsrange",
		Key:   "103564437",
	}
	result := HbaseGet(param)
	if result == nil {
		t.Error("HbaseGet returned nil")
		return
	}
	// 简单的输出验证，实际应根据业务判断 Code
	t.Logf("HbaseGet result - Code: %d, Msg: %s", result.Code, result.Msg)
	if result.Result != nil {
		t.Logf("Result cells count: %d", len(result.Result.Cells))
	}
}

type GpsPoint struct {
	Lat          float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng          float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	LocationType int32   `protobuf:"varint,3,opt,name=location_type,json=locationType,proto3" json:"location_type,omitempty"`
	Accuracy     float64 `protobuf:"fixed64,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	PositionTime int64   `protobuf:"varint,5,opt,name=position_time,json=positionTime,proto3" json:"position_time,omitempty"`
	OrderId      int64   `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status       int64   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	FenceIdx     int64   `protobuf:"varint,8,opt,name=fence_idx,json=fenceIdx,proto3" json:"fence_idx,omitempty"`
	IsFilter     bool    `protobuf:"varint,9,opt,name=is_filter,json=isFilter,proto3" json:"is_filter,omitempty"`
	FilterMemo   string  `protobuf:"bytes,10,opt,name=filter_memo,json=filterMemo,proto3" json:"filter_memo,omitempty"`
}

func TestHbasePut(t *testing.T) {
	// 测试参数
	//2026/06/04 11:19:54 SDK代理路径: /hbase/get
	//2026/06/04 11:19:54 SDK代理参数: {"table":"gpsrange","key":"103564437"}
	//2026/06/04 11:19:54 SDK代理结果: {"code":1,"msg":"","result":{"Cells":[{"row":"MTAzNTY0NDM3","family":"Z3Bz","qualifier":"ZGF0YQ==","timestamp":1778554792214,"cell_type":4,"value":"eyJsYXQiOjMwLjk2NDI3OSwibG5nIjoxMDQuMjc3OTk0LCJsb2NhdGlvbl90eXBlIjo1LCJhY2N1cmFjeSI6NDMsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODMsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MSwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQ0NjIsImxuZyI6MTA0LjI3ODI3NSwibG9jYXRpb25fdHlwZSI6NSwiYWNjdXJhY3kiOjkwLCJwb3NpdGlvbl90aW1lIjoxNzc4NTU0Nzg0LCJvcmRlcl9pZCI6MTAzNTY0NDM3LCJzdGF0dXMiOjEsImZlbmNlX2lkeCI6MCwiaXNfZmlsdGVyIjpmYWxzZSwiZmlsdGVyX21lbW8iOiIifSx7ImxhdCI6MzAuOTY0MjI0NzE2MTQ1MDY1LCJsbmciOjEwNC4yNzgwMzA2NjUwNzg4LCJsb2NhdGlvbl90eXBlIjoxLCJhY2N1cmFjeSI6MjAsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODYsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MiwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQyNTMzNjQ2MzY3OCwibG5nIjoxMDQuMjc4MDUzNTg5MDkwMzksImxvY2F0aW9uX3R5cGUiOjEsImFjY3VyYWN5Ijo0LjkwMDAwMDA5NTM2NzQzMiwicG9zaXRpb25fdGltZSI6MTc3ODU1NDc5MSwib3JkZXJfaWQiOjEwMzU2NDQzNywic3RhdHVzIjoyLCJmZW5jZV9pZHgiOjAsImlzX2ZpbHRlciI6ZmFsc2UsImZpbHRlcl9tZW1vIjoiIn0s"}],"Stale":false,"Partial":false,"Exists":null}}

	str := "{\"code\":1,\"msg\":\"\",\"result\":{\"Cells\":[{\"row\":\"MTAzNTY0NDM3\",\"family\":\"Z3Bz\",\"qualifier\":\"ZGF0YQ==\",\"timestamp\":1778554792214,\"cell_type\":4,\"value\":\"eyJsYXQiOjMwLjk2NDI3OSwibG5nIjoxMDQuMjc3OTk0LCJsb2NhdGlvbl90eXBlIjo1LCJhY2N1cmFjeSI6NDMsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODMsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MSwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQ0NjIsImxuZyI6MTA0LjI3ODI3NSwibG9jYXRpb25fdHlwZSI6NSwiYWNjdXJhY3kiOjkwLCJwb3NpdGlvbl90aW1lIjoxNzc4NTU0Nzg0LCJvcmRlcl9pZCI6MTAzNTY0NDM3LCJzdGF0dXMiOjEsImZlbmNlX2lkeCI6MCwiaXNfZmlsdGVyIjpmYWxzZSwiZmlsdGVyX21lbW8iOiIifSx7ImxhdCI6MzAuOTY0MjI0NzE2MTQ1MDY1LCJsbmciOjEwNC4yNzgwMzA2NjUwNzg4LCJsb2NhdGlvbl90eXBlIjoxLCJhY2N1cmFjeSI6MjAsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODYsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MiwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQyNTMzNjQ2MzY3OCwibG5nIjoxMDQuMjc4MDUzNTg5MDkwMzksImxvY2F0aW9uX3R5cGUiOjEsImFjY3VyYWN5Ijo0LjkwMDAwMDA5NTM2NzQzMiwicG9zaXRpb25fdGltZSI6MTc3ODU1NDc5MSwib3JkZXJfaWQiOjEwMzU2NDQzNywic3RhdHVzIjoyLCJmZW5jZV9pZHgiOjAsImlzX2ZpbHRlciI6ZmFsc2UsImZpbHRlcl9tZW1vIjoiIn0s\"}],\"Stale\":false,\"Partial\":false,\"Exists\":null}}"
	var hbaseGetResult HbaseGetResult
	_ = json.Unmarshal([]byte(str), &hbaseGetResult)
	val := "[" + string(hbaseGetResult.Result.Cells[0].Value)[:len(string(hbaseGetResult.Result.Cells[0].Value))-1] + "]"
	points := new([]*GpsPoint)
	_ = json.Unmarshal([]byte(val), &points)

	var putStr string
	for _, p := range *points {
		point, err := json.Marshal(p)
		if err != nil {
			continue
		}
		putStr += string(point) + ","
	}
	param := HbasePutParam{
		Table: "gpsrange",
		Key:   "103564437",
		Values: map[string]map[string]string{
			"gps": {
				"data": putStr,
			},
		},
	}
	result := HbasePut(param)
	if result == nil {
		t.Error("HbasePut returned nil")
		return
	}
	t.Logf("HbasePut result - Code: %d, Msg: %s", result.Code, result.Msg)
}

func TestHbaseAppend(t *testing.T) {
	// 测试参数
	//2026/06/04 11:19:54 SDK代理路径: /hbase/get
	//2026/06/04 11:19:54 SDK代理参数: {"table":"gpsrange","key":"103564437"}
	//2026/06/04 11:19:54 SDK代理结果: {"code":1,"msg":"","result":{"Cells":[{"row":"MTAzNTY0NDM3","family":"Z3Bz","qualifier":"ZGF0YQ==","timestamp":1778554792214,"cell_type":4,"value":"eyJsYXQiOjMwLjk2NDI3OSwibG5nIjoxMDQuMjc3OTk0LCJsb2NhdGlvbl90eXBlIjo1LCJhY2N1cmFjeSI6NDMsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODMsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MSwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQ0NjIsImxuZyI6MTA0LjI3ODI3NSwibG9jYXRpb25fdHlwZSI6NSwiYWNjdXJhY3kiOjkwLCJwb3NpdGlvbl90aW1lIjoxNzc4NTU0Nzg0LCJvcmRlcl9pZCI6MTAzNTY0NDM3LCJzdGF0dXMiOjEsImZlbmNlX2lkeCI6MCwiaXNfZmlsdGVyIjpmYWxzZSwiZmlsdGVyX21lbW8iOiIifSx7ImxhdCI6MzAuOTY0MjI0NzE2MTQ1MDY1LCJsbmciOjEwNC4yNzgwMzA2NjUwNzg4LCJsb2NhdGlvbl90eXBlIjoxLCJhY2N1cmFjeSI6MjAsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODYsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MiwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQyNTMzNjQ2MzY3OCwibG5nIjoxMDQuMjc4MDUzNTg5MDkwMzksImxvY2F0aW9uX3R5cGUiOjEsImFjY3VyYWN5Ijo0LjkwMDAwMDA5NTM2NzQzMiwicG9zaXRpb25fdGltZSI6MTc3ODU1NDc5MSwib3JkZXJfaWQiOjEwMzU2NDQzNywic3RhdHVzIjoyLCJmZW5jZV9pZHgiOjAsImlzX2ZpbHRlciI6ZmFsc2UsImZpbHRlcl9tZW1vIjoiIn0s"}],"Stale":false,"Partial":false,"Exists":null}}

	str := "{\"code\":1,\"msg\":\"\",\"result\":{\"Cells\":[{\"row\":\"MTAzNTY0NDM3\",\"family\":\"Z3Bz\",\"qualifier\":\"ZGF0YQ==\",\"timestamp\":1778554792214,\"cell_type\":4,\"value\":\"eyJsYXQiOjMwLjk2NDI3OSwibG5nIjoxMDQuMjc3OTk0LCJsb2NhdGlvbl90eXBlIjo1LCJhY2N1cmFjeSI6NDMsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODMsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MSwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQ0NjIsImxuZyI6MTA0LjI3ODI3NSwibG9jYXRpb25fdHlwZSI6NSwiYWNjdXJhY3kiOjkwLCJwb3NpdGlvbl90aW1lIjoxNzc4NTU0Nzg0LCJvcmRlcl9pZCI6MTAzNTY0NDM3LCJzdGF0dXMiOjEsImZlbmNlX2lkeCI6MCwiaXNfZmlsdGVyIjpmYWxzZSwiZmlsdGVyX21lbW8iOiIifSx7ImxhdCI6MzAuOTY0MjI0NzE2MTQ1MDY1LCJsbmciOjEwNC4yNzgwMzA2NjUwNzg4LCJsb2NhdGlvbl90eXBlIjoxLCJhY2N1cmFjeSI6MjAsInBvc2l0aW9uX3RpbWUiOjE3Nzg1NTQ3ODYsIm9yZGVyX2lkIjoxMDM1NjQ0MzcsInN0YXR1cyI6MiwiZmVuY2VfaWR4IjowLCJpc19maWx0ZXIiOmZhbHNlLCJmaWx0ZXJfbWVtbyI6IiJ9LHsibGF0IjozMC45NjQyNTMzNjQ2MzY3OCwibG5nIjoxMDQuMjc4MDUzNTg5MDkwMzksImxvY2F0aW9uX3R5cGUiOjEsImFjY3VyYWN5Ijo0LjkwMDAwMDA5NTM2NzQzMiwicG9zaXRpb25fdGltZSI6MTc3ODU1NDc5MSwib3JkZXJfaWQiOjEwMzU2NDQzNywic3RhdHVzIjoyLCJmZW5jZV9pZHgiOjAsImlzX2ZpbHRlciI6ZmFsc2UsImZpbHRlcl9tZW1vIjoiIn0s\"}],\"Stale\":false,\"Partial\":false,\"Exists\":null}}"
	var hbaseGetResult HbaseGetResult
	_ = json.Unmarshal([]byte(str), &hbaseGetResult)
	val := "[" + string(hbaseGetResult.Result.Cells[0].Value)[:len(string(hbaseGetResult.Result.Cells[0].Value))-1] + "]"
	points := new([]*GpsPoint)
	_ = json.Unmarshal([]byte(val), &points)

	var putStr string
	for _, p := range *points {
		point, err := json.Marshal(p)
		if err != nil {
			continue
		}
		putStr += string(point) + ","
	}
	param := HbaseAppendParam{
		Table: "gpsrange",
		Key:   "103564437",
		Values: map[string]map[string]string{
			"gps": {
				"data": putStr,
			},
		},
	}
	result := HbaseAppend(param)
	if result == nil {
		t.Error("HbaseAppend returned nil")
		return
	}
	t.Logf("HbaseAppend result - Code: %d, Msg: %s", result.Code, result.Msg)
}

func TestHbaseDelete(t *testing.T) {
	param := HbaseDeleteParam{
		Table: "gpsrange",
		Key:   "103564437",
	}
	result := HbaseDelete(param)
	if result == nil {
		t.Error("HbaseDelete returned nil")
		return
	}
	t.Logf("HbaseDelete result - Code: %d, Msg: %s", result.Code, result.Msg)
}
