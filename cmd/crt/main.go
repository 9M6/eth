package main

import (
    "context"
    "fmt"
    "log"
    "os"

    _ "github.com/joho/godotenv/autoload"

    "github.com/kebani/eth/crt"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    conn, err := crt.NewCRT(os.Getenv("TESTNET"))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connecting to Ethereum Network...")

    event := make(chan *crt.Event)
    if err = conn.Watch(ctx, event); err == nil {
        log.Fatal(err)
    }

    fmt.Println("Connected")
    fmt.Println("Looking up for Chain Reorganisations...")

    for e := range event {
        fmt.Println(e, e.Message)
        fmt.Println(conn)
    }
}