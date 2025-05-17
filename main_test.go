package main

import (
	"testing"

	"fmt"
)

// creating a map[string]interface for used in testing
var testQuery = map[string]any{
	"query": map[string]any{
		"pages": map[string]any{
			"555": map[string]any{
				"pageid":  "555",
				"ns":      0,
				"title":   "Test Title",
				"extract": "test content",
			},
		},
	},
}

// test to make sure the newURL function returns the correct strings
func TestNewURL(t *testing.T) {

	in := []string{"https://en.wikipedia.org/wiki/Robotics", "https://en.wikipedia.org/wiki/Robot"}

	out := newURL(in)

	expected := []string{"https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exlimit=max&explaintext&titles=Robotics",
		"https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exlimit=max&explaintext&titles=Robot"}

	if len(expected) != len(out) {
		t.Errorf("Did not match the expected length")
	}
	for i := range expected {
		if expected[i] != out[i] {
			t.Errorf("Expected %s but got %s", expected[i], out[i])
		}
	}

}

// test to make sure that formattng the wiki works out correctly
func TestFormatWiki(t *testing.T) {
	in := testQuery

	outTitle, outContent := formatWiki(in)

	expectedTitle := "Test Title"

	expectedContent := "test content"

	if outTitle != expectedTitle {
		t.Errorf("Expected %s but got %s", expectedTitle, outTitle)
	}

	if outContent != expectedContent {
		t.Errorf("Expected %s but got %s", expectedContent, outContent)
	}

}

// benchmark for the main.go functions
func BenchmarkScrapeAll(b *testing.B) {
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
