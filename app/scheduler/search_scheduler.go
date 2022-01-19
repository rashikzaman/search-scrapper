package scheduler

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"rashik/search-scrapper/app/domain"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	TotalSearchResult string
	TotalAdword       string
	TotalLink         string
	HtmlBody          []byte
}

func ScheduleKeywordParser(ctx context.Context, repo domain.KeywordRepository) {
	userAgents, err := readLines("./user_agents.txt")
	if err != nil {
		fmt.Println("Cannot read user agents file", userAgents)
		return
	}
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			result, err := repo.FetchPendingKeyword(ctx)
			if err != nil {
				fmt.Println("error fetching keywords", err)
			} else if result != nil {
				userAgent := userAgents[rand.Intn(len(userAgents))]
				keyword := result.Word
				searchResult, err := GetSearchResult(keyword, userAgent)
				if err != nil {
					fmt.Println("Err", err)
				} else {
					fmt.Println("Result", searchResult.TotalSearchResult, searchResult.TotalAdword, searchResult.TotalAdword, searchResult.TotalLink)
					fmt.Println("User agent", userAgent)
					filepath := fmt.Sprintf("./public/results/result_%d.html", result.ID)
					htmlFilePath := strings.TrimPrefix(filepath, ".")
					err := repo.UpdateKeyword(ctx, result.ID, "completed", searchResult.TotalSearchResult, searchResult.TotalAdword, searchResult.TotalLink, htmlFilePath)
					if err != nil {
						fmt.Println("Error updating keyword", err)
					} else {
						err := storeHtml(filepath, searchResult.HtmlBody)
						if err != nil {
							fmt.Println("error creating html file", err)
						}
					}
				}
			}
		}
	}()
}

func GetSearchResult(keyword string, userAgent string) (*SearchResult, error) {
	data, err := GetHtml(keyword, userAgent)

	if err != nil {
		fmt.Printf("Cannnot read html: %d", err)
		return nil, err
	}

	reader := bytes.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Printf("Cannot create html document, err:%d", err)
		return nil, err
	}

	strlink := strconv.Itoa(doc.Find("a").Length())
	strAds := strconv.Itoa(doc.Find("#tads > div").Length())
	result := &SearchResult{
		TotalSearchResult: doc.Find("#result-stats").Text(),
		TotalAdword:       strAds,
		TotalLink:         strlink,
		HtmlBody:          data,
	}

	return result, nil
}

func GetHtml(keyword string, userAgent string) ([]byte, error) {
	client := &http.Client{}
	keyword = strings.ReplaceAll(keyword, " ", "+") //replaincing all whitespace with +
	fmt.Println("keyword", keyword)
	req, err := http.NewRequest("GET", "https://www.google.ru/search?q="+keyword+"&hl=en", nil)
	if err != nil {
		return nil, nil
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	return body, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func storeHtml(path string, body []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err2 := f.Write(body)
	return err2
}
