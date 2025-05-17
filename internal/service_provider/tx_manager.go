package service_provider

import "github.com/GP-Hacks/auth/internal/utils/transaction"

func (s *ServiceProvider) TxManager() transaction.TransactionManager {
	return transaction.NewTransactionManager(s.DB())
}
