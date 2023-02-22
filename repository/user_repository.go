package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"simple-payment/model"
	"simple-payment/util"
)

type UserRepository interface {
	Insert(user *model.User) error
	UserLogin(userCredential *model.UserCredential) (model.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func (ur *userRepository) Insert(user *model.User) error {
	createdUser := new(model.User)
	row := ur.db.QueryRowx(util.CREATE_USER, user.Email, user.Password)
	if err := row.StructScan(createdUser); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) UserLogin(userCredential *model.UserCredential) (model.User, error) {
	var user model.User
	row := ur.db.QueryRowx(util.LOGIN_USER, userCredential.Email)
	err := row.Scan(&user.UserId, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}
	err = util.VerifyPassword(user.Password, userCredential.Password)
	if err != nil {
		return model.User{}, errors.New("invalid password")
	}

	return user, nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
