package app

import (
	"consensus-wallet/app/abi"
	"fmt"
	"log"
)

func Initialize() error {

	err := LoadConfig()
	if err != nil {
		log.Fatalf("can't parse config file")
	}

	err = abi.LoadABI()
	if err != nil {
		return fmt.Errorf("Abi loading error: %v", err)
	}
	
	return nil
}
