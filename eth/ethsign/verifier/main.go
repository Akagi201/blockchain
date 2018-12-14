package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

// SignHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
//
// https://github.com/ethereum/go-ethereum/blob/master/internal/ethapi/api.go#L412
func SignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func main() {
	privateKey, err := crypto.HexToECDSA("107be946709e41b7895eea9f2dacf998a0a9124acbb786f0fd1a826101581a07")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	publicAddrStr := hex.EncodeToString(publicAddr[:])

	log.Infof("public key: %v", publicAddrStr)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("foobar")
	{
		hash := SignHash(data)
		log.Infof("eth hash: %v", common.ToHex(hash))
	}

	hash := crypto.Keccak256Hash(data)
	log.Infof("hash: %v", hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("signature: %v", hexutil.Encode(signature))

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	log.Infof("public key matches 1: %v", matches)

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	log.Infof("public key matches 2: %v", matches)

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)

	log.Infof("signature verified: %v", verified)
}
