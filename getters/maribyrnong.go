package maribyrnong

import "github.com/gocolly/colly"

func maribyrnong() string {
	url := "https://www.maribyrnong.vic.gov.au/About-us/Council-and-committee-meetings/Agendas-and-minutes"
	c := colly.NewCollector()

	c.Visit(url)

	link := ""
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link = e.Attr("href")
	})

	return link
	
}