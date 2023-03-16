package client

import (
	"github.com/alexgao001/crosschain-tx-tool/tokenhub"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BSCClient struct {
	RpcClient *ethclient.Client
	ThClient  *tokenhub.Tokenhub
}

func NewBSCClient(url, contractAddress string) *BSCClient {
	rpcClient, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	thClient, err := tokenhub.NewTokenhub(
		common.HexToAddress(contractAddress),
		rpcClient)
	if err != nil {
		panic("new crossChain client error")
	}
	client := &BSCClient{
		RpcClient: rpcClient,
		ThClient:  thClient,
	}
	return client
}
