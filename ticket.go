package jssdk

import (
	"encoding/json"
	"fmt"
)

const (
	TICKET_API_URL string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

type TicketData struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func (self *WechatJSSDK) GetTicket() (TicketData, error) {
	// get access token
	accessToken, err := self.GetAccessToken()
	if err != nil || accessToken.ErrCode != 0 {
		return TicketData{
			ErrCode: accessToken.ErrCode,
			ErrMsg:  accessToken.ErrMsg,
		}, err
	}

	// get ticket
	return self.GetTicketByToken(accessToken.AccessToken)
}

func (self *WechatJSSDK) GetTicketByToken(accessToken string) (TicketData, error) {
	// get remote data
	res, err := HttpGet(fmt.Sprintf(TICKET_API_URL, accessToken))
	if err != nil {
		return TicketData{}, err
	}

	// parse json data
	var data TicketData
	if err := json.Unmarshal(res, &data); err != nil {
		return data, err
	}
	// on success

	self.Ticket = data.Ticket
	return data, nil
}
