package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	data gin.H
)

func InitCovid() {
	res, err := http.Get("https://koronavirus.gov.hu")
	if err != nil {
		log.Fatalf("Error downloading page!\n%s", err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error downloading data, non 200 status code: %o", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Error parsing data!\n%s", err.Error())
	}

	infP, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-fertozott-pest").Text(), " ", ""), 10, 32)
	infV, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-fertozott-videk").Text(), " ", ""), 10, 32)

	vacc, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-beoltottak").Text(), " ", ""), 10, 32)

	recP, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-gyogyult-pest").Text(), " ", ""), 10, 32)
	recV, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-gyogyult-videk").Text(), " ", ""), 10, 32)

	deadP, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-elhunyt-pest").Text(), " ", ""), 10, 32)
	deadV, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-elhunyt-videk").Text(), " ", ""), 10, 32)

	isol, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-karantenban").Text(), " ", ""), 10, 32)
	samp, _ := strconv.ParseInt(strings.ReplaceAll(doc.Find("#api-mintavetel").Text(), " ", ""), 10, 32)

	data = gin.H{
		"_comment": "Covid data provided by web scraping. https://github.com/EdwardBot/data-source",
		"infected": gin.H{
			"capital":     infP,
			"countryside": infV,
		},
		"recovered": gin.H{
			"capital":     recP,
			"countryside": recV,
		},
		"died": gin.H{
			"capital":     deadP,
			"countryside": deadV,
		},
		"vaccinated": vacc,
		"quarantine": isol,
		"sampled":    samp,
	}
}

func GetCovidData() gin.H {
	return data
}
