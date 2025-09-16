package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/QFdXM7qFQqQaB-A48af2NlEeC8bZhvLe")
	if err != nil {
		log.Fatal(err)
	}

	// 获取当前网络的连ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Chain ID:", chainID)

	// 查询某个区块的信息
	blockNumber := big.NewInt(9215591)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		log.Println("Transaction Hash:", tx.Hash().Hex())
		log.Println("To:", tx.To().Hex())
		log.Println("Value:", tx.Value().String())
		log.Println("Gas Limit:", tx.Gas())
		log.Println("Gas Price:", tx.GasPrice())
		log.Println("Nonce:", tx.Nonce())
		log.Println("Data:", tx.Data())

		// 从交易中恢复发送者地址
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err != nil {
			log.Println("Sender:", sender.Hex())
		} else {
			log.Println(err)
		}

		// 获取交易详情
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Transaction Receipt Status:", receipt.Status)
		fmt.Println("receipt.Logs:", receipt.Logs)
		break;
	}
}