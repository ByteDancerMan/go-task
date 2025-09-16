package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"go-task/Dapp/task3/store"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0x5Bc2C08bfbA16CDcD757Eb5CfE6C3Cd7F5CB95F2"
)


func main() {
	// 连接到以太坊Sepolia测试网络
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/QFdXM7qFQqQaB-A48af2NlEeC8bZhvLe")
	if err != nil {
		log.Fatal(err)
	}

	// 创建与已部署合约交互的实例
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	// 从私钥字符串加载私钥
	privateKey, err := crypto.HexToECDSA("5e4fec68dd3d33a2cba292356c72440e2ec1b81c3fda9bbe49993edf06aea4c9")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个带有链ID的交易签名者
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	// 调用合约的SetItem方法将键值对存储到合约中
	tx, err := storeContract.Add(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	// 创建调用选项用于查询合约状态
	callOpt := &bind.CallOpts{Context: context.Background()}

	// 从合约中读取刚才存储的值
	valueInContract, err := storeContract.A(callOpt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value in contract:", valueInContract)
}