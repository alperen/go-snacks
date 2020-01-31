package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type jobRecord struct {
	arguments map[string]string
	result    string
	rid       int
}

type wikipediaResponse struct {
	Title   string `json:"displaytitle"`
	Content string `json:"extract"`
}

type resultRecord struct {
	Tile    string `json:"title"`
	Content  string `json:"content"`
	Language string `json:"language"`
}

var (
	headlines = []string{"MongoDB", "MySQL", "PHP", "JavaScript"}
	languages = []string{"tr", "fr", "en"}
)

const (
	outputFilePath = "./result.json"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateJobQueue() (result []jobRecord) {
	for _, headline := range headlines {
		for _, language := range languages {

			arguments := make(map[string]string)
			arguments["headline"] = headline
			arguments["language"] = language

			result = append(result, jobRecord{arguments: arguments, rid: rand.Intn(1000)})
		}
	}

	return
}

func worker(id int, jobs <-chan jobRecord, result chan<- jobRecord) {
	log.Printf("%d[th] worker started \n", id)

	for job := range jobs {
		headline, language := job.arguments["headline"], job.arguments["language"]
		log.Printf("Job Sended %s - %s \n", headline, language)

		job.result = fetchWikipediaSummary(headline, language)

		result <- job
	}
}

func fetchWikipediaSummary(headline string, language string) (summary string) {
	requestURI := fmt.Sprintf("https://%s.wikipedia.org/api/rest_v1/page/summary/%s", language, headline)
	res, _ := http.Get(requestURI)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var beautyResponse wikipediaResponse
	_ = json.Unmarshal(body, &beautyResponse)

	summary = beautyResponse.Content
	return
}

func main() {
	jobQueue := generateJobQueue()
	jobQueueSize := len(jobQueue)

	jobs := make(chan jobRecord, jobQueueSize)
	results := make(chan jobRecord, jobQueueSize)

	defer close(jobs)
	defer close(results)

	// create workers

	for w := 1; w <= jobQueueSize; w++ {
		go worker(w, jobs, results)
	}

	// send jobs

	for j := 0; j < jobQueueSize; j++ {
		jobs <- jobQueue[j]
	}

	// wait results

	var apiResults []resultRecord

	for a := 1; a <= jobQueueSize; a++ {
		rr := <-results
		headline := rr.arguments["headline"]
		language := rr.arguments["language"]
		content := rr.result

		apiResults = append(apiResults, resultRecord{headline, language, content})

		log.Println("result reached.")
	}

	jsonByte, _ := json.MarshalIndent(&apiResults, "", "	")

	_ = ioutil.WriteFile(outputFilePath, jsonByte, os.ModePerm)
}
