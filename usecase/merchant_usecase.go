package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
	"time"
)

type MerchantUseCase interface {
	Insert(merchant *model.Merchant) error
	Merchants() (*[]model.Merchant, error)
	MerchantById(id int) (*model.Merchant, error)
	Delete(id int) error
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

func (cu *merchantUseCase) Insert(merchant *model.Merchant) error {
	merchant.CreatedAt = time.Now()
	return cu.repo.Insert(merchant)
}

func (cu *merchantUseCase) Merchants() (*[]model.Merchant, error) {
	return cu.repo.Merchants()
}

func (cu *merchantUseCase) MerchantById(id int) (*model.Merchant, error) {
	return cu.repo.MerchantById(id)
}

func (cu *merchantUseCase) Delete(id int) error {
	return cu.repo.Delete(id)
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{
		repo: repo,
	}
}
