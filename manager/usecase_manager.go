package manager

import "simple-payment/usecase"

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUseCase
	MerchantUseCase() usecase.MerchantUseCase
	BankUseCase() usecase.BankUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
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

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
