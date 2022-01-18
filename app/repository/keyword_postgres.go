package repository

import (
	"context"
	"rashik/search-scrapper/app/domain"
	"time"

	"gorm.io/gorm"
)

type PostgresKeywordRepository struct {
	Conn *gorm.DB
}

func NewPostgresKeywordRepository(Conn *gorm.DB) domain.KeywordRepository {
	return &PostgresKeywordRepository{Conn}
}

func (m *PostgresKeywordRepository) FetchKeywordsForUser(ctx context.Context, user domain.User) ([]*domain.Keyword, error) {
	var keywords []*domain.Keyword
	result := m.Conn.Where("user_id = ?", user.ID).Find(&keywords)
	return keywords, result.Error
}

func (m *PostgresKeywordRepository) StoreKeywords(ctx context.Context, keywords [][]string, user domain.User) ([]*domain.Keyword, error) {
	var words = []*domain.Keyword{}
	for i := 0; i < len(keywords); i++ {
		words = append(words, &domain.Keyword{
			Word:      keywords[i][0],
			User:      user,
			Status:    "pending",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()})
	}
	result := m.Conn.Create(&words)
	return words, result.Error
}

func (m *PostgresKeywordRepository) FetchPendingKeyword(ctx context.Context) (*domain.Keyword, error) {
	var keyword *domain.Keyword
	result := m.Conn.Where("status = ?", "pending").First(&keyword)
	return keyword, result.Error
}
