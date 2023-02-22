package repository

import (
	"github.com/jmoiron/sqlx"
	"simple-payment/model"
	"simple-payment/util"
)

type HistoryRepository interface {
	Historys() (*[]model.LogHistory, error)
}

type historyRepository struct {
	db *sqlx.DB
}

func (cr *historyRepository) Historys() (*[]model.LogHistory, error) {
	historys := new([]model.LogHistory)

	if err := cr.db.Select(historys, util.ALL_HISTORIES); err != nil {
		return nil, err
	}

	return historys, nil
}

func NewHistoryRepository(db *sqlx.DB) HistoryRepository {
	return &historyRepository{db: db}
}
