package main

//使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。
//  具体任务
//  环境搭建
//  安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
//  注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
//  查询区块
//  编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
//  实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
//  输出查询结果到控制台。
//  发送交易
//  准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
//  编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
//  构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
//  对交易进行签名，并将签名后的交易发送到网络。
//  输出交易的哈希值。

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() { 
	// 连接ETH节点
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/QFdXM7qFQqQaB-A48af2NlEeC8bZhvLe")
	if err != nil {
        log.Fatal(err)
	}

	// 从十六进制私钥字符串加载私钥
	privateKey, err := crypto.HexToECDSA("5e4fec68dd3d33a2cba292356c72440e2ec1b81c3fda9bbe49993edf06aea4c9")
	if err != nil {
		log.Fatal(err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取发送方地址并查询待处理的nonce值
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易参数：转账金额、Gas 价格和 Gas 限制
	value := big.NewInt(1000000000000000) // 0.001 ETH
	gasLimit := uint64(21000)		   // 标准以太币转账的Gas限制
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 创建交易并签名
	toAddress := common.HexToAddress("0xDc35F50D6f7654FD3Ed096EE093789Dbe2875c61")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 获取链ID用于EIP-155签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tx sent: ", signedTx.Hash().Hex())
}