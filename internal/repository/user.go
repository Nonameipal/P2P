package repository

import (
	"os"

	dbModels "github.com/Nonameipal/P2P/internal/models/db"
	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/rs/zerolog"
)

func (r *Repository) CreateUser(user domain.User) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateUser").Logger()
	_, err = r.db.Exec(`INSERT INTO users (full_name, username, password, role, phone)
					VALUES ($1, $2, $3, $4, $5)`,
		user.FullName,
		user.Username,
		user.Password,
		user.Role,
		user.Phone,
	)
	if err != nil {
		logger.Err(err).Msg("error inserting user")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetUserByID(id string) (domain.User, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByID").Logger()
	var dbUser dbModels.User
	if err := r.db.Get(&dbUser, `
		SELECT id, full_name, username, password, role, created_at, updated_at , phone
		FROM users
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting user")
		return domain.User{}, r.translateError(err)
	}

	return dbUser.ToDomain(), nil
}

func (r *Repository) GetUserByUsername(username string) (domain.User, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByUsername").Logger()

	var dbUser dbModels.User
	if err := r.db.Get(&dbUser, `
		SELECT id, full_name, username, password, role, created_at, updated_at, phone
		FROM users
		WHERE username = $1`, username); err != nil {
		logger.Err(err).Msg("error selecting user")
		return domain.User{}, r.translateError(err)
	}

	return dbUser.ToDomain(), nil
}

func (r *Repository) GetAdminCount() (int, error) {
	var count int
	err := r.db.Get(&count, `SELECT COUNT(*) FROM users WHERE role = $1`, domain.AdminRole)
	return count, err
}

func (r *Repository) GetUserCount() (int, error) {
	var count int
	err := r.db.Get(&count, `SELECT COUNT(*) FROM users`)
	return count, err
}
