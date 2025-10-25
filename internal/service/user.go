package service

import (
	"errors"

	"github.com/Nonameipal/P2P/internal/errs"
	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/Nonameipal/P2P/utils"
)

func (s *Service) CreateUser(user domain.User) (err error) {
	// Проверить существует ли пользователь с таким username'ом в бд
	_, err = s.repository.GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUsernameAlreadyExists
	}

	// За хэшировать пароль
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	user.Role = domain.UserRole

	// Добавить пользователя в бд
	if err = s.repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}
func (s *Service) Authenticate(user domain.User) (int, string, error) {
	// проверить существует ли пользователь с таким username
	userFromDB, err := s.repository.GetUserByUsername(user.Username)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return 0, "", errs.ErrUserNotFound
		}
		return 0, "", err
	}

	// пароль от пользователя не нужно хешировать,
	// т.к. мы будем сравнивать его с хешем из базы

	// проверить правильно ли он указал пароль
	if err := utils.CompareHash(userFromDB.Password, user.Password); err != nil {
		return 0, "", errs.ErrIncorrectUsernameOrPassword
	}

	return userFromDB.ID, userFromDB.Role, nil
}
