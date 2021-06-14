package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	infP := doc.Find("#api-fertozott-pest").Text()
	infV := doc.Find("#api-fertozott-videk").Text()

	vacc := doc.Find("#api-beoltottak").Text()

	recP := doc.Find("#api-gyogyult-pest").Text()
	recV := doc.Find("#api-gyogyult-videk").Text()

	deadP := doc.Find("#api-elhunyt-pest").Text()
	deadV := doc.Find("#api-elhunyt-videk").Text()

	isol := doc.Find("#api-karantenban").Text()

	samp := doc.Find("#api-mintavetel").Text()

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
