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

func (m *PostgresKeywordRepository) FetchKeywordsForUser(ctx context.Context, id int) ([]*domain.Keyword, error) {
	//user, err := m.FetchUserById(ctx, id)
	//return user, err

	return nil, nil
}

func (m *PostgresKeywordRepository) StoreKeywords(ctx context.Context, keywords [][]string, userId uint) ([]*domain.Keyword, error) {
	var words = []*domain.Keyword{}
	for i := 0; i < len(keywords); i++ {
		words = append(words, &domain.Keyword{
			Word:      keywords[i][0],
			UserId:    userId,
			Status:    "pending",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()})
	}
	result := m.Conn.Create(&words)
	return words, result.Error
}
