package crawler

import (
	"strconv"
	"testing"
)

func Testcrawl(t *testing.T) {
	testcases := []struct {
		url string
	}{
		{"http://vnexpress.net/rss/tin-moi-nhat.rss"},
		{"http://www.tienphong.vn/rss/the-thao-11.rss"},
		{"http://www.tienphong.vn/rss/cong-nghe-45.rss"},
		{"http://www.tienphong.vn/rss/giao-duc-71.rss"},
	}
	for i, c := range testcases {
		tempFileName := "news_" + strconv.Itoa(i)
		body := crawl(c.url)
		writeToXML(tempFileName, body)
	}
}
