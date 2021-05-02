package service

import (
	"errors"
	"fmt"
	"github.com/mohamedveron/paytabs_task/domains"
	"strconv"
	"sync"
)

func (s *Service) MakeTransfer(from , to string, balance float32) error {

	var fromAccount domains.Account
	var toAccount domains.Account

	if _, ok := s.AccountsDB[from]; ok {
		fromAccount = s.AccountsDB[from]
	}else{
		return errors.New("wrong from account")
	}

	if _, ok := s.AccountsDB[to]; ok {
		toAccount = s.AccountsDB[to]
	}else{
		return errors.New("wrong to account")
	}


	fromBalance, err := strconv.ParseFloat(fromAccount.Balance, 64)
	toBalance, err := strconv.ParseFloat(toAccount.Balance, 64)

	if err != nil {
		return err
	}

	if balance > float32(fromBalance) {
		return errors.New("don't have enough money")
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {

		defer wg.Done()

		toBalance+= float64(balance)
		newToBalance := fmt.Sprintf("%f", toBalance)

		fromBalance-= float64(balance)
		newFromBalance := fmt.Sprintf("%f", fromBalance)

		toAccount.Balance = newToBalance
		fromAccount.Balance = newFromBalance

		s.mutex.Lock()

		s.AccountsDB[to] = toAccount
		s.AccountsDB[from] = fromAccount

		s.mutex.Unlock()

	}()

	wg.Wait()

	return nil
}
