/**
 * @Author: Jacky
 * @Description:
 * @File: submail
 * @Version: 1.0.0
 * @Date: 2022/4/27 10:52
 */
package submail

type Submail struct {
	AppID  string
	AppKey string
}

func NewSubmail(appid string, appkey string) *Submail {
	return &Submail{
		AppID:  appid,
		AppKey: appkey,
	}
}
