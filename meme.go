package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var (
	amount = 20
	memes  []gin.H
)

func FetchMemes() {
	memes = make([]gin.H, 0)
	fetchPart(false, "")
	for i := 0; i < amount-1; i++ {
		log.Printf("Fetching meme part %o/%o", i+2, amount-1)
		fetchPart(false, memes[len(memes)-1]["id"].(string))
	}
}

func fetchPart(first bool, after string) {
	var url string
	if first {
		url = "https://reddit.com/u/bendimester_23/m/meme.json?sort=rising&limit=100"
	} else {
		url = "https://reddit.com/u/bendimester_23/m/meme.json?sort=rising&limit=100?after=" + after
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error fetching memes!\n%s", err.Error())
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	res, err := new(http.Client).Do(req)
	if err != nil {
		log.Fatalf("Error fetching memes!\n%s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	var data map[string]interface{}
	rawData, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		log.Println(string(rawData))
		log.Fatalf("Error fetching memes! Got code: %o", res.StatusCode)
	}
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		log.Fatalf("Error parsing JSON.\n%s", err.Error())
	}

	var posts []interface{}
	posts = data["data"].(map[string]interface{})["children"].([]interface{})

	result := make([]gin.H, 0, 0)

	for _, i := range posts {
		iii := i.(map[string]interface{})["data"]
		ii := iii.(map[string]interface{})

		r := gin.H{
			"subreddit": ii["subreddit"],
			"id":        ii["id"],
			"image":     ii["url_overridden_by_dest"],
			"url":       "https://reddit.com" + ii["permalink"].(string),
			"author":    ii["author"],
			"title":     ii["title"],
		}
		result = append(result, r)
	}

	memes = append(memes, result...)
}

func GetRandomMeme() gin.H {
	return memes[rand.Intn(len(memes))]
}

func GetMemeCount() int {
	return len(memes)
}

func GetAllMemes() []gin.H {
	return memes
}
