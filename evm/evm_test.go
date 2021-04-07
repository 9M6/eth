package evm

import (
    "encoding/hex"
    "fmt"
    "testing"

    "github.com/ethereum/go-ethereum/crypto"
)

func TestNewEVM(t *testing.T) {}

func TestEVM_Execute(t *testing.T) {
    mapOfByteCode := map[string]string{
        // "#1": "60016020526002606452600361ff0052600362ffffff526005601053",
        // "#2": "7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00016000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00026020527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff05604052",
        "#3": "7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000a6000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000a6020527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0a604052",
    }

    for _, value := range mapOfByteCode {
        bytecode, err := hex.DecodeString(value)
        if err != nil {
            t.Fatalf("something went wrong converting the bytecode from hex to byte")
        }

        evm := NewEVM(bytecode)
        bytes, gas := evm.Execute()

        fmt.Println(crypto.Keccak256Hash(bytes), gas)
    }
}

func TestEVM_GasConsumption(t *testing.T) {}

func TestEVM_StackDoublePop(t *testing.T) {}

func TestEVM_MergePushN(t *testing.T) {}

func TestEVM_IndexPushN(t *testing.T) {}