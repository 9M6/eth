package crt

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

type CRT struct {
    rpc *ethclient.Client
    top *big.Int
    log map[*big.Int]*types.Header
    chn chan *types.Header
}

func NewCRT(host string) (*CRT, error) {
    rpc, err := ethclient.Dial(host)
    if err != nil {
        return nil, err
    }
    return &CRT{
        rpc: rpc,
        log: make(map[*big.Int]*types.Header),
        top: big.NewInt(0),
        chn:  make(chan *types.Header),
    }, nil
}

func (c CRT) Watch(ctx context.Context, event chan<- string) error {
    conn, err := c.rpc.SubscribeNewHead(ctx, c.chn)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Unsubscribe()

    for ch := range c.chn {
        if ch.Number.Int64() <= c.top.Int64() {

        }
        fmt.Println(ch.Number, ch.Root, ch.ParentHash, ch.Time, ch.Nonce)
    }
}