package repository

import (
	"github.com/jmoiron/sqlx"
	"simple-payment/model"
	"simple-payment/util"
)

type MerchantRepository interface {
	Insert(merchant *model.Merchant) error
	Merchants() (*[]model.Merchant, error)
	MerchantById(id int) (*model.Merchant, error)
	Delete(id int) error
}

type merchantRepository struct {
	db *sqlx.DB
}

func (cr *merchantRepository) Insert(merchant *model.Merchant) error {
	createdMerchant := new(model.Merchant)
	row := cr.db.QueryRowx(util.CREATE_MERCHANT, merchant.UserId, merchant.Name, merchant.CreatedAt)
	if err := row.StructScan(createdMerchant); err != nil {
		return err
	}

	return nil
}

func (cr *merchantRepository) Merchants() (*[]model.Merchant, error) {
	merchants := new([]model.Merchant)

	if err := cr.db.Select(merchants, util.ALL_MERCHANT); err != nil {
		return nil, err
	}

	return merchants, nil
}

func (cr *merchantRepository) MerchantById(id int) (*model.Merchant, error) {
	merchant := new(model.Merchant)

	if err := cr.db.Get(merchant, util.READ_MERCHANT, id); err != nil {
		return nil, err
	}

	return merchant, nil
}

func (cr *merchantRepository) Delete(id int) error {
	if _, err := cr.db.Exec(util.DELETE_MERCHANT, id); err != nil {
		return err
	}

	return nil
}

func NewMerchantRepository(db *sqlx.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
