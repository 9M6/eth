package evm

import (
    "github.com/holiman/uint256"
)

const (
    ADD     = 0x01
    MUL     = 0x02
    SDIV    = 0x05
    EXP     = 0x0A
    MSTORE  = 0x52
    MSTORE8 = 0x53
    PUSH1   = 0x60
    PUSH2   = 0x61
    PUSH3   = 0x62
    PUSH32  = 0x7F
)

func Add(x, y *uint256.Int) *uint256.Int {
    return uint256.NewInt().Add(x, y)
}

func Mul(x, y *uint256.Int) *uint256.Int {
    return uint256.NewInt().Mul(x, y)
}

func SDiv(x, y *uint256.Int) *uint256.Int {
    return uint256.NewInt().SDiv(x, y)
}

func Exp(x, y *uint256.Int) *uint256.Int {
    return uint256.NewInt().Exp(x, y)
}