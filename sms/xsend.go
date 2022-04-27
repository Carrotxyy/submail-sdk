/**
 * @Author: Jacky
 * @Description:
 * @File: xsend
 * @Version: 1.0.0
 * @Date: 2022/4/27 17:04
 */
package sms

import "github.com/Carrotxyy/submail-sdk"

const xsendURL = "https://api-v4.mysubmail.com/sms/xsend"

type XSendRequest struct {
	AppID       string `json:"appid"`
	To          string `json:"to"`
	Project     string `json:"project"` // 模板 ID
	Vars        string `json:"vars"`    // 变量 json 字符串
	Tag         string `json:"tag,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	SignType    string `json:"sign_type,omitempty"`
	SignVersion string `json:"sign_version,omitempty"`
	Signature   string `json:"signature"`
}

type XSendResponse struct {
	Status string `json:"status"`
	SendID string `json:"send_id"`
	Fee    int    `json:"fee"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func NewXSendRequest(to string, project string, vars string) *XSendRequest {
	return &XSendRequest{
		To:      to,
		Project: project,
		Vars:    vars,
	}
}

func (r *XSendRequest) SignMeta() string {
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
func (r *XSendRequest) SetAppID(appID string) {
	r.AppID = appID
}
func (r *XSendRequest) SetTag(tag string) {
	r.Tag = tag
}
func (r *XSendRequest) SetTimestamp(timestamp string) {
	r.Timestamp = timestamp
}
func (r *XSendRequest) SetSignType(signType string) {
	r.SignType = signType
}
func (r *XSendRequest) SetSignVersion(signVersion string) {
	r.SignVersion = signVersion
}
func (r *XSendRequest) SetSignature(signature string) {
	r.Signature = signature
}
