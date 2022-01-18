package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     *string   `json:"email" gorm:"unique;type:varchar(255);not null"`
	Password  *string   `json:"-" gorm:"type:varchar(255);not null"` //prevent password appearing in json
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Keyword   []Keyword
}

type UserRepository interface {
	FetchUserById(ctx context.Context, id int) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	StoreUser(context.Context, *User) (*User, error)
}

type UserUseCase interface {
	FetchUserById(ctx context.Context, id int) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	StoreUser(context.Context, *User) (*User, error)
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}
