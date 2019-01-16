package jssdk

import (
	"encoding/json"
	"fmt"
)

const (
	ACCESS_TOKEN_API_URL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type AccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func (self *WechatJSSDK) GetAccessToken() (AccessTokenData, error) {
	var err error
	var data AccessTokenData
	// get remote data
	res, err := HttpGet(fmt.Sprintf(ACCESS_TOKEN_API_URL, self.AppID, self.AppSecret))
	if err != nil {
		return data, err
	}

	// parse json data
	if err := json.Unmarshal(res, &data); err != nil {
		return data, err
	}
	self.AccessToken = data.AccessToken
	return data, nil
}
