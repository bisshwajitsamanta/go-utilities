package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Urlset struct {
	Urls []Url `xml:"url"`
}
type Url struct {
	Loc string `xml:"loc"`
}

func (u Url) String() string {
	return fmt.Sprintf(u.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-opinions-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var u Urlset
	xml.Unmarshal(bytes, &u)
	//u.Urls is the slice containing all the URL's
	for _, location := range u.Urls {
		fmt.Printf("\n%s", location)
	}
	resp.Body.Close()
}
