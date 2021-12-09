package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Films struct {
	ID    int     `json:"ID"`
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
		var idFilm string
		idFilm = string(workString[0] + workString[1] + workString[2])
		//titleFilm := e.DOM.Find("td:nth-child(2)").Text()
		//yearFilm := e.DOM.Find("td:nth-child(3)").Text()
		//rateFilm := e.DOM.Find("td:nth-child(3)").Text()
		//fmt.Print(workString)
		fmt.Println(idFilm)
	})

	c.Visit(url)
}
