package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type CheckResult struct {
	URL		string
	Status	string
	Title	string
	Error	error
}

func getTitle(doc *html.Node) (string, bool) {
	if doc.Type == html.ElementNode && doc.Data == "title" {
		return doc.FirstChild.Data, true
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		title, ok := getTitle(c)
		if ok {
			return title, true
		}
	}

	return "", false
}

func checkURL(url string) CheckResult {
	result := CheckResult{URL: url}

	resp, err := http.Get(url)
	if err != nil {
		result.Error = err
		return result
	}
	defer resp.Body.Close() 

	result.Status = resp.Status
	if resp.StatusCode != http.StatusOK {
		result.Error = fmt.Errorf("Status code is not OK: %d", resp.StatusCode)
		return result
	}

	limitRecorder := &io.LimitedReader{R: resp.Body, N: 1000000}

	doc, err := html.Parse(limitRecorder)
	if err != nil {
		result.Error = fmt.Errorf("Failed to parsing HTML: %w", err )
		return result
	}

	if title, ok := getTitle(doc); ok {
		result.Title = strings.TrimSpace(title)
	} else {
		result.Title = "Title not found!"
	}
	
	return result
}

func main() {
	urlToCheck := "https://www.tiktok.com"

	fmt.Printf("Checking URL", urlToCheck)
	result := checkURL(urlToCheck)

	if result.Error != nil {
		fmt.Printf("Failed: %+v\n", result)
	} else {
		fmt.Printf("Success: %+v\n", result)
	}
}