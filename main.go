package main

import (
	"github.com/tjdghks994/go-webcrawler/crawler"
)

var sss string = "https://www.cau.ac.kr/cms/FR_CON/index.do?MENU_ID=2130#page1"
var cauURL string = "https://www.cau.ac.kr/cms/FR_CON/index.do?MENU_ID=100#page1"
var paging string = "https://www.cau.ac.kr/bvs/pagingFH.do"

func main() {

	// secCAUURL := "http://security.cau.ac.kr/board.htm?bbsid=notice"
	// //naverURL := "https://finance.naver.com/sise/lastsearch2.nhn"

	// c := new(crawler.SecurityCAUCrawler)
	// c.Crawling(secCAUURL)

	CAU := new(crawler.CAUNotice)
	CAU.Crawling(sss)

}
