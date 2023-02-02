package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractJob struct {
	id      string
	title   string
	summary string
	date    string
	writer  string
}

var baseURL string = "https://www.mbn.co.kr/news/politics/"

func main() {
	totalPages := getPages() - 5 // MBN 페이징은 10까지만 돼있음 & 1page만 이상하게 tit 하위에 a href id 값이 다름. 그래서 2 page 부터 10page 까지만 진행함.
	fmt.Println(totalPages)

	for i := 2; i <= totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "?page=" + strconv.Itoa((page)) + "&vod=&category=politics" // Itoa 는 숫자를 텍스트로 변환
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tit")

	searchCards.Each(func(i int, s *goquery.Selection) { // s means each cards
		id, _ := s.Find("a").Attr("href")
		// fmt.Println(exists)
		fmt.Println(id)
		// if exists {
		// 	fmt.Println(id)
		// }

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
