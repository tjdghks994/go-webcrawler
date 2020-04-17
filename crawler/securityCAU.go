package crawler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

// CAUCrawler security CAU News Crwaling
type SecurityCAUCrawler struct {
	cont []securityCAU
}

type securityCAU struct {
	title string
	href  string
}

func (cau *SecurityCAUCrawler) getPage(url string) string {
	res, err := http.Get(url)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	pages, _ := doc.Find(".paging>a").Last().Attr("href")
	page := strings.Split(pages, "page=")

	return page[1]
}

func (cau *SecurityCAUCrawler) getContents(url string, page int, c chan<- []securityCAU) {
	noticeURL := url + "&ctg_cd=&skey=&keyword=&mode=list&page=" + strconv.Itoa(page)
	res, err := http.Get(noticeURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	cont := []securityCAU{}
	contentChan := make(chan securityCAU)

	doc.Find(".al").Each(func(i int, s *goquery.Selection) {
		go cau.putContents(s, contentChan)
	})
	for i := 0; i < doc.Find(".al").Length(); i++ {
		temp := <-contentChan
		cont = append(cont, temp)
	}

	fmt.Println("page", page, "done")
	c <- cont
}

func (cau *SecurityCAUCrawler) putContents(s *goquery.Selection, c chan<- securityCAU) {
	cont := securityCAU{}

	temp := s.Find("a").Text()
	convTemp, _ := iconv.ConvertString(temp, "euc-kr", "utf-8")
	cont.title = convTemp
	cont.href, _ = s.Find("a").Attr("href")

	c <- cont
}

// Crawling c
func (cau *SecurityCAUCrawler) Crawling(url string) {
	c := make(chan []securityCAU)

	pages, _ := strconv.Atoi(cau.getPage(url))

	for i := 0; i < pages; i++ {
		go cau.getContents(url, i+1, c)
	}
	for i := 0; i < pages; i++ {
		temp := <-c
		cau.cont = append(cau.cont, temp...)
	}

}
