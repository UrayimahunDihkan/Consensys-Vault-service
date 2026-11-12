package abi

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	pledgePoolABIPath = "app/abi/ProposalSign.abi.json"
)

var (
	proposalSign abi.ABI
	once         sync.Once
)

func LoadABI() error {
	var err error
	once.Do(func() {
		proposalSign, err = loadABI(pledgePoolABIPath)
		if err != nil {
			return
		}
	})
	return err
}

func loadABI(path string) (abi.ABI, error) {
	abiBytes, err := os.ReadFile(path)
	if err != nil {
		return abi.ABI{}, err
	}

	var artifact struct {
		ABI json.RawMessage `json:"abi"`
	}
	if json.Unmarshal(abiBytes, &artifact) == nil && len(artifact.ABI) > 0 {
		abiBytes = artifact.ABI
	}

	return abi.JSON(strings.NewReader(string(abiBytes)))
}

// PledgePoolABI 返回池子合约的 ABI
func PledgePoolABI() abi.ABI {
	return proposalSign
}
