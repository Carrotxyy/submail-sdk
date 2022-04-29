# Submail
对接短信服务提供商 **SUBMAIL** 开放接口！

## 下载
```go
go get github.com/Carrotxyy/submail-sdk
```

## 实例

```go
func main(){
	// 密钥明文验证模式
	core := submail.NewSubmail("appid", "appkey")
   	// 数字签名验证模式 
   	//core := submail.NewSubmail(
	//	"appid",
	//	"appkey",
	//	submail.WithSignType(submail.MD5),
	//	submail.WithSignVersion(submail.Version2),
	//)

	// 创建 SMS
	smsCli := sms.NewSMS(core)

	// 创建消息体
	vars := map[string]string{
		"name": "肖二阳",
		"time": "2022/04/28 12:00:00",
	}
	request := sms.NewXSendRequest("phone number", "template id", vars)

	// 发送模板消息
	err, response := smsCli.XSend(context.TODO(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.Code != submail.SUCCESS {
		fmt.Println(response)
		return
	}
}
```

