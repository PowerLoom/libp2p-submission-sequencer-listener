package prost

import (
	"Listen/config"
	"Listen/pkgs/contract"
	"context"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"time"
)

var Instance *contract.Contract

var (
	Client *ethclient.Client
)

func ConfigureClient() {
	var err error
	Client, err = ethclient.Dial(config.SettingsObj.ClientUrl)
	if err != nil {
		log.Fatal(err)
	}
}
func ConfigureContractInstance() {
	Instance, _ = contract.NewContract(common.HexToAddress(config.SettingsObj.ContractAddress), Client)
}

func MustQuery[K any](ctx context.Context, call func(opts *bind.CallOpts) (val K, err error)) (K, error) {
	expBackOff := backoff.NewExponentialBackOff()
	expBackOff.MaxElapsedTime = 1 * time.Minute
	var val K
	operation := func() error {
		var err error
		val, err = call(&bind.CallOpts{})
		return err
	}
	// Use the retry package to execute the operation with backoff
	err := backoff.Retry(operation, backoff.WithContext(expBackOff, ctx))
	if err != nil {
		return *new(K), err
	}
	return val, err
}
