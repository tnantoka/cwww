package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Chatwork struct {
	Settings
}

func (chatwork Chatwork) PostMessage(body string) {
	urlString := "https://api.chatwork.com/v1/rooms/" + chatwork.RoomID + "/messages"

	values := url.Values{}
	values.Add("body", body)

	req, _ := http.NewRequest("POST", urlString, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-ChatWorkToken", chatwork.APIToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
}
