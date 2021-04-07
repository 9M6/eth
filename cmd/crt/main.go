package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    _ "github.com/joho/godotenv/autoload"
)

func main() {

    // Notes:
    // https://medium.com/blockvigil/how-we-deal-with-chain-reorganization-at-ethvigil-5a8c06859c7
    //
    // https://goerli.etherscan.io/blocks_forked
    //
    // Use a Doubli Linked List to maintain a list of headers
    // and
    client, err := ethclient.Dial(os.Getenv("TESTNET"))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx, _ := context.WithCancel(context.Background())

    headers := make(chan *types.Header)
    _, err = client.SubscribeNewHead(ctx, headers)
    if err != nil {
        log.Fatal(err)
    }

    for ch := range headers {
        fmt.Println(ch.Number, ch.Root, ch.ParentHash, ch.Time, ch.Nonce)
    }

}