package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
)

type BankUseCase interface {
	Insert(bank *model.Bank) error
	Banks() (*[]model.Bank, error)
	BankById(id int) (*model.Bank, error)
	Delete(id int) error
}

type bankUseCase struct {
	repo repository.BankRepository
}

func (cu *bankUseCase) Insert(bank *model.Bank) error {
	return cu.repo.Insert(bank)
}

func (cu *bankUseCase) Banks() (*[]model.Bank, error) {
	return cu.repo.Banks()
}

func (cu *bankUseCase) BankById(id int) (*model.Bank, error) {
	return cu.repo.BankById(id)
}

func (cu *bankUseCase) Delete(id int) error {
	return cu.repo.Delete(id)
}

func NewBankUseCase(repo repository.BankRepository) BankUseCase {
	return &bankUseCase{
		repo: repo,
	}
}
