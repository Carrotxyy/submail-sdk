/**
 * @Author: Jacky
 * @Description:
 * @File: multixsend
 * @Version: 1.0.0
 * @Date: 2022/4/28 14:51
 */
package sms

import "github.com/Carrotxyy/submail-sdk"

const multiXSendURL = "https://api-v4.mysubmail.com/sms/multixsend"

type MultiXSendFieldMulti struct {
	To   string            `json:"to"`
	Vars map[string]string `json:"vars"`
}

type MultiXSendRequest struct {
	AppID       string `json:"appid"`
	Project     string `json:"project"` // 模板 ID
	Multi       string `json:"multi"`
	Tag         string `json:"tag,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	SignType    string `json:"sign_type,omitempty"`
	SignVersion string `json:"sign_version,omitempty"`
	Signature   string `json:"signature"`
}

type MultiXSendResponse struct {
	To     string `json:"to"`
	Status string `json:"status"`
	SendID string `json:"send_id"`
	Fee    int    `json:"fee"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func NewMultiXSendRequest(project string, multi []MultiXSendFieldMulti) *MultiXSendRequest {
	return &MultiXSendRequest{
		Project: project,
		Multi:   submail.JsonMarshal(multi),
	}
}

func (r *MultiXSendRequest) SignMeta() string {
	data := make(map[string]string)
	data["appid"] = r.AppID
	data["project"] = r.Project
	data["timestamp"] = r.Timestamp
	data["sign_type"] = r.SignType
	data["sign_version"] = r.SignVersion
	if submail.Version1.Is(r.SignVersion) {
		data["multi"] = r.Multi
	}

	return submail.SortAndJoin(data, "&")
}
func (r *MultiXSendRequest) SetAppID(appID string) {
	r.AppID = appID
}
func (r *MultiXSendRequest) SetTag(tag string) {
	r.Tag = tag
}
func (r *MultiXSendRequest) SetTimestamp(timestamp string) {
	r.Timestamp = timestamp
}
func (r *MultiXSendRequest) SetSignType(signType string) {
	r.SignType = signType
}
func (r *MultiXSendRequest) SetSignVersion(signVersion string) {
	r.SignVersion = signVersion
}
func (r *MultiXSendRequest) SetSignature(signature string) {
	r.Signature = signature
}
