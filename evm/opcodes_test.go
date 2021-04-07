package evm

import (
    "encoding/hex"
    "fmt"
    "testing"

    "github.com/holiman/uint256"
)

const (
    FFFF = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
    FF00 = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00"
)

func TestAdd(t *testing.T) {
    bFFFF, _ := hex.DecodeString(FFFF)
    bFF00, _ := hex.DecodeString(FF00)

    x := uint256.NewInt().SetBytes(bFFFF)
    y := uint256.NewInt().SetBytes(bFF00)

    z := Add(x, y)

    fmt.Println(z, z.AddOverflow(x, y))
}

func TestMul(t *testing.T) {
    bFFFF, _ := hex.DecodeString(FFFF)
    bFF00, _ := hex.DecodeString(FF00)

    x := uint256.NewInt().SetBytes(bFFFF)
    y := uint256.NewInt().SetBytes(bFF00)

    z := Mul(x, y)

    fmt.Println(z, z.MulOverflow(x, y))
}

func TestSDiv(t *testing.T) {
    bFFFF, _ := hex.DecodeString(FFFF)
    bFF00, _ := hex.DecodeString(FF00)

    x := uint256.NewInt().SetBytes(bFFFF)
    y := uint256.NewInt().SetBytes(bFF00)

    z := SDiv(x, y)

    fmt.Println(z)
}