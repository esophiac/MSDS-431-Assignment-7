package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// struct to hold all of the data get from the website
type Article struct {
	Url, Title, Content string
}

// format the URL to perform better with the wikipedia API
func newURL(urlList []string) (newList []string) {

	for _, url := range urlList {

		splitURL := strings.Split(url, "/")

		articleName := splitURL[len(splitURL)-1]

		newFront := "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exlimit=max&explaintext&titles="

		finalNew := newFront + articleName

		newList = append(newList, finalNew)

	}

	return newList
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

			// pulling the title from the URL as a test
			splitURL := strings.Split(e.Request.URL.String(), "/")
			article.Title = splitURL[len(splitURL)-1]
		})

		c.Visit(url)

		finalList = append(finalList, article)
	}

	return finalList
}

// get a slice of JSON files from the new list of wikipedia articles
func scrapeWiki(urlList []string) (listJSON []Article) {

	return listJSON
}

func createJL(article []Article) (final string) {
	// open a new file named output.jl
	f, err := os.OpenFile("output.jl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal the slice of articles and add each one as a new entry to the jsonl document
	for _, entry := range article {
		jsonData, err := json.Marshal(entry)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write(jsonData); err != nil {
			log.Fatal(err)
		}
		// write a /n to the document
		if _, err := f.Write([]byte("\n")); err != nil {
			log.Fatal(err)
		}
	}

	// close the new document that was created
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	final = "Your file is ready"

	return final

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

	//wikiList := scrapeWikis(urls)

	//finalResult := createJL(wikiList)

	//fmt.Println(finalResult)

	testUrls := []string{"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot"}

	fmt.Println(newURL(testUrls))

}
