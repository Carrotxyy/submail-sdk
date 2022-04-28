/**
 * @Author: Jacky
 * @Description:
 * @File: sms_test
 * @Version: 1.0.0
 * @Date: 2022/4/27 17:17
 */
package sms

import (
	"context"
	"os"
	"testing"

	"github.com/Carrotxyy/submail-sdk"
)

var _sms *SMS

func TestMain(m *testing.M) {
	_s := submail.NewSubmail(
		"75738",
		"e7f9d4e7bdf66b13c0c3d1eebb317f18",
		submail.WithSignType(submail.MD5),
		submail.WithSignVersion(submail.Version2),
	)
	_sms = NewSMS(_s)
	os.Exit(m.Run())
}

func TestSMS_XSend(t *testing.T) {

	vars := map[string]string{
		"v_name": "eryang",
		"v_time": "2022/04/28 08:00:00",
		"url":    "info",
		"code":   "07325",
	}

	err, xsendResponse := _sms.XSend(context.TODO(), NewXSendRequest("15220091565", "5mEYH3", vars))
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", xsendResponse)
	}
}

func TestSMS_MultiXSend(t *testing.T) {
	mulit := []MultiXSendFieldMulti{
		{
			To: "15220091565",
			Vars: map[string]string{
				"v_name": "肖二阳",
				"v_time": "2022/04/29 08:00:00",
				"url":    "test",
				"code":   "07325",
			},
		},
		{
			To: "18938632226",
			Vars: map[string]string{
				"v_name": "李诚",
				"v_time": "2022/04/29 08:00:00",
				"url":    "test",
				"code":   "07326",
			},
		},
	}

	request := NewMultiXSendRequest("5mEYH3", mulit)

	err, response := _sms.MultiXSend(context.TODO(), request)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", response)
	}
}
