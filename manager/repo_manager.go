package manager

import "simple-payment/repository"

type RepositoryManager interface {
	CustomerRepository() repository.CustomerRepository
	MerchantRepository() repository.MerchantRepository
	BankRepository() repository.BankRepository
	PaymentRepository() repository.PaymentRepository
	UserRepository() repository.UserRepository
	HistoryRepository() repository.HistoryRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (rm *repositoryManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) MerchantRepository() repository.MerchantRepository {
	return repository.NewMerchantRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) BankRepository() repository.BankRepository {
	return repository.NewBankRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) PaymentRepository() repository.PaymentRepository {
	return repository.NewPaymentRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(rm.infra.SqlDB())
}

func (rm *repositoryManager) HistoryRepository() repository.HistoryRepository {
	return repository.NewHistoryRepository(rm.infra.SqlDB())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
