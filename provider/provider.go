package provider

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

type provider struct {
	client *ethclient.Client
}

func NewProvider(url string) IProvider {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	return &provider{
		client: client,
	}
}

func (a *provider) BLockNumber() (uint64, error) {
	block, err := a.client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return block, nil
}
