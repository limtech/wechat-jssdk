package jssdk

type WechatJSSDK struct {
	AppID       string
	AppSecret   string
	AccessToken string
	Ticket      string
	NonceStr    string
	Timestamp   int64
	Signature   string
}

func New(appId string, appSecret string) *WechatJSSDK {
	return &WechatJSSDK{
		AppID:     appId,
		AppSecret: appSecret,
	}
}
