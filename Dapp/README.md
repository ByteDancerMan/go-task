## taks1查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。

```shell
go mod init
go get github.com/ethereum/go-ethereum/ethclient
go mod tidy
go run .\Dapp\task1\main.go

```
## task2发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。

```shell
go mod init
go get github.com/ethereum/go-ethereum/ethclient
go mod tidy
go run .\Dapp\task2\main.go

```

## task3查询交易
编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。

```shell
go mod init
go get github.com/ethereum/go-ethereum/ethclient

// 安装 abigen 工具
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

// 安装Solidity编译器的JavaScript版本
npm install -g solc

// 将智能合约生成abi
solcjs --abi .\Dapp\task3\add.sol

// 将智能合约生成bin
solcjs --bin .\Dapp\task3\add.sol


// 使用abigen生成Go绑定代码
abigen --abi .\Dapp\task3\add.abi --bin .\Dapp\task3\add.bin --pkg store --out .\Dapp\task3\store\add.go

go mod tidy
go run .\Dapp\task3\main.go
```

