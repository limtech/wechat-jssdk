package jssdk

type WechatJSSDK struct {
	AppID       string
	AppSecret   string
	AccessToken string
	Ticket      string
}

func New(appId string, appSecret string) *WechatJSSDK {
	return &WechatJSSDK{
		AppID:     appId,
		AppSecret: appSecret,
	}
}
