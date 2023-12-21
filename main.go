package main

import (
	"fmt"
	"log"

	"net/http"
	// "net/smtp"
	// "os/exec"
	// "strings"

	"github.com/gocolly/colly/v2"
	// "github.com/pdfcpu/pdfcpu/pkg/api"
	// "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	fmt.Println("hello world")
	url := "https://www.maribyrnong.vic.gov.au/About-us/Council-and-committee-meetings/Agendas-and-minutes"

	page, response := http.Get(url)
	// fmt.Println(page)
	fmt.Println(response)


	c := colly.NewCollector()

	c.CheckHead = true

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
			
	// 	fmt.Println(link)
	// })

	err := c.Visit(url)
	if err != nil {
			log.Fatal(err)
	}
}
