/**
 * @Author: Jacky
 * @Description:
 * @File: auth
 * @Version: 1.0.0
 * @Date: 2022/4/27 14:40
 */
package submail

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

type Message interface {
	SignMeta() string                  // 返回签名元数据
	SetAppID(appID string)             // 设置 appid
	SetTag(tag string)                 // 设置 tag
	SetTimestamp(timestamp string)     // 设置时间戳
	SetSignType(signType string)       // 设置 api 授权模式
	SetSignVersion(signVersion string) // 设置加密计算方式
	SetSignature(signature string)     // 设置签名
}

type (
	SignType    string
	SignVersion string
)

const (
	MD5    SignType = "md5"
	SHA1   SignType = "sha1"
	Normal SignType = "normal"

	Version1 SignVersion = "1"
	Version2 SignVersion = "2"
)

func (v SignVersion) Is(version string) bool {
	return string(v) == version
}

type auth struct {
	appid       string
	appkey      string
	tag         string
	signType    SignType
	signVersion SignVersion
}

func (a *auth) sign(message Message) {
	message.SetAppID(a.appid)
	if a.signType == Normal {
		message.SetSignature(a.appkey)
		return
	}

	// 设置 message 字段值
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	message.SetTag(a.tag)
	message.SetTimestamp(timestamp)
	message.SetSignType(string(a.signType))
	message.SetSignVersion(string(a.signVersion))

	signMeta := message.SignMeta()
	signStr := a.appid + a.appkey + signMeta + a.appid + a.appkey

	var sign string
	switch a.signType {
	case MD5:
		sign = a.md5(signStr)
	case SHA1:
		sign = a.sha1(signStr)
	}

	message.SetSignature(sign)

	return
}

func (a *auth) md5(data string) string {
	_md5 := md5.New()
	io.WriteString(_md5, data)

	return hex.EncodeToString(_md5.Sum(nil))
}

func (a *auth) sha1(data string) string {
	_sha1 := sha1.New()
	io.WriteString(_sha1, data)

	return hex.EncodeToString(_sha1.Sum(nil))
}
