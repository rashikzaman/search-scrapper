package usecase

import (
	"context"
	"encoding/csv"
	"io"
	"rashik/search-scrapper/app/domain"
)

type KeywordRepository struct {
	KeywordRepository domain.KeywordRepository
	UserUseCase       domain.UserUseCase
}

func NewKeywordUseCase(a domain.KeywordRepository, b domain.UserUseCase) domain.KeywordUseCase {
	return &KeywordRepository{
		KeywordRepository: a,
		UserUseCase:       b,
	}
}

func (m *KeywordRepository) FetchKeywordsForUser(ctx context.Context, userId int, searchKey string) ([]*domain.Keyword, error) {
	user, err := m.UserUseCase.FetchUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	keywords, err := m.KeywordRepository.FetchKeywordsForUser(ctx, *user, searchKey)
	return keywords, err
}

func (m *KeywordRepository) StoreKeywordsFromFile(ctx context.Context, file io.Reader, userId int) ([]*domain.Keyword, error) {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) > 0 {
		user, err := m.UserUseCase.FetchUserById(ctx, userId)
		if err != nil {
			return nil, err
		}
		result, err := m.KeywordRepository.StoreKeywords(ctx, records, *user)
		return result, err
	}
	return nil, nil
}
