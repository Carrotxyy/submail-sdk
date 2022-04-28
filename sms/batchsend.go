/**
 * @Author: Jacky
 * @Description:
 * @File: batchsend
 * @Version: 1.0.0
 * @Date: 2022/4/28 15:17
 */
package sms

import "github.com/Carrotxyy/submail-sdk"

const batchSendURL = "https://api-v4.mysubmail.com/sms/batchsend"

type BatchSendRequest struct {
	AppID       string `json:"appid"`
	To          string `json:"to"`
	Content     string `json:"content"`
	Tag         string `json:"tag,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	SignType    string `json:"sign_type,omitempty"`
	SignVersion string `json:"sign_version,omitempty"`
	Signature   string `json:"signature"`
}

type BatchSendResponse struct {
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

func NewBatchSendRequest(to string, content string) *BatchSendRequest {
	return &BatchSendRequest{
		To:      to,
		Content: content,
	}
}

func (r *BatchSendRequest) SignMeta() string {
	data := make(map[string]string)
	data["appid"] = r.AppID
	data["to"] = r.To
	data["timestamp"] = r.Timestamp
	data["sign_type"] = r.SignType
	data["sign_version"] = r.SignVersion
	if submail.Version1.Is(r.SignVersion) {
		data["content"] = r.Content
	}

	return submail.SortAndJoin(data, "&")
}
func (r *BatchSendRequest) SetAppID(appID string) {
	r.AppID = appID
}
func (r *BatchSendRequest) SetTag(tag string) {
	r.Tag = tag
}
func (r *BatchSendRequest) SetTimestamp(timestamp string) {
	r.Timestamp = timestamp
}
func (r *BatchSendRequest) SetSignType(signType string) {
	r.SignType = signType
}
func (r *BatchSendRequest) SetSignVersion(signVersion string) {
	r.SignVersion = signVersion
}
func (r *BatchSendRequest) SetSignature(signature string) {
	r.Signature = signature
}
