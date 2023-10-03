package models

import (
	"time"
)

type Account struct {
	ID          int
	Username    string
	DisplayName string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
