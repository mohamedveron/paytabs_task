package tests

import (
	"github.com/mohamedveron/paytabs_task/service"
	"testing"
)

func TestMakeRightTransfer(t *testing.T){

	serviceLayer := service.NewService()
	serviceLayer.ConsumeAccount()

	err := serviceLayer.MakeTransfer("793b565c-1ec0-4f21-8edd-fd7ee3011659", "76dc4db2-cb0a-490a-a6d4-05c850195916", 358)

	if err != nil {
		t.Errorf("Transfer failed")
	}
}

func TestMakeNegativeTransfer(t *testing.T){

	serviceLayer := service.NewService()
	serviceLayer.ConsumeAccount()

	err := serviceLayer.MakeTransfer("793b565c-1ec0-4f21-8edd-fd7ee3011659", "76dc4db2-cb0a-490a-a6d4-05c850195916", 3000)

	if err == nil {
		t.Errorf("Transfer which exceded the amount must not be affected")
	}
}
