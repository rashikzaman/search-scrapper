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

func (m *PostgresKeywordRepository) FetchKeywordsForUser(ctx context.Context, user domain.User, searchKey string) ([]*domain.Keyword, error) {
	var keywords []*domain.Keyword
	var result *gorm.DB
	query := m.Conn.Where("user_id = ?", user.ID)
	if searchKey == "" {
		result = query.Find(&keywords)
	} else {
		result = query.Where("word like ?", "%"+searchKey+"%").Find(&keywords)
	}
	return keywords, result.Error
}

func (m *PostgresKeywordRepository) StoreKeywords(ctx context.Context, keywords [][]string, user domain.User) ([]*domain.Keyword, error) {
	var words = []*domain.Keyword{}
	for i := 0; i < len(keywords); i++ {
		words = append(words, &domain.Keyword{
			Word:         keywords[i][0],
			User:         user,
			Status:       "pending",
			HtmlFilePath: "",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now()})
	}
	result := m.Conn.Create(&words)
	return words, result.Error
}

func (m *PostgresKeywordRepository) FetchPendingKeyword(ctx context.Context) (*domain.Keyword, error) {
	var keyword *domain.Keyword
	result := m.Conn.Where("status = ?", "pending").First(&keyword)
	return keyword, result.Error
}

func (m *PostgresKeywordRepository) UpdateKeyword(ctx context.Context, id uint,
	status string,
	searchResult string,
	adwords string,
	totalLink string,
	htmlFile string) error {

	var keyword *domain.Keyword
	result := m.Conn.First(&keyword, id)

	if result.Error != nil {
		return result.Error
	} else {
		keyword.Adword = adwords
		keyword.SearchResult = searchResult
		keyword.Status = status
		keyword.Link = totalLink
		keyword.HtmlFilePath = htmlFile
		result = m.Conn.Save(&keyword)
		return result.Error
	}
}
