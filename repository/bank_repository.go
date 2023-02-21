package repository

import (
	"github.com/jmoiron/sqlx"
	"simple-payment/model"
	"simple-payment/util"
)

type BankRepository interface {
	Insert(bank *model.Bank) error
	Banks() (*[]model.Bank, error)
	BankById(id int) (*model.Bank, error)
	Delete(id int) error
}

type bankRepository struct {
	db *sqlx.DB
}

func (cr *bankRepository) Insert(bank *model.Bank) error {
	createdBank := new(model.Bank)
	row := cr.db.QueryRowx(util.CREATE_BANK, bank.BankAccountNumber, bank.Name)
	if err := row.StructScan(createdBank); err != nil {
		return err
	}

	return nil
}

func (cr *bankRepository) Banks() (*[]model.Bank, error) {
	banks := new([]model.Bank)

	if err := cr.db.Select(banks, util.ALL_BANK); err != nil {
		return nil, err
	}

	return banks, nil
}

func (cr *bankRepository) BankById(id int) (*model.Bank, error) {
	bank := new(model.Bank)

	if err := cr.db.Get(bank, util.READ_BANK, id); err != nil {
		return nil, err
	}

	return bank, nil
}

func (cr *bankRepository) Delete(id int) error {
	if _, err := cr.db.Exec(util.DELETE_BANK, id); err != nil {
		return err
	}

	return nil
}

func NewBankRepository(db *sqlx.DB) BankRepository {
	return &bankRepository{db: db}
}
