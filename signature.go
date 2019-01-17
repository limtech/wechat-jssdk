package jssdk

import (
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	SIGNATURE_ENCODE_STRING string = "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
)

func (self *WechatJSSDK) GetSignature(url string, ticket string) string {
	self.NonceStr = RandomString()
	self.Timestamp = time.Now().Unix()
	self.Signature = fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf(SIGNATURE_ENCODE_STRING, ticket, self.NonceStr, self.Timestamp, url))))
	return self.Signature
}
