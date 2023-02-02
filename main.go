package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id      string
	title   string
	summary string
	date    string
	writer  string
}

var baseURL string = "https://www.mbn.co.kr/news/entertain/"

func main() {
	var jobs []extractedJob
	totalPages := getPages() - 5 // MBN 페이징은 10까지만 돼있음 & 1page만 이상하게 tit 하위에 a href id 값이 다름. 그래서 2 page 부터 10page 까지만 진행함.

	for i := 2; i <= totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "?page=" + strconv.Itoa((page)) + "&vod=&category=entertain" // Itoa 는 숫자를 텍스트로 변환
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".article_list")

	searchCards.Each(func(i int, s *goquery.Selection) { // s means each cards
		job := extractJob(s)
		jobs = append(jobs, job)
	})
	return jobs

}

func extractJob(s *goquery.Selection) extractedJob {
	id, _ := s.Find(".tit>a").Attr("href")
	title := cleanString(s.Find(".tit>a").Text())
	summary := cleanString(s.Find(".desctxt>a").Text())
	date := cleanString(s.Find(".date").Text())
	writer := cleanString(s.Find(".name_b").Text())

	return extractedJob{
		id:      id,
		title:   title,
		summary: summary,
		date:    date,
		writer:  writer,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL + "?page=2&vod=&category=entertain")
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
