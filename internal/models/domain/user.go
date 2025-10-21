package domain

import "time"

type User struct {
	ID           int
	FullName     string
	Username     string
	Password     string
	Email        string
	Phone        string
	Role         string
	IsIdentified bool
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

var (
	UserRole  = "USER"
	AdminRole = "ADMIN"
)
