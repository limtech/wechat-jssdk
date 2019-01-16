package jssdk

import (
	"fmt"
	"time"
)

const (
	SIGNATURE_ENCODE_STRING string = "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
)

func (self *WechatJSSDK) Signature(url string, ticket string) string {
	noncestr := RandomString()
	timestamp := time.Now().Unix()
	return fmt.Sprintf(SIGNATURE_ENCODE_STRING, ticket, noncestr, timestamp, url)
}
