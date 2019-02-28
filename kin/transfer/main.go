package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kinecosystem/go/build"
	"github.com/kinecosystem/go/clients/horizon"
	"github.com/kinecosystem/go/keypair"
)

const (
	horizonURL = "https://horizon-testnet.kininfrastructure.com"
	// horizonURL = "http://172.18.16.176:80"
)

func fillAccounts(addresses [2]string) {
	for _, address := range addresses {
		friendBotResp, err := http.Get("https://friendbot-testnet.kininfrastructure.com/?addr=" + address)
		if err != nil {
			log.Fatal(err)
		}
		defer friendBotResp.Body.Close()
	}
}

var horizonClient *horizon.Client

func logBalances(addresses [2]string) {
	for _, address := range addresses {
		account, err := horizonClient.LoadAccount(address)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Balances for address:", address)
		for _, balance := range account.Balances {
			log.Println(balance)
		}
	}
}

func sendLumens(amount string, sourceSeed string, destinationAddress string, memo string) {
	tx, err := build.Transaction(
		build.SourceAccount{sourceSeed},
		build.Network{"Kin Testnet ; December 2018"},
		build.AutoSequence{SequenceProvider: horizonClient},
		build.MemoText{memo},
		build.Payment(
			build.Destination{AddressOrSeed: destinationAddress},
			build.NativeAmount{Amount: amount},
		),
	)

	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(sourceSeed)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	resp, err := horizonClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully sent", amount, "lumens to", destinationAddress, ". Hash:", resp.Hash)
}

func main() {
	fromPair := keypair.MustParse(opts.IssuerSeed)
	toPair := keypair.MustParse(opts.ToAddr)

	addresses := [2]string{fromPair.Address(), toPair.Address()}

	horizonClient = &horizon.Client{
		URL:  horizonURL,
		HTTP: &http.Client{Timeout: 10 * time.Second},
	}
	// fillAccounts(addresses)
	logBalances(addresses)
	sendLumens(opts.Amount, opts.IssuerSeed, toPair.Address(), opts.Memo)
	logBalances(addresses)
}
