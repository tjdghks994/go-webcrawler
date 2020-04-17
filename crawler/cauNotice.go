package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type CAUNotice struct {
	title string
}

func (cau *CAUNotice) getHTML(url string) {
	res, err := http.Get(url)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	fmt.Print(string(data))
}

func (cau *CAUNotice) Crawling(url string) {
	cau.getHTML(url)
}
