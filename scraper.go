package main

import (
  "log"
  "github.com/gocolly/colly/v2"
)

type Lession struct {
  Title       string
  Description string
  YoutubeURL  string
  Duration    string
  Transcript  string
}

type Section struct {
  Title    string
  Lessions []Lession
}

func main()  {
  webpage := "https://www.professormesser.com/security-plus/sy0-701/sy0-701-video/sy0-701-comptia-security-plus-course/"
  collector := colly.NewCollector(
    colly.AllowedDomains("professormesser.com", "www.professormesser.com"),
    colly.MaxDepth(1),
  )

  	// Error handling
	collector.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", r.StatusCode, err)
	})

  log.Println("Info: Scrapping", webpage)
  collector.Visit(webpage)
}

