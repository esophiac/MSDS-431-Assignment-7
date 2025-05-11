package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// struct to hold all of the data get from the website
type Article struct {
	Url, Title, Content string
}

// scrapes a list of wikipedia articles and returns a struct of article information
func scrapeWikis(urlList []string) (finalList []Article) {

	// set up the colly collector
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnRequest(func(r *colly.Request) {
		// print URL of request
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	// iterate through the list of URLs and scrape them
	for _, url := range urlList {
		// create an article to hold the data
		var article Article

		c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
			article.Url = url
			article.Content = e.Text
			article.Title = e.Name
		})

		c.Visit(url)

		finalList = append(finalList, article)
	}

	return finalList
}

func main() {

	// from assignment description
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)"}

	wikiList := scrapeWikis(urls)

	fmt.Println(len(wikiList))

}
