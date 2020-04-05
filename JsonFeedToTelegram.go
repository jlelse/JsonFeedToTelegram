package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	lastArticleFile, lastArticleFileSet := os.LookupEnv("LAST_ARTICLE_FILE")
	feed, feedSet := os.LookupEnv("FEED")
	botToken, botTokenSet := os.LookupEnv("BOT_TOKEN")
	channel, channelSet := os.LookupEnv("CHANNEL")
	if lastArticleFileSet && feedSet && botTokenSet && channelSet {
		telegram := Telegram{botToken: botToken, channel: channel}
		http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Wrong HTTP method", http.StatusMethodNotAllowed)
				return
			}
			fmt.Println("Fetch feed: ", time.Now().Format(time.RFC3339))
			article, err := LatestArticle(feed)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if lastArticle := lastArticleUrl(lastArticleFile); lastArticle != article.Url {
				err = telegram.Post(createMessage(article))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				err = updateLastArticleUrl(lastArticleFile, article.Url)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				http.Error(w, errors.New("no new article").Error(), http.StatusInternalServerError)
				return
			}
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		log.Fatal("Not configured")
	}
}

func lastArticleUrl(filename string) string {
	fileContent, _ := ioutil.ReadFile(filename)
	return string(fileContent)
}

func updateLastArticleUrl(filename, url string) error {
	return ioutil.WriteFile(filename, []byte(url), 0644)
}

func createMessage(article *Article) string {
	var message bytes.Buffer
	message.WriteString("🔔 Something new was published")
	message.WriteString("\n\n")
	if article.Title != "" {
		message.WriteString(article.Title)
		message.WriteString("\n\n")
	}
	message.WriteString(article.Url)
	return message.String()
}
