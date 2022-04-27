/**
 * @Author: Jacky
 * @Description:
 * @File: submail
 * @Version: 1.0.0
 * @Date: 2022/4/27 10:52
 */
package submail

import (
	"github.com/go-resty/resty/v2"
)

type Option func(c *config)
type config struct {
	tag         string
	signType    SignType
	signVersion SignVersion
}

func defaultConfig() *config {
	return &config{
		signType: Normal,
	}
}

func WithTag(tag string) Option {
	return func(c *config) {
		c.tag = tag
	}
}
func WithSignType(signType SignType) Option {
	return func(c *config) {
		c.signType = signType
	}
}
func WithSignVersion(signVersion SignVersion) Option {
	return func(c *config) {
		c.signVersion = signVersion
	}
}

type Submail struct {
	auth   *auth
	client *resty.Client
}

func NewSubmail(appid string, appkey string, options ...Option) *Submail {
	conf := defaultConfig()
	for _, option := range options {
		option(conf)
	}

	return &Submail{
		auth: &auth{
			appid:       appid,
			appkey:      appkey,
			tag:         conf.tag,
			signType:    conf.signType,
			signVersion: conf.signVersion,
		},
		client: resty.New(),
	}
}

func (s *Submail) Sign(message Message) {
	s.auth.sign(message)
}

func (s *Submail) Client() *resty.Request {
	return s.client.R()
}
