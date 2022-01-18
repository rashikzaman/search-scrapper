package usecase

import (
	"context"
	"encoding/csv"
	"io"
	"rashik/search-scrapper/app/domain"
)

type KeywordRepository struct {
	KeywordRepository domain.KeywordRepository
}

func NewKeywordUseCase(a domain.KeywordRepository) domain.KeywordUseCase {
	return &KeywordRepository{
		KeywordRepository: a,
	}
}

func (m *KeywordRepository) FetchKeywordsForUser(ctx context.Context, id int) ([]*domain.Keyword, error) {
	//user, err := m.FetchUserById(ctx, id)
	//return user, err
	return nil, nil
}

func (m *KeywordRepository) StoreKeywordsFromFile(ctx context.Context, file io.Reader, userId int) ([]*domain.Keyword, error) {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) > 0 {
		result, err := m.KeywordRepository.StoreKeywords(ctx, records, uint(userId))
		return result, err
	}
	return nil, nil
}
