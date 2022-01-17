package usecase

import (
	"context"
	"rashik/search-scrapper/app/domain"
	"time"
)

type userUserCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(a domain.UserRepository) domain.UserUseCase {
	return &userUserCase{
		UserRepository: a,
	}
}

func (m *userUserCase) FetchUserById(ctx context.Context) {

}

func (m *userUserCase) FetchUserByEmail(ctx context.Context) {

}

func (m *userUserCase) StoreUser(ctx context.Context, data *domain.User) (*domain.User, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	user, err := m.UserRepository.StoreUser(ctx, data)
	return user, err
}

func (m *userUserCase) UpdateUser(ctx context.Context) {

}

func (m *userUserCase) DeleteUser(ctx context.Context) {

}
