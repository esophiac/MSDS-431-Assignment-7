# Week 7 Assignment: Crawling and Scraping the Web
For this assignment, we tested the web scraping abilities of Go on ten wikipedia pages. We used the Colly package instead of the net/http package in the Go standard library. Our goal was to get results that were similar to the Scrapy program in Python, which we succeeded at. It took the Scrapy program 15.77 seconds to retrieve all ten websites, but it only took 0.236 seconds to accomplish the same task. Our deliverable was a JSON lines file (output.jl) which is created in the same directory as the main.go file.

## Background
This assignment used the [Colly Package](https://github.com/gocolly/colly), which is a batteries-included web scraping package for Go.  <br>

For the python portion of the assignment, we used the example provided on Canvas for Scrapy but added timing information. The example can be found [here](https://canvas.northwestern.edu/courses/231183/pages/python-slash-scrapy-wikipedia-example?wrap=1). 

## Recommendation to Management
Management's goal is to create a knowledge base, beginning by scraping websites. For this trial, we used scraping ten wikipedia pages as a test. This trail found that Go vastly out performs the Scrapy method. Additionally, the simplicity of the Go code is another point in the language's favor. The program can easily be extended with additional websites and run without worrying about changes to the directory structure.

## Roles of Programs and Data
### The Data
The data was provided in the assignment description on Canvas as a Go slice:

// Wikipedia URLs for topic of interest  <br>
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
    "https://en.wikipedia.org/wiki/Android_(robot)"
}

### Programs
All Go programs used in this test are found in the Go folder in this repository.  <br>
go.mod: defines the module's properties  <br>
go.sum: record of the library the project depends on  <br>
main_test.go: tests and benchmarks the fuctions in the main.go file  <br>
main.go: function and execution of code to scrape the ten wikipedia sites  <br>
output.jl: output of the main.go file when it is run  <br>
README.md: the readme file for the repository  <br>

## Application
An executable for this project was created using Windows. To create your own executable, run go build in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

## Use of AI
AI was not used for this assignment.