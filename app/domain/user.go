package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     *string   `json:"email" gorm:"unique;type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	FetchUserById(ctx context.Context)
	FetchUserByEmail(ctx context.Context)
	StoreUser(context.Context, *User) (*User, error)
}

type UserUseCase interface {
}
