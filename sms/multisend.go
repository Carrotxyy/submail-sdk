/**
 * @Author: Jacky
 * @Description:
 * @File: multisend
 * @Version: 1.0.0
 * @Date: 2022/4/28 11:43
 */
package sms

import "github.com/Carrotxyy/submail-sdk"

const multiSendURL = "https://api-v4.mysubmail.com/sms/multisend"

type MultiSendFieldMulti struct {
	To   string            `json:"to"`
	Vars map[string]string `json:"vars"`
}

type MultiSendRequest struct {
	AppID       string `json:"appid"`
	Content     string `json:"content"`
	Multi       string `json:"multi"`
	Tag         string `json:"tag,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	SignType    string `json:"sign_type,omitempty"`
	SignVersion string `json:"sign_version,omitempty"`
	Signature   string `json:"signature"`
}

type MultiSendResponse struct {
	To     string `json:"to"`
	Status string `json:"status"`
	SendID string `json:"send_id"`
	Fee    int    `json:"fee"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func NewMultiSendRequest(content string, multi []MultiSendFieldMulti) *MultiSendRequest {
	return &MultiSendRequest{
		Content: content,
		Multi:   submail.JsonMarshal(multi),
	}
}

func (r *MultiSendRequest) SignMeta() string {
	data := make(map[string]string)
	data["appid"] = r.AppID
	data["timestamp"] = r.Timestamp
	data["sign_type"] = r.SignType
	data["sign_version"] = r.SignVersion
	if submail.Version1.Is(r.SignVersion) {
		data["content"] = r.Content
		data["multi"] = r.Multi
	}

	return submail.SortAndJoin(data, "&")
}
func (r *MultiSendRequest) SetAppID(appID string) {
	r.AppID = appID
}
func (r *MultiSendRequest) SetTag(tag string) {
	r.Tag = tag
}
func (r *MultiSendRequest) SetTimestamp(timestamp string) {
	r.Timestamp = timestamp
}
func (r *MultiSendRequest) SetSignType(signType string) {
	r.SignType = signType
}
func (r *MultiSendRequest) SetSignVersion(signVersion string) {
	r.SignVersion = signVersion
}
func (r *MultiSendRequest) SetSignature(signature string) {
	r.Signature = signature
}
