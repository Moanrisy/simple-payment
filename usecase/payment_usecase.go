package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
	"time"
)

type PaymentUseCase interface {
	Insert(payment *model.Payment) error
	Payments() (*[]model.Payment, error)
	PaymentById(id int) (*model.Payment, error)
}

type paymentUseCase struct {
	repo repository.PaymentRepository
}

func (cu *paymentUseCase) Insert(payment *model.Payment) error {
	payment.CreatedAt = time.Now()
	return cu.repo.Insert(payment)
}

func (cu *paymentUseCase) Payments() (*[]model.Payment, error) {
	return cu.repo.Payments()
}

func (cu *paymentUseCase) PaymentById(id int) (*model.Payment, error) {
	return cu.repo.PaymentById(id)
}

func NewPaymentUseCase(repo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCase{
		repo: repo,
	}
}
