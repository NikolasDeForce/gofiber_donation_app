package models

import (
	"time"
)

type Donate struct {
	ID             int       `db:"id" json:"id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	LoginWhoDonate string    `db:"loginwhodonate" json:"loginwhodonate" validate:"required,lte=255"`
	LoginToDonate  string    `db:"logintodonate" json:"logintodonate" validate:"required,lte=255"`
	Message        string    `db:"message" json:"message" validate:"required,lte=255"`
	Summary        int       `db:"summary" json:"summary"`
}
