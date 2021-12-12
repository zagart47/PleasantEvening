package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"strings"
	"time"
)

type Films struct {
	Title string
}

var filmCollection = []Films{}

func main() {
	scrapPage("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
	rand.Seed(time.Now().UnixNano())
	output := filmCollection[rand.Intn(250)]
	fmt.Println(output)
	fmt.Scanln()

}
func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		workString := e.DOM.Find("td:nth-child(2)").Text()
		var film string
		workString = strings.ReplaceAll(workString, "\n", " ")
		workString = strings.Replace(workString, " ", "", 2)
		workString = strings.Replace(workString, "*. ", "", 2)
		workString = strings.ReplaceAll(workString, "       ", " ")
		workString = strings.ReplaceAll(workString, "   ", " ")
		workString = strings.ReplaceAll(workString, "   ", "")
		film = workString
		filmCollection = append(filmCollection, Films{film})
	})

	c.Visit(url)
}
