package domain

import "time"

type User struct {
	ID                int64
	Email             string
	Name              string
	Password          string
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
