package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	var urls []string
	var url string

	fmt.Print("Enter the URL : ")
	fmt.Scan(&url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find all URLs in the HTML
	re := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9\+&@#/%?=~_|!:,.;]*[-A-Za-z0-9\+&@#/%=~_|]`)
	matches := re.FindAllString(string(body), -1)

	// Separate URLs by file type
	for _, match := range matches {
		if strings.HasSuffix(match, ".php") {
			urls = append(urls, "PHP:\n"+match)
		} else if strings.HasSuffix(match, ".js") {
			urls = append(urls, "JS:\n"+match)
		} else if strings.HasSuffix(match, ".html") {
			urls = append(urls, "HTML:\n"+match)
		} else if strings.HasSuffix(match, ".css") {
			urls = append(urls, "CSS:\n"+match)
		} else if strings.HasSuffix(match, ".jpg") || strings.HasSuffix(match, ".png") || strings.HasSuffix(match, ".gif") {
			urls = append(urls, "Image:\n"+match)
		} else if strings.HasSuffix(match, ".mp3") || strings.HasSuffix(match, ".wav") || strings.HasSuffix(match, ".ogg") {
			urls = append(urls, "Audio:\n"+match)
		} else if strings.HasSuffix(match, ".mp4") || strings.HasSuffix(match, ".avi") || strings.HasSuffix(match, ".mov") {
			urls = append(urls, "Video:\n"+match)
		}
	}

	// Write URLs to a text file
	err = ioutil.WriteFile("urls.txt", []byte(strings.Join(urls, "\n")), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("URLs have been scraped and written to 'urls.txt'.")
}
