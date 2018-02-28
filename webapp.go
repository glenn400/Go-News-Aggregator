package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

// sync variable
var wg sync.WaitGroup

// struct for site map at top level
type UrlIndex struct {
	Locations []string `xml:"url>loc"`
}

// new struct for site map inside above site map
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywards  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	// send value to the channel
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// instances of structures
	var s UrlIndex
	news_map := make(map[string]NewsMap)

	// if you want to pull info from a website
	resp, _ := http.Get("https://hiphopdx.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	queue := make(chan News, 1000)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)

	for elem := range queue {
		for idx, _ := range elem.Keywards {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywards[idx], elem.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>I love GO</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8080", nil)

}
