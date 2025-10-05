package main

import (
    "context"
    "fmt"
    "log"

    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
    "github.com/gagliardetto/solana-go/programs/system"
)

func main() {
    // 创建 Solana RPC 客户端
    rpcClient := rpc.New("https://solana-devnet.g.alchemy.com/v2/QFdXM7qFQqQaB-A48af2NlEeC8bZhvLe")

    // 获取最新区块哈希 (使用 GetLatestBlockhash 替代 GetRecentBlockhash)
    resp, err := rpcClient.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
    if err != nil {
        log.Fatalf("Failed to get latest blockhash: %v", err)
    }
    blockhash := resp.Value.Blockhash
    fmt.Printf("Latest blockhash: %s\n", blockhash)

    // 查询账户余额
    pubKey := solana.MustPublicKeyFromBase58("BpVZJvGsYJGuxhiZZyRuuisxefpvHiXnWau8EnzSf5Fz")
    balanceResp, err := rpcClient.GetBalance(context.TODO(), pubKey, rpc.CommitmentFinalized)
    if err != nil {
        log.Fatalf("Failed to get balance: %v", err)
    }
    fmt.Printf("Balance: %d lamports\n", balanceResp.Value)

    // 构造原生转账交易
    from := solana.MustPublicKeyFromBase58("BpVZJvGsYJGuxhiZZyRuuisxefpvHiXnWau8EnzSf5Fz")
    to := solana.MustPublicKeyFromBase58("5YtuiQ9JEtLdgrNXrmSKrK2PLCpa74khydfRLj2qPqWe")
    lamports := uint64(10000000) // 转账金额（0.001 SOL）

    // 创建转账指令
    transferInst, err := system.NewTransferInstruction(
        lamports,
        from,
        to,
    ).ValidateAndBuild()
    if err != nil {
        log.Fatalf("Failed to build transfer instruction: %v", err)
    }

    // 构造交易
    tx, err := solana.NewTransaction(
        []solana.Instruction{transferInst},
        blockhash,
        solana.TransactionPayer(from),
    )
    if err != nil {
        log.Fatalf("Failed to create transaction: %v", err)
    }

    // 使用Solana私钥进行签名交易
    // 方式1: 从Base58编码的私钥字符串创建私钥
    privateKeyBase58 := "4uUHuFoviR9vx8tCGumBzKvq71WPjAwr97UuUujCGiUZxdRJiYg19pN5a3ABdZZy1jVkhZTmF42ETBuV5N7ZTNNG" // 替换为您的实际私钥
    accountFromPrivateKey, err := solana.PrivateKeyFromBase58(privateKeyBase58)
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }

    _, err = tx.Sign(
        func(key solana.PublicKey) *solana.PrivateKey {
            if key.Equals(from) {
                return &accountFromPrivateKey
            }
            return nil
        },
    )
    if err != nil {
        log.Fatalf("Failed to sign transaction: %v", err)
    }

    // 发送交易
    txResp, err := rpcClient.SendTransaction(context.TODO(), tx)
    if err != nil {
        log.Fatalf("Failed to send transaction: %v", err)
    }

    fmt.Printf("Transaction sent: %s\n", txResp)
}