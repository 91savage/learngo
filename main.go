package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.mbn.co.kr/news/politics/"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "?page=" + strconv.Itoa((page)) + "&vod=&category=politics"
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tit")
	fmt.Println(res.Body)

	searchCards.Each(func(i int, s *goquery.Selection) { // s means each cards
		id, _ := s.Attr("href")
		fmt.Println(id)
	})

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL + "?page=1&vod=&category=politics")
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".paging_2018").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()

	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with Status: ", res.StatusCode)
	}
}
