package usecase

import (
	"errors"
	"simple-payment/model"
	"simple-payment/repository"
	"simple-payment/util"
	"simple-payment/util/authenthicator"
)

type UserUseCase interface {
	Insert(user *model.User) error
	Login(user model.UserCredential) (token string, err error)
}

type userUseCase struct {
	repo         repository.UserRepository
	tokenService authenthicator.AccessToken
}

func (cu *userUseCase) Insert(user *model.User) error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return cu.repo.Insert(user)
}

func (cu *userUseCase) Login(userCredential model.UserCredential) (token string, err error) {
	user, err := cu.repo.UserLogin(&userCredential)

	if err != nil && err.Error() == "invalid password" {
		return "", errors.New("invalid password")
	}

	if user.UserId == "" {
		return "", errors.New("user does not exist, please sign up first")
	}

	userCredential.UserId = user.UserId
	if err == nil {
		token, err := cu.tokenService.CreateAccessToken(&userCredential)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", err
	}
}

func NewUserUseCase(repo repository.UserRepository, tokenService authenthicator.AccessToken) UserUseCase {
	return &userUseCase{
		repo:         repo,
		tokenService: tokenService,
	}
}
