package domain

import (
	"context"
	"io"
	"time"
)

type Keyword struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Word         string    `json:"word"`
	Adword       string    `json:"adword"`
	Link         string    `json:"link"`
	SearchResult string    `json:"search_result"`
	Status       string    `json:"status"`
	UserId       uint      `json:"user_id"`
	User         User      `json:"-" gorm:"foreignKey:UserId;references:ID"`
	HtmlFilePath string    `json:"html_file_path"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type KeywordRepository interface {
	StoreKeywords(context.Context, [][]string, User) ([]*Keyword, error)
	FetchKeywordsForUser(context.Context, User) ([]*Keyword, error)
	FetchPendingKeyword(ctx context.Context) (*Keyword, error)
	UpdateKeyword(context.Context, uint, string, string, string, string, string) error
}

type KeywordUseCase interface {
	StoreKeywordsFromFile(context.Context, io.Reader, int) ([]*Keyword, error)
	FetchKeywordsForUser(context.Context, int) ([]*Keyword, error)
}
