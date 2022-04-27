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
	"encoding/json"
	"os"
	"testing"

	"github.com/Carrotxyy/submail-sdk"
)

var _sms *Sms

func TestMain(m *testing.M) {
	_s := submail.NewSubmail(
		"75738",
		"e7f9d4e7bdf66b13c0c3d1eebb317f18",
		submail.WithSignType(submail.MD5),
		submail.WithSignVersion(submail.Version2),
	)
	_sms = NewSms(_s)
	os.Exit(m.Run())
}

func TestSms_XSend(t *testing.T) {

	vars := map[string]string{
		"v_name": "eryang",
		"v_time": "2022/04/28 08:00:00",
		"url":    "info",
		"code":   "07325",
	}

	varsData, _ := json.Marshal(vars)

	err, xsendResponse := _sms.XSend(context.TODO(), &XSendRequest{
		To:      "15220091565",
		Project: "5mEYH3",
		Vars:    string(varsData),
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", xsendResponse)
	}
}
