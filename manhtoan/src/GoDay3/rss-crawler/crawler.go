package main

import (
	"encoding/xml"
	// "flag"
	//"io"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml: "version,attr"`

	//Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`

	//Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []item `xml:"channel>item"`
}

type item struct {
	//Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`

	//Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

func main() {
	testcases := []struct {
		url string
	}{
		{"http://vnexpress.net/rss/tin-moi-nhat.rss"},
		{"http://www.tienphong.vn/rss/the-thao-11.rss"},
		{"http://www.tienphong.vn/rss/cong-nghe-45.rss"},
		{"http://www.tienphong.vn/rss/giao-duc-71.rss"},
	}
	for i, c := range testcases {
		tempFileName := "news_" + strconv.Itoa(i) + ".xml"
		ch := make(chan int)
		go crawl(c.url, tempFileName, ch)
		ch <- 1
	}

	// r := rss2{}
	// xmlContent, _ := ioutil.ReadFile("test.xml")
	// err := xml.Unmarshal(xmlContent, &r)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, item := range r.ItemList {
	// 	fmt.Println(item)
	// }
}

func crawl(uri string, tempFileName string, ch chan int) {
	fmt.Println("Starting crawling ... ")
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	writeToXML(tempFileName, string(body))
	<-ch
}

func writeToXML(filename string, source string) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Write to XML file
	//xmlWriter := io.Writer(file)
	file.WriteString(source)
}
