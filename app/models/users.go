package models

import (
	"time"
)

type User struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Login     string    `db:"login" json:"login" validate:"required,lte=255"`
	Email     string    `db:"email" json:"email" validate:"required,lte=255"`
	Password  string    `db:"password" json:"password" validate:"required,lte=255"`
}
