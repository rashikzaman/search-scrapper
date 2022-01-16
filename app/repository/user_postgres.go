package repository

import (
	"context"
	"rashik/search-scrapper/app/domain"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &MysqlUserRepository{Conn}
}

func (m *MysqlUserRepository) FetchUserById(ctx context.Context) {

}

func (m *MysqlUserRepository) FetchUserByEmail(ctx context.Context) {

}

func (m *MysqlUserRepository) StoreUser(ctx context.Context, data *domain.User) (*domain.User, error) {
	result := m.Conn.Create(data)
	return data, result.Error
}
