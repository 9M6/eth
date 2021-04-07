package evm

import (
    "github.com/holiman/uint256"
)

// EVM is a struct type that executes a series of
// instructions on a stack based approach.
//
// The VM supports a limited series of OpCodes, for
// more about the OpCode supported go to `opcodes.go`
type EVM struct {
    Code  []byte
    Index int
    Mem   *Mem
    Stack Stack
    Gas   uint64
}

// NewEVM Initialises the Ethereum Virtual Machine
// and takes a argument of type string, where the argument
// is bytecode encoded smart contract.
func NewEVM(byteCode []byte) *EVM {
    return &EVM{
        Code:    byteCode,
        Index:   0,
        Mem:     NewMem(),
        Stack:   Stack{},
        Gas:     0,
    }
}

// Execute runs the assigned code at initialisation
// and follows through the OpCodes through an Index
// that holds the location next available OpCode to
// be Executed.
//
// The function returns one item from the stack
// which should be the last and finished computed value
func (e *EVM) Execute() ([]byte, uint64) {
    for e.Index < len(e.Code) {
        switch OpCode := e.Code[e.Index]; OpCode {
        case ADD:
            x, y := e.StackDoublePop()
            e.Stack = e.Stack.Push(Add(x, y))
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, nil)
        case MUL:
            x, y := e.StackDoublePop()
            e.Stack = e.Stack.Push(Mul(x, y))
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, nil)
        case SDIV:
            x, y := e.StackDoublePop()
            e.Stack = e.Stack.Push(SDiv(x, y))
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, nil)
        case EXP:
            x, y := e.StackDoublePop()
            e.Stack = e.Stack.Push(Exp(x, y))
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, func() uint64 {
                return uint64(y.ByteLen())
            })
        case MSTORE:
            offset, value := e.StackDoublePop()
            e.Mem.Set32(offset.Uint64(), value)
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, func() uint64 {
                return 0
            })
        case MSTORE8:
            offset, value := e.StackDoublePop()
            e.Mem.Set(offset.Uint64(), uint64(len(value.Bytes())), value.Bytes())
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, nil)
        case PUSH1, PUSH2, PUSH3, PUSH32:
            value := make([]byte, 0)
            e.Index++
            e.MergePushN(OpCode, func(index int) {
                value = append(value, e.Code[e.Index+index])
            })
            e.Stack = e.Stack.Push(uint256.NewInt().SetBytes(value))
            e.IndexPushN(OpCode)
            e.GasConsumption(OpCode, nil)
        }
    }

    return e.Mem.Bytes(), e.Gas
}

// GasConsumption increases gas counter by the specific opcode
// and takes a secondary argument as a formlua, for more complicated
// calculations within the OP Codes
func (e *EVM) GasConsumption(opcode byte, formula func() uint64) {
    switch opcode {
    case ADD, PUSH1, PUSH2, PUSH3, PUSH32, MSTORE8:
        e.Gas += 3
    case MUL, SDIV:
        e.Gas += 5
    case EXP:
        e.Gas += 50 * formula()
    case MSTORE:
        e.Gas += formula() - e.Mem.lastGasCost
        e.Mem.lastGasCost = e.Gas
    }
}

// StackDoublePop Stack.Pop two items from the top of the
// Stack and returns those values.
func (e *EVM) StackDoublePop() (*uint256.Int, *uint256.Int) {
    var x, y *uint256.Int
    e.Stack, x = e.Stack.Pop()
    e.Stack, y = e.Stack.Pop()
    return x, y
}

// StackPushN takes a times byte which maps to an OpCode.
//
// Current accepted OpCodes for iteration are:
// {PUSH1: 1} | {PUSH2: 2} | {PUSH3: 3} | {PUSH32: 32}
func (e *EVM) MergePushN(times byte, callback func(int)) {
    pushNTimes := map[byte]int{
        PUSH1:  1,
        PUSH2:  2,
        PUSH3:  3,
        PUSH32: 32,
    }
    for i := 0; i < pushNTimes[times]; i++ {
        callback(i)
    }
}

// IndexPushN Pushes the e.Index n Times depending
// OpCode being called.
func (e *EVM) IndexPushN(times byte) int {
    switch times {
    case ADD, MUL, SDIV, EXP, PUSH1, MSTORE, MSTORE8:
        e.Index += 1
    case PUSH2:
        e.Index += 2
    case PUSH3:
        e.Index += 3
    case PUSH32:
        e.Index += 32
    }
    return e.Index
}