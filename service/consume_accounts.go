package service

import (
	"encoding/json"
	"fmt"
	"github.com/mohamedveron/paytabs_task/domains"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (s *Service) ConsumeAccount() error{

	response, err := http.Get("https://git.io/Jm76h")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var accounts []domains.Account
	err = json.Unmarshal(responseData, &accounts)

	if err != nil {
		log.Fatal(err)
	}

	accountChannel := make(chan domains.Account)

	// consume the accounts data by 3 concurrent goroutines
	for i := 0; i < 3; i++{
		go s.accountIngestor(accountChannel)
	}


	for idx := range accounts {
		accountChannel <- accounts[idx]
	}

	return err
}

func (s *Service) accountIngestor(accounts chan domains.Account){

	for account := range accounts {
		s.mutex.Lock()
		s.AccountsDB[account.ID] = account
		s.mutex.Unlock()
	}
}