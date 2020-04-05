package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func LatestArticle(url string) (*Article, error) {
	jsonFeed := &struct {
		Items []Article `json:"items"`
	}{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("failed to create req to get json feed")
	}
	req.Header.Add("User-Agent", "JsonFeedToTelegram")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("failed to get json feed")
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonFeed)
	if err != nil {
		return nil, errors.New("failed to parse json feed")
	}
	if len(jsonFeed.Items) < 1 {
		return nil, errors.New("no articles in feed")
	}
	return &jsonFeed.Items[0], nil
}
