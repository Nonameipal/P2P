package db

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type User struct {
	ID        int        `db:"id"`
	FullName  string     `db:"full_name"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	Phone     string     `db:"phone"`
	Role      string     `db:"role"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:        u.ID,
		FullName:  u.FullName,
		Username:  u.Username,
		Password:  u.Password,
		Phone:     u.Phone,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) FromDomain(du domain.User) {
	u.ID = du.ID
	u.FullName = du.FullName
	u.Username = du.Username
	u.Password = du.Password
	u.Phone = du.Phone
	u.Role = du.Role
	u.CreatedAt = du.CreatedAt
	u.UpdatedAt = du.UpdatedAt
}
