package jssdk

import (
	"fmt"
	"time"
)

const (
	SIGNATURE_ENCODE_STRING string = "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
)

type JSSDK_CONFIG struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
}

func (self *WechatJSSDK) GetSignature(url string, ticket string) string {
	self.NonceStr = RandomString()
	self.Timestamp = time.Now().Unix()
	self.Signature = fmt.Sprintf(SIGNATURE_ENCODE_STRING, ticket, self.NonceStr, self.Timestamp, url)
	return self.Signature
}

func (self *WechatJSSDK) GetConfig() JSSDK_CONFIG {
	return JSSDK_CONFIG{
		AppID:     self.AppID,
		Timestamp: self.Timestamp,
		NonceStr:  self.NonceStr,
		Signature: self.Signature,
	}
}
