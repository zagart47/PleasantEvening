package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Films struct {
	Rank  int     `json:"Rank"`
	Title string  `json:"Title"`
	Year  string  `json:"Year"`
	Rate  float32 `json:"Rate"`
}

var filmCollection = []Films{}

func main() {
	scrapPage("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

}
func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		workString := e.DOM.Find("td:nth-child(2)").Text()
		var rankFilm string
		workString = strings.ReplaceAll(workString, "\n", " ")
		workString = strings.ReplaceAll(workString, "       ", " ")
		workString = strings.ReplaceAll(workString, "   ", " ")
		workString = strings.ReplaceAll(workString, "   ", "")
		rankFilm = workString
		//titleFilm := e.DOM.Find("td:nth-child(2)").Text()
		//yearFilm := e.DOM.Find("td:nth-child(3)").Text()
		//rateFilm := e.DOM.Find("td:nth-child(3)").Text()
		//fmt.Print(workString)
		fmt.Println(rankFilm)
	})

	c.Visit(url)
}
