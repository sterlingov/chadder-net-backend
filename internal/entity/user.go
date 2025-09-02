package entity

import "time"

type User struct {
	ID           int64
	Username     string
	Name         string
	Email        string
	Bio          string
	PasswordHash string
	CreatedAt    time.Time
	Avatar       string
}
