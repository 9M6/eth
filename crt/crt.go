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
    conn   *ethclient.Client
    ch     chan *types.Header
    height *big.Int
    chain  map[*big.Int]*types.Header
}

type Event struct {
    Height  *big.Int
    Message string
    Length  *big.Int
}

func NewCRT(host string) (*CRT, error) {
    rpc, err := ethclient.Dial(host)
    if err != nil {
        return nil, err
    }

    return &CRT{
        conn:   rpc,
        ch:     make(chan *types.Header),
        height: big.NewInt(0),
        chain:  make(map[*big.Int]*types.Header),
    }, nil
}

func (c *CRT) Watch(ctx context.Context, event chan<- *Event) (err error) {
    conn, err := c.conn.SubscribeNewHead(ctx, c.ch)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Unsubscribe()

    fmt.Println("Connected...")
    fmt.Println("Looking up for Chain Reorganisations...")

    for ch := range c.ch {
        switch {
        case ch.Number.Uint64() == c.height.Uint64():
            event <- &Event{
                Message: "chain reorganisation event",
                Height:  ch.Number,
                Length:  big.NewInt(0),
            }
        case ch.Number.Uint64() <= c.height.Uint64():
            length := big.NewInt(0).Sub(c.height, ch.Number)
            event <- &Event{
                Height:  ch.Number,
                Length:  length,
                Message: "chain reorganisation event",
            }
            for i := uint64(0); i < length.Uint64(); i++ {
                c.chain[ch.Number] = ch
                c.height = ch.Number
            }
        }
        c.chain[ch.Number] = ch
        c.height = ch.Number
    }

    return
}