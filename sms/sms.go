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

type SMS struct {
	s *submail.Submail
}

func NewSMS(s *submail.Submail) *SMS {
	return &SMS{
		s: s,
	}
}

// Send
// @Description: 短信发送:https://www.mysubmail.com/documents/FppOR3
// @receiver sms
// @param ctx
// @param sendRequest
// @return error
// @return *SendResponse
func (sms *SMS) Send(ctx context.Context, sendRequest *SendRequest) (error, *SendResponse) {
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

// XSend
// @Description: 短信模板发送:https://www.mysubmail.com/documents/OOVyh
// @receiver sms
// @param ctx
// @param xsendRequest
// @return error
// @return *XSendResponse
func (sms *SMS) XSend(ctx context.Context, xsendRequest *XSendRequest) (error, *XSendResponse) {
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

// MultiSend
// @Description: 短信一对多发送:https://www.mysubmail.com/documents/KZjET3
// @receiver sms
// @param ctx
// @param multiSendRequest
// @return error
// @return []MultiSendResponse
func (sms *SMS) MultiSend(ctx context.Context, multiSendRequest *MultiSendRequest) (error, []MultiSendResponse) {
	var multiSendResponseList []MultiSendResponse

	sms.s.Sign(multiSendRequest)
	_, err := sms.s.Client().
		SetContext(ctx).
		SetHeader(submail.JsonContentType()).
		SetBody(multiSendRequest).
		SetResult(&multiSendResponseList).
		Post(multiSendURL)

	return err, multiSendResponseList
}

// MultiXSend
// @Description: 短信模板一对多发送:https://www.mysubmail.com/documents/eM4rY2
// @receiver sms
// @param ctx
// @param multiXSendRequest
// @return error
// @return []MultiXSendResponse
func (sms *SMS) MultiXSend(ctx context.Context, multiXSendRequest *MultiXSendRequest) (error, []MultiXSendResponse) {
	var multiXSendResponseList []MultiXSendResponse

	sms.s.Sign(multiXSendRequest)
	_, err := sms.s.Client().
		SetContext(ctx).
		SetHeader(submail.JsonContentType()).
		SetBody(multiXSendRequest).
		SetResult(&multiXSendResponseList).
		Post(multiXSendURL)

	return err, multiXSendResponseList
}

// BatchSend
// @Description: 短信批量群发:https://www.mysubmail.com/documents/AzD4Z4
// @receiver sms
// @param ctx
// @param batchSendRequest
// @return error
// @return *BatchSendResponse
func (sms *SMS) BatchSend(ctx context.Context, batchSendRequest *BatchSendRequest) (error, *BatchSendResponse) {
	batchSendResponse := &BatchSendResponse{}

	sms.s.Sign(batchSendRequest)
	_, err := sms.s.Client().
		SetContext(ctx).
		SetHeader(submail.JsonContentType()).
		SetBody(batchSendRequest).
		SetResult(batchSendResponse).
		Post(batchSendURL)

	return err, batchSendResponse
}

// BatchXSend
// @Description: 短信批量模板群发:https://www.mysubmail.com/documents/G5KBR
// @receiver sms
// @param ctx
// @param batchXSendRequest
// @return error
// @return *BatchXSendURLResponse
func (sms *SMS) BatchXSend(ctx context.Context, batchXSendRequest *BatchXSendURLRequest) (error, *BatchXSendURLResponse) {
	batchXSendResponse := &BatchXSendURLResponse{}

	sms.s.Sign(batchXSendRequest)
	_, err := sms.s.Client().
		SetContext(ctx).
		SetHeader(submail.JsonContentType()).
		SetBody(batchXSendRequest).
		SetResult(batchXSendResponse).
		Post(batchXSendURL)

	return err, batchXSendResponse
}
