package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Crawler() {
	// Create file
	file, err := os.Create("info.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	str := "https://www.techinasia.com/startup?&page="
	fmt.Println("Start crawling")
	for i := 0; i < 100; i++ {
		count := 501
		fmt.Println("Page ", i+1)
		count = count + i
		url := str + strconv.Itoa(count)

		// Query page to chose element
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}

		// Find element
		doc.Find(".datatable .datatable-cell .datatable-cell-block .datatable-cell-title").Each(func(j int, s *goquery.Selection) {

			// Line 1
			number := strconv.Itoa(count) + "-" + strconv.Itoa(j+1) + "\n"
			file.WriteString(number)

			// Line 2
			line2 := "Start Up: " + s.Text() + "\n"
			file.WriteString(line2)

			// Line 3
			line3, flag := s.Attr("href")
			if flag {
				line3 = "Page: " + line3 + "\n"
				file.WriteString(line3)
			}

			nextLink, ok := s.Attr("href")
			if ok {
				doc2, err := goquery.NewDocument(nextLink)
				if err != nil {
					log.Fatal(err)
				}

				s1 := doc2.Find(".company-profile-main .box .company-card .media-body .company-website")
				// Line 4
				line4, ok := s1.Attr("href")
				if ok {
					line4 = "Website: " + line4 + "\n"
					file.WriteString(line4)
				}

				doc2.Find(".company-profile-main .box .media").Each(func(k int, s2 *goquery.Selection) {
					s2.Find(".media-body").Each(func(l int, s3 *goquery.Selection) {
						// Line 5
						if k > 0 {
							line5 := s3.Find(".title").Text() + ": " + s3.Find(".name").Text() + "\n"
							file.WriteString(line5)
						}
					})
				})
			}
			file.WriteString("\n")
		})
	}
}

func main() {
	Crawler()
}
