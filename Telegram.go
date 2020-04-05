package main

import (
	"errors"
	"net/http"
	"net/url"
)

type Telegram struct {
	channel  string
	botToken string
}

var telegramBaseUrl = "https://api.telegram.org/bot"

func (t *Telegram) Post(message string) error {
	params := url.Values{}
	params.Add("chat_id", t.channel)
	params.Add("text", message)
	tgUrl, err := url.Parse(telegramBaseUrl + t.botToken + "/sendMessage")
	if err != nil {
		return errors.New("failed to create Telegram request")
	}
	tgUrl.RawQuery = params.Encode()
	req, _ := http.NewRequest(http.MethodPost, tgUrl.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return errors.New("failed to send Telegram message")
	}
	return nil
}
