package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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

// Scrape MBN by a term
func Scrape(term string) {
	var baseURL string = "https://www.mbn.co.kr/news/" + term + "/"
	fmt.Println(baseURL)
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL, term) - 5 // MBN 페이징은 10까지만 돼있음 & 1page만 이상하게 tit 하위에 a href id 값이 다름. 그래서 2 page 부터 10page 까지만 진행함.

	for i := 2; i <= totalPages; i++ {
		go getPage(i, baseURL, c, term)
	}

	for i := 2; i <= totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob, term string) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "?page=" + strconv.Itoa((page)) + "&vod=&category=" + term // Itoa 는 숫자를 텍스트로 변환
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".article_list")

	searchCards.Each(func(i int, s *goquery.Selection) { // s means each cards
		go extractJob(s, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)

	}

	mainC <- jobs

}

func extractJob(s *goquery.Selection, c chan<- extractedJob) {
	id, _ := s.Find(".tit>a").Attr("href")
	title := CleanString(s.Find(".tit>a").Text())
	summary := CleanString(s.Find(".desctxt>a").Text())
	date := CleanString(s.Find(".date").Text())
	writer := CleanString(s.Find(".name_b").Text())

	c <- extractedJob{
		id:      id,
		title:   title,
		summary: summary,
		date:    date,
		writer:  writer}
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

}

func getPages(url string, term string) int {
	pages := 0
	res, err := http.Get(url + "?page=2&vod=&category=" + term)
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

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Summary", "Date", "Writer"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.summary, job.date, job.writer}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

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
