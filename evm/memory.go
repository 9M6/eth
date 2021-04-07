package evm

import (
    "github.com/holiman/uint256"
)

// Mem is a simple in-memory representation using slice of bytes.
type Mem struct {
    store       []byte
    lastGasCost uint64
}

// NewMem initialises a new Mem and returns a reference to it.
func NewMem() *Mem {
    return &Mem{}
}

// Set sets a value of type []byte to store
func (m *Mem) Set(offset, size uint64, value []byte) {
    if size > 0 {
        if offset+size >= uint64(len(m.store)) {
            m.Resize(offset + size)
        }
        copy(m.store[offset:offset+size], value)
    }
}

// Set32 sets a 32bit value of type *uint256.Int to store
func (m *Mem) Set32(offset uint64, value *uint256.Int) {
    if offset+32 > uint64(len(m.store)) {
        m.Resize(offset + 32)
    }

    // Creates a []byte padded with offset + 32byte zeroes
    paddedValue := value.PaddedBytes(32)

    // Write the value bytes to the padded value
    value.WriteToSlice(paddedValue)

    // Copies everything except the sign of the *uint256.Int
    copy(m.store[offset:], paddedValue)
}

// Load pulls from Mem the value depending on the offset
func (m *Mem) Load(offset, size uint64) []byte {
    return m.store[offset : offset+size]
}

// Resize the memory store to the provided length
func (m *Mem) Resize(size uint64) {
    if size > uint64(len(m.store)) {
        m.store = append(m.store, make([]byte, size-uint64(len(m.store)))...)
    }
}

// Bytes
func (m Mem) Bytes() []byte {
    return m.store
}

// Length
func (m Mem) Length() int {
    return len(m.store)
}