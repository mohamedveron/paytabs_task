package service

import (
	"github.com/mohamedveron/paytabs_task/domains"
	"sync"
)

type Service struct {
	AccountsDB	map[string]domains.Account
	mutex sync.Mutex
}

func NewService() *Service {

	AccountsDB := make(map[string]domains.Account)

	return &Service{
		AccountsDB: AccountsDB,
	}
}
