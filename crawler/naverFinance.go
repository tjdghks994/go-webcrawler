package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

type NaverFinanceCrawler struct {
	cont naverFinance
}

type naverFinance struct {
	title string
}

func getContent(url string) {
	resp, err := http.Get(url)
	checkErr(err)
	checkStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchBox := doc.Find(".box_type_l")
	searchName := searchBox.Find("tr")
	searchName.Each(func(i int, s *goquery.Selection) {
		if s.Find("a").Text() != "" {
			title := s.Find("a").Text()
			convTitle, _ := iconv.ConvertString(title, "euc-kr", "utf-8")
			num := strings.Fields(strings.TrimSpace(s.Find(".number").Text()))
			fmt.Println(convTitle, num)
		}
	})

}
