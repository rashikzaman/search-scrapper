package usecase

import (
	"context"
	"rashik/search-scrapper/app/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userUserCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(a domain.UserRepository) domain.UserUseCase {
	return &userUserCase{
		UserRepository: a,
	}
}

func (m *userUserCase) FetchUserById(ctx context.Context, id int) (*domain.User, error) {
	user, err := m.UserRepository.FetchUserById(ctx, id)
	return user, err
}

func (m *userUserCase) FetchUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := m.UserRepository.FetchUserByEmail(ctx, email)
	return user, err
}

func (m *userUserCase) StoreUser(ctx context.Context, data *domain.User) (*domain.User, error) {

	hasedPassword, err := m.HashPassword(*data.Password)
	if err != nil {
		return nil, err
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.Password = &hasedPassword
	user, err := m.UserRepository.StoreUser(ctx, data)
	return user, err
}

func (m *userUserCase) UpdateUser(ctx context.Context) {

}

func (m *userUserCase) DeleteUser(ctx context.Context) {

}

func (m *userUserCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *userUserCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
