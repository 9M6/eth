package main

import (
    "encoding/hex"
    "fmt"
    "log"
    "os"

    "github.com/kebani/eth/rlp"
)

func main() {
    message := []byte(os.Args[1])
    mLength := len(message)
    if mLength == 0 {
        panic("Please enter your argument")
    }

    encodedMsg := make([]byte, hex.DecodedLen(mLength))
    _, err := hex.Decode(encodedMsg, message)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(rlp.Decode(encodedMsg)))
}