package main

import (
	"encoding/xml"
	"fmt"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

var washingtonPost = []byte(`
<sitemapindex>
	<sitemap>
		<loc> https://www.google.com </loc>
	</sitemap>
	<sitemap>
		<loc> https://www.google.com </loc>
	</sitemap>
	<sitemap>
		<loc> https://www.google.com </loc>
	</sitemap>
</sitemapindex>
`)

func main() {
	//resp, _:= http.Get("https://www.washingtonpost.com/news-opinions-sitemap.xml")
	//bytes, _:= ioutil.ReadAll(resp.Body)
	bytes := washingtonPost
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	//resp.Body.Close()
}
