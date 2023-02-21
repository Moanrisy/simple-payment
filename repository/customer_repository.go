package repository

import (
	"github.com/jmoiron/sqlx"
	"simple-payment/model"
	"simple-payment/util"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Customers() (*[]model.Customer, error)
	CustomerById(id int) (model.Customer, error)
	Update(customer *model.Customer) error
	Delete(id int) error
}

type customerRepository struct {
	db *sqlx.DB
}

func (cr *customerRepository) Insert(customer *model.Customer) error {
	createdCustomer := new(model.Customer)
	row := cr.db.QueryRowx(util.CREATE_CUSTOMER, customer.UserId, customer.Name, customer.CreatedAt)
	if err := row.StructScan(createdCustomer); err != nil {
		return err
	}

	return nil
}

func (cr *customerRepository) Customers() (*[]model.Customer, error) {
	customers := new([]model.Customer)

	if err := cr.db.Select(customers, util.ALL_CUSTOMER); err != nil {
		return nil, err
	}

	return customers, nil
}

func (cr *customerRepository) CustomerById(id int) (model.Customer, error) {
	return model.Customer{}, nil
}

func (cr *customerRepository) Update(customer *model.Customer) error {
	return nil
}

func (cr *customerRepository) Delete(id int) error {
	return nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{db: db}
}
