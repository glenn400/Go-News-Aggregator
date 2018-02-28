package main

/*
Had to comment out in order to run main in other file as this file
was used to test the original news aggregator
the functions that were in this file are now in the webapp

func main() {
	// instances of structures
	var n News
	var s UrlIndex
	news_map := make(map[string]NewsMap)

	// if you want to pull info from a website
	resp, _ := http.Get("https://hiphopdx.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywards[idx], n.Locations[idx]}
			//fmt.Println(idx)
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
*/
