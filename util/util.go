package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetBscPrivateKey(privateKey string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return privKey
}
