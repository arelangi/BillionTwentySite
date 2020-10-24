package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	xj "github.com/basgys/goxml2json"
)

func main() {
	// We can use GET form to get result.
	resp, err := http.Get("http://feeds.feedburner.com/billiontwenty/intheend")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := xj.Convert(resp.Body)
	if err != nil {
		panic(err)
	}

	var feed Feed
	err = json.Unmarshal(response.Bytes(), &feed)
	if err != nil {
		panic(err)
	}

	fmt.Println(feed.Rss.Channel.Item[0].Summary)

	latestPodcast := getLatest(feed)
	mainContent := getMainContent(feed)
	footer := getFooter(feed)
	episodes := getEpisodes(feed)

	indexPage := header + latestPodcast + mainContent + aboutUsTeam + newsletter + footer
	episodePage := header + latestPodcast + episodes + footer
	aboutUsPage := header + aboutUsHeader + aboutUsTeam + newsletter + footer
	contactUsPage := header + contactUs + newsletter + footer

	err = ioutil.WriteFile("docs/index.html", []byte(indexPage), 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("docs/intheend.html", []byte(episodePage), 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("docs/about-us.html", []byte(aboutUsPage), 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("docs/contact.html", []byte(contactUsPage), 0644)
	if err != nil {
		panic(err)
	}

}
