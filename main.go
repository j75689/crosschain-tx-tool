package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/alexgao001/crosschain-tx-tool/client"
	"github.com/alexgao001/crosschain-tx-tool/util"
	gnfdclient "github.com/bnb-chain/greenfield-go-sdk/client/chain"
	"github.com/bnb-chain/greenfield-go-sdk/keys"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bridgetypes "github.com/bnb-chain/greenfield/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagFromChain            = "from-chain"
	flagKey                  = "key"
	flagChainId              = "chain-id"
	flagUrl                  = "url"
	flagToAddress            = "to"
	flagAmount               = "amount"
	flagTokenHubContractAddr = "contract-addr"
	flagNumberOfTx           = "tx-count"

	bsc  = "bsc"
	gnfd = "gnfd"
)

func initFlags() {
	flag.String(flagFromChain, "", "from chain")
	flag.String(flagKey, "", "private key")
	flag.String(flagChainId, "", "chain id")
	flag.String(flagUrl, "", "url")
	flag.String(flagToAddress, "", "to")
	flag.String(flagAmount, "0", "amount")
	flag.String(flagTokenHubContractAddr, "", "smart contract address")
	flag.Int64(flagNumberOfTx, 0, "number of tx in batch")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func main() {
	initFlags()
	fromChain := viper.GetString(flagFromChain)
	key := viper.GetString(flagKey)
	chainId := viper.GetString(flagChainId)
	url := viper.GetString(flagUrl)
	to := viper.GetString(flagToAddress)
	amountStr := viper.GetString(flagAmount)
	amount, _ := new(big.Int).SetString(amountStr, 10)
	txCount := viper.GetInt64(flagNumberOfTx)

	if fromChain == gnfd {
		err := gnfdCrossChainTx(key, chainId, to, url, amount, txCount)
		if err != nil {
			panic(err)
		}
	} else if fromChain == bsc {
		cAddr := viper.GetString(flagTokenHubContractAddr)
		err := bscCrossChainTx(key, chainId, to, url, cAddr, amount, txCount)
		if err != nil {
			panic(err)
		}
	} else {
		panic("chain not correct")
	}
}

func gnfdCrossChainTx(key, chainId, to, url string, amount *big.Int, txCount int64) error {
	km, err := keys.NewPrivateKeyManager(key)
	if err != nil {
		return err
	}
	gnfdClient := gnfdclient.NewGreenfieldClient(url, chainId, gnfdclient.WithKeyManager(km),
		gnfdclient.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))

	msgTransferOut := bridgetypes.NewMsgTransferOut(km.GetAddr().String(), to, &sdk.Coin{
		Denom:  "BNB",
		Amount: sdk.NewIntFromBigInt(amount),
	})

	nonce, err := gnfdClient.GetNonce()
	if err != nil {
		return err
	}
	for i := 0; i < int(txCount); i++ {
		txOpt := &types.TxOption{
			Nonce: nonce,
		}
		response, err := gnfdClient.BroadcastTx([]sdk.Msg{msgTransferOut}, txOpt)
		if err != nil {
			return err
		}
		nonce++
		fmt.Println(response.TxResponse.String())
	}
	return nil
}

func bscCrossChainTx(key, chainId, to, url, cAddr string, amount *big.Int, txCount int64) error {
	chainIdInt, err := strconv.Atoi(chainId)
	if err != nil {
		return err
	}

	privKey := util.GetBscPrivateKey(key)
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("get public key error")
	}
	txSender := crypto.PubkeyToAddress(*publicKeyECDSA)
	bscClient := client.NewBSCClient(url, cAddr)
	nonce, err := bscClient.RpcClient.PendingNonceAt(context.Background(), txSender)
	if err != nil {
		return err
	}
	for i := 0; i < int(txCount); i++ {
		txOpts, err := getTransactor(privKey, nonce, big.NewInt(int64(chainIdInt)), amount)
		if err != nil {
			return err
		}
		fmt.Printf("tranfer amount %d bnb to addres %s\n", amount, to)
		res, err := bscClient.ThClient.TransferOut(
			txOpts,
			common.HexToAddress(to),
			amount,
		)
		if err != nil {
			return err
		}
		nonce++
		fmt.Printf("tx hash %s\n", res.Hash().String())
	}
	return nil
}

func getTransactor(privateKey *ecdsa.PrivateKey, nonce uint64, chainID *big.Int, amount *big.Int) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	relayFee := 4 * math.Pow(10, 15)
	f := int64(relayFee)
	txOpts.Value = new(big.Int).Add(amount, big.NewInt(f))
	txOpts.GasLimit = 4700000
	txOpts.GasPrice = big.NewInt(20000000000)
	return txOpts, nil
}
