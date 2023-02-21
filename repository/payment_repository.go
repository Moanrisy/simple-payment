package repository

import (
	"errors"
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

	tx := cr.db.MustBegin()

	// reduce customer balance and increase bank balance
	customer := new(model.Customer)
	if err := tx.Get(customer, util.READ_CUSTOMER, payment.SenderId); err != nil {
		err = tx.Rollback()
		err = errors.New("customer not found")
		return err
	}
	customer.Balance = customer.Balance - payment.Amount
	if customer.Balance < 0 {
		err := tx.Rollback()
		err = errors.New("insufficient customer balance")
		return err
	}
	_, err := tx.Exec(util.UPDATE_CUSTOMER_BALANCE, customer.Balance, payment.SenderId)
	if err != nil {
		err := tx.Rollback()
		err = errors.New("failed reducing customer balance")
		return err
	}

	bank := new(model.Bank)
	if err := tx.Get(bank, util.READ_BANK_BY_ACCOUNT_NUMBER, payment.BankAccountNumber); err != nil {
		err = tx.Rollback()
		err = errors.New("bank not found")
		return err
	}
	bank.Balance = bank.Balance + payment.Amount
	_, err = tx.Exec(util.UPDATE_BANK_BALANCE, bank.Balance, payment.BankAccountNumber)
	if err != nil {
		err := tx.Rollback()
		err = errors.New("failed increasing bank balance")
		return err
	}

	// increase merchant balance and reduce bank balance
	merchant := new(model.Merchant)
	if err := tx.Get(merchant, util.READ_MERCHANT, payment.ReceiverId); err != nil {
		err = tx.Rollback()
		err = errors.New("merchant not found")
		return err
	}
	merchant.Balance = merchant.Balance + payment.Amount
	_, err = tx.Exec(util.UPDATE_MERCHANT_BALANCE, merchant.Balance, payment.ReceiverId)
	if err != nil {
		err := tx.Rollback()
		err = errors.New("failed increasing merchant balance")
		return err
	}

	if err := tx.Get(bank, util.READ_BANK_BY_ACCOUNT_NUMBER, payment.BankAccountNumber); err != nil {
		err = tx.Rollback()
		err = errors.New("bank not found")
		return err
	}
	bank.Balance = bank.Balance - payment.Amount
	_, err = tx.Exec(util.UPDATE_BANK_BALANCE, bank.Balance, payment.BankAccountNumber)
	if err != nil {
		err := tx.Rollback()
		err = errors.New("failed reducing bank balance")
		return err
	}

	// insert payment if all previous steps are successful
	createdPayment := new(model.Payment)
	row := cr.db.QueryRowx(util.CREATE_PAYMENT, payment.SenderId, payment.ReceiverId, payment.Amount, payment.BankAccountNumber, payment.CreatedAt)
	if err := row.StructScan(createdPayment); err != nil {
		err = tx.Rollback()
		err = errors.New("failed inserting payment")
		return err
	}

	err = tx.Commit()

	return err
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
