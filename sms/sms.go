/**
 * @Author: Jacky
 * @Description:
 * @File: sms
 * @Version: 1.0.0
 * @Date: 2022/4/27 14:30
 */
package sms

import (
	"context"

	"github.com/Carrotxyy/submail-sdk"
)

type Sms struct {
	s *submail.Submail
}

func NewSms(s *submail.Submail) *Sms {
	return &Sms{
		s: s,
	}
}

func (sms *Sms) Send(ctx context.Context, sendRequest *SendRequest) (error, *SendResponse) {
	sendResponse := &SendResponse{}

	sms.s.Sign(sendRequest)
	_, err := sms.s.Client().
		SetContext(ctx).
		SetHeader(submail.JsonContentType()).
		SetBody(sendRequest).
		SetResult(sendResponse).
		Post(sendURL)

	return err, sendResponse
}

func (sms *Sms) XSend(ctx context.Context, xsendRequest *XSendRequest) (error, *XSendResponse) {
	xsendResponse := &XSendResponse{}

	sms.s.Sign(xsendRequest)
	_, err := sms.s.Client().
		SetContext(ctx).
		SetHeader(submail.JsonContentType()).
		SetBody(xsendRequest).
		SetResult(xsendResponse).
		Post(xsendURL)

	return err, xsendResponse

}
