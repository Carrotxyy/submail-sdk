/**
 * @Author: Jacky
 * @Description:
 * @File: send
 * @Version: 1.0.0
 * @Date: 2022/4/27 14:33
 */
package sms

import (
	"github.com/Carrotxyy/submail-sdk"
)

const sendURL = "https://api-v4.mysubmail.com/sms/send"

type SendRequest struct {
	AppID       string `json:"appid"`
	To          string `json:"to"`
	Content     string `json:"content"`
	Tag         string `json:"tag,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	SignType    string `json:"sign_type,omitempty"`
	SignVersion string `json:"sign_version,omitempty"`
	Signature   string `json:"signature"`
}

type SendResponse struct {
	Status string `json:"status"`
	SendID string `json:"send_id"`
	Fee    int    `json:"fee"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}

func NewSendRequest(to string, content string) *SendRequest {
	return &SendRequest{
		To:      to,
		Content: content,
	}
}

func (r *SendRequest) SignMeta() string {
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
func (r *SendRequest) SetAppID(appID string) {
	r.AppID = appID
}
func (r *SendRequest) SetTag(tag string) {
	r.Tag = tag
}
func (r *SendRequest) SetTimestamp(timestamp string) {
	r.Timestamp = timestamp
}
func (r *SendRequest) SetSignType(signType string) {
	r.SignType = signType
}
func (r *SendRequest) SetSignVersion(signVersion string) {
	r.SignVersion = signVersion
}
func (r *SendRequest) SetSignature(signature string) {
	r.Signature = signature
}
