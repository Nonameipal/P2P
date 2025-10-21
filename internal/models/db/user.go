package db

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type User struct {
	ID           int        `db:"id"`
	FullName     string     `db:"full_name"`
	Username     string     `db:"username"`
	Password     string     `db:"password"`
	Role         string     `db:"role"`
	IsIdentified bool       `db:"is_identified"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:           u.ID,
		FullName:     u.FullName,
		Username:     u.Username,
		Password:     u.Password,
		Role:         u.Role,
		IsIdentified: u.IsIdentified,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func (u *User) FromDomain(du domain.User) {
	u.ID = du.ID
	u.FullName = du.FullName
	u.Username = du.Username
	u.Password = du.Password
	u.Role = du.Role
	u.IsIdentified = du.IsIdentified
	u.CreatedAt = du.CreatedAt
	u.UpdatedAt = du.UpdatedAt
}
