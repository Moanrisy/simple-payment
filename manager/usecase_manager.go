package manager

import (
	"simple-payment/usecase"
	"simple-payment/util/authenthicator"
)

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUseCase
	MerchantUseCase() usecase.MerchantUseCase
	BankUseCase() usecase.BankUseCase
	PaymentUseCase() usecase.PaymentUseCase
	UserUseCase() usecase.UserUseCase
	HistoryUseCase() usecase.HistoryUseCase
}

type useCaseManager struct {
	repoManager  RepositoryManager
	tokenService authenthicator.AccessToken
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepository())
}

func (u *useCaseManager) MerchantUseCase() usecase.MerchantUseCase {
	return usecase.NewMerchantUseCase(u.repoManager.MerchantRepository())
}

func (u *useCaseManager) BankUseCase() usecase.BankUseCase {
	return usecase.NewBankUseCase(u.repoManager.BankRepository())
}

func (u *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repoManager.PaymentRepository())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepository(), u.tokenService)
}

func (u *useCaseManager) HistoryUseCase() usecase.HistoryUseCase {
	return usecase.NewHistoryUseCase(u.repoManager.HistoryRepository())
}

func NewUseCaseManager(repoManager RepositoryManager, tokenService authenthicator.AccessToken) UseCaseManager {
	return &useCaseManager{
		repoManager:  repoManager,
		tokenService: tokenService,
	}
}
