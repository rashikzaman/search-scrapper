package scheduler

import (
	"context"
	"fmt"
	"rashik/search-scrapper/app/domain"
	"time"
)

func ScheduleKeywordParser(ctx context.Context, repo domain.KeywordRepository) {
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for range ticker.C {
			result, err := repo.FetchPendingKeyword(ctx)
			if err != nil {
				fmt.Println("error", err)
			} else {
				fmt.Println("keyword", result.Word)
			}
		}
	}()
}
