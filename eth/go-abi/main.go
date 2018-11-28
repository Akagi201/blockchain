package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func transfer() {
	//首先导入上面生成的账户密钥（json）和密码
	auth, err := bind.NewTransactor(strings.NewReader("json"), "password")
	//这句用的是生成的token.go里面的方法
	//client变量是我们第一步连接以太坊rpc节点的时候创建的
	//contractAddress 是代币地址，比如eos 的地址是0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0
	//那么我们转账针对的就是账户里的eos代币
	//具体看这里 https://etherscan.io/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0
	token, err := NewToken(common.HexToAddress("0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0"), client)
	if err != nil {
		panic(err)
	}
	//每个代币都会有相应的位数，例如eos是18位，那么我们转账的时候，需要在金额后面加18个0
	decimal, err := token.Decimals(nil)
	if err != nil {
		panic(err)
	}
	//这是处理位数的代码段
	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})
	//然后就可以转账到你需要接受的账户上了
	//toAddress 是接受eos的账户地址
	txs, err := token.Transfer(auth, common.HexToAddress(toAddress), convertAmount)
}

func main() {
	rpcDial, err := rpc.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(rpcDial)
	fmt.Printf("eth client: %#v\n", client)

	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)

	address, _ := ks.NewAccount("password")
	account, err := ks.Export(address, "password", "password")
	if err != nil {
		panic(err)
	}

	fmt.Printf("account: %v\n", string(account))

	fmt.Printf("account addr: %v\n", address.Address.Hex())
}
