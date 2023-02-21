package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
)

type CustomerUseCase interface {
	Insert(store *model.Customer) error
	Customers() (*[]model.Customer, error)
	CustomerById(id int) (model.Customer, error)
	Update(store *model.Customer) error
	Delete(id int) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (cu *customerUseCase) Insert(store *model.Customer) error {
	return cu.repo.Insert(store)
}

func (cu *customerUseCase) Customers() (*[]model.Customer, error) {
	return cu.repo.Customers()
}

func (cu *customerUseCase) CustomerById(id int) (model.Customer, error) {
	return cu.repo.CustomerById(id)
}

func (cu *customerUseCase) Update(store *model.Customer) error {
	return cu.repo.Update(store)
}

func (cu *customerUseCase) Delete(id int) error {
	return cu.repo.Delete(id)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repo: repo,
	}
}
