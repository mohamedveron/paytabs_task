package service

import "github.com/mohamedveron/paytabs_task/domains"

func (s *Service) GetAccounts() ([]domains.Account, error){

	var accounts []domains.Account

	for _, account := range s.AccountsDB {
		accounts = append(accounts, account)
	}

	return accounts, nil

}
