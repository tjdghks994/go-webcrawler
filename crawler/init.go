package crawler

import (
	"fmt"
	"net/http"
)

// Crawler static web site crawling
type Crawler interface {
	Crawling(string)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		fmt.Println("failed : ", res.StatusCode)
	}
}
