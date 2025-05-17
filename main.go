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

// function to get JSON files from wikipedia article API
func scrapeWiki(urlArticle []string) (finalList []Article) {

	// set up the colly collector
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnRequest(func(r *colly.Request) {
		// print URL of request
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	// iterate through and define responses, then visit the api
	for _, entry := range urlArticle {
		// create an article to hold the data
		var article Article

		c.OnResponse(func(r *colly.Response) {

			var msgMapTemplate interface{}
			json.Unmarshal([]byte(r.Body), &msgMapTemplate)
			msgMap := msgMapTemplate.(map[string]interface{})

			article.Title, article.Content = formatWiki(msgMap)
			article.Url = entry

		})

		c.Visit(entry)

		finalList = append(finalList, article)

	}

	return finalList
}

// handling the result of the wikipedia page
func formatWiki(result map[string]interface{}) (string, string) {
	//move back to return

	var step2 map[string]interface{}

	step1 := result["query"].(map[string]interface{})["pages"].(map[string]interface{})

	for _, value := range step1 {
		step2 = value.(map[string]interface{})
	}

	articleTitle := step2["title"].(string)

	articleExtract := step2["extract"].(string)

	return articleTitle, articleExtract

}

// create a json lines file from a slice of articles
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

	reFormat := newURL(urls)

	wikiList := scrapeWiki(reFormat)

	finalResult := createJL(wikiList)

	fmt.Println(finalResult)

}
