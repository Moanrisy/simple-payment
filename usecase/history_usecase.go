package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
)

type HistoryUseCase interface {
	Historys() (*[]model.LogHistory, error)
}

type historyUseCase struct {
	repo repository.HistoryRepository
}

func (cu *historyUseCase) Historys() (*[]model.LogHistory, error) {
	return cu.repo.Historys()
}

func NewHistoryUseCase(repo repository.HistoryRepository) HistoryUseCase {
	return &historyUseCase{
		repo: repo,
	}
}
