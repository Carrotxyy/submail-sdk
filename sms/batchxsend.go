/**
 * @Author: Jacky
 * @Description:
 * @File: batchxsend
 * @Version: 1.0.0
 * @Date: 2022/4/28 15:27
 */
package sms

import "github.com/Carrotxyy/submail-sdk"

const batchXSendURL = "https://api-v4.mysubmail.com/sms/batchxsend"

type BatchXSendURLRequest struct {
	AppID       string `json:"appid"`
	To          string `json:"to"`
	Project     string `json:"project"` // 模板 ID
	Vars        string `json:"vars"`
	Tag         string `json:"tag,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	SignType    string `json:"sign_type,omitempty"`
	SignVersion string `json:"sign_version,omitempty"`
	Signature   string `json:"signature"`
}

type BatchXSendURLResponse struct {
	Status    string `json:"status"`
	Batchlist string `json:"batchlist"` // 任务 ID
	TotalFee  int    `json:"total_fee"` // 总计费条数
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Response  []struct {
		To     string `json:"to"`
		Status string `json:"status"`
		SendID string `json:"send_id"`
		Fee    int    `json:"fee"`
	} `json:"response"`
}

func NewBatchXSendURLRequest(to string, project string, vars map[string]string) *BatchXSendURLRequest {
	return &BatchXSendURLRequest{
		To:      to,
		Project: project,
		Vars:    submail.JsonMarshal(vars),
	}
}

func (r *BatchXSendURLRequest) SignMeta() string {
	data := make(map[string]string)
	data["appid"] = r.AppID
	data["to"] = r.To
	data["project"] = r.Project
	data["timestamp"] = r.Timestamp
	data["sign_type"] = r.SignType
	data["sign_version"] = r.SignVersion
	if submail.Version1.Is(r.SignVersion) {
		data["vars"] = r.Vars
	}

	return submail.SortAndJoin(data, "&")
}
func (r *BatchXSendURLRequest) SetAppID(appID string) {
	r.AppID = appID
}
func (r *BatchXSendURLRequest) SetTag(tag string) {
	r.Tag = tag
}
func (r *BatchXSendURLRequest) SetTimestamp(timestamp string) {
	r.Timestamp = timestamp
}
func (r *BatchXSendURLRequest) SetSignType(signType string) {
	r.SignType = signType
}
func (r *BatchXSendURLRequest) SetSignVersion(signVersion string) {
	r.SignVersion = signVersion
}
func (r *BatchXSendURLRequest) SetSignature(signature string) {
	r.Signature = signature
}
