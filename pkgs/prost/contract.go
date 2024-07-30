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
	"math/big"
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

func UpdateSubmissionLimit(curBlock *big.Int) *big.Int {
	var submissionLimit *big.Int
	if window, err := Instance.SnapshotSubmissionWindow(&bind.CallOpts{}); err != nil {
		log.Errorf("Failed to fetch snapshot submission window: %s\n", err.Error())
	} else {
		submissionLimit = new(big.Int).Add(curBlock, window)
		submissionLimit = submissionLimit.Add(submissionLimit, big.NewInt(1))
		log.Debugln("Snapshot Submission Limit:", submissionLimit)
	}
	return submissionLimit
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
