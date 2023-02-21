package repository

import (
	"github.com/jmoiron/sqlx"
	"simple-payment/model"
	"simple-payment/util"
)

type PaymentRepository interface {
	Insert(payment *model.Payment) error
	Payments() (*[]model.Payment, error)
	PaymentById(id int) (*model.Payment, error)
}

type paymentRepository struct {
	db *sqlx.DB
}

func (cr *paymentRepository) Insert(payment *model.Payment) error {
	createdPayment := new(model.Payment)
	row := cr.db.QueryRowx(util.CREATE_PAYMENT, payment.SenderId, payment.ReceiverId, payment.Amount, payment.BankAccountNumber, payment.CreatedAt)
	if err := row.StructScan(createdPayment); err != nil {
		return err
	}

	return nil
}

func (cr *paymentRepository) Payments() (*[]model.Payment, error) {
	payments := new([]model.Payment)

	if err := cr.db.Select(payments, util.ALL_PAYMENT); err != nil {
		return nil, err
	}

	return payments, nil
}

func (cr *paymentRepository) PaymentById(id int) (*model.Payment, error) {
	payment := new(model.Payment)

	if err := cr.db.Get(payment, util.READ_PAYMENT, id); err != nil {
		return nil, err
	}

	return payment, nil
}

func NewPaymentRepository(db *sqlx.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
