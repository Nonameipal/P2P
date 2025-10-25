package domain

import "time"

type User struct {
	ID        int
	FullName  string
	Username  string
	Password  string
	Phone     string
	Role      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

var (
	UserRole  = "USER"
	AdminRole = "ADMIN"
)
