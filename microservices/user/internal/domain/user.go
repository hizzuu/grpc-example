package domain

import "time"

type User struct {
	ID                int64
	Email             string
	Password          string
	EncryptedPassword string
	Name              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
