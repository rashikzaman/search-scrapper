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
	"rashik/search-scrapper/app/logger"
	"rashik/search-scrapper/config"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
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
		logger.GetLog().WithFields(logrus.Fields{
			"error": err,
		}).Error("Cannot read user agents file")
		return
	}
	ticker := time.NewTicker(time.Duration(config.GetConfig().GetSchedulerInterval()) * time.Millisecond)
	go func() { //go routine
		for range ticker.C {
			result, err := repo.FetchPendingKeyword(ctx)
			if err != nil {
				logger.GetLog().WithFields(logrus.Fields{
					"error": err,
				}).Error("Error fetching pending keywords")
			} else if result != nil {
				go scrapResult(ctx, repo, userAgents, result) //spawning go routine
			}
		}
	}()
}

func scrapResult(ctx context.Context, repo domain.KeywordRepository, userAgents []string, result *domain.Keyword) {
	userAgent := userAgents[rand.Intn(len(userAgents))]
	keyword := result.Word
	searchResult, err := getSearchResult(keyword, userAgent)
	if err != nil {
		logger.GetLog().WithFields(logrus.Fields{
			"error": err,
		}).Error("Error getting search result")
	} else {
		printSearchResult(searchResult, userAgent)
		filepath := fmt.Sprintf("./public/results/result_%d.html", result.ID)
		htmlFilePath := strings.TrimPrefix(filepath, ".")
		err := repo.UpdateKeyword(ctx, result.ID, "completed", searchResult.TotalSearchResult, searchResult.TotalAdword, searchResult.TotalLink, htmlFilePath)
		if err != nil {
			logger.GetLog().WithFields(logrus.Fields{
				"error": err,
			}).Error("Error updating pending keywords")
		} else {
			err := storeHtml(filepath, searchResult.HtmlBody)
			if err != nil {
				logger.GetLog().WithFields(logrus.Fields{
					"error": err,
				}).Error("Error creating html file")
			}
		}
	}
}

func printSearchResult(searchResult *SearchResult, userAgent string) {
	logger.GetLog().WithFields(logrus.Fields{
		"search_result": searchResult.TotalSearchResult,
		"total_adword":  searchResult.TotalAdword,
		"total_link":    searchResult.TotalLink,
		"user_agent":    userAgent,
	}).Debug("Scrapper Result")
}

func getSearchResult(keyword string, userAgent string) (*SearchResult, error) {
	data, err := getHtml(keyword, userAgent)
	if err != nil {
		logger.GetLog().WithFields(logrus.Fields{
			"error": err,
		}).Error("Error fetching pending keywords")
		return nil, err
	}
	result, err := parseHtml(data)
	return result, err
}

func parseHtml(data []byte) (*SearchResult, error) {
	reader := bytes.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		logger.GetLog().WithFields(logrus.Fields{
			"error": err,
		}).Error("Cannot create html document")
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

func getHtml(keyword string, userAgent string) ([]byte, error) {
	client := &http.Client{}
	keyword = strings.ReplaceAll(keyword, " ", "+") //replaincing all whitespace with +
	logrus.Debug("Keyword: ", keyword)
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
