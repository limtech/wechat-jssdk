package jssdk

type JSSDK_CONFIG struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
}

func (self *WechatJSSDK) GetConfig(url string, ticket string) JSSDK_CONFIG {
	self.GetSignature(url, ticket)
	return JSSDK_CONFIG{
		AppID:     self.AppID,
		Timestamp: self.Timestamp,
		NonceStr:  self.NonceStr,
		Signature: self.Signature,
	}
}
