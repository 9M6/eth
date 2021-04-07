package rlp

import (
    "encoding/hex"
    "fmt"
    "reflect"
    "testing"
)

func TestDecode(t *testing.T) {
    mapOfEncodedMessages := map[string][]byte{
        "#1": []byte("ea8d4976616e20426a656c616a61638e4d616c697361205075736f6e6a618c536c61766b6f204a656e6963"),
        "#2": []byte("e5922034342e38313538393735343033373334319132302e3435343733343334343535353435"),
        "#3": []byte("c4c131c132"),
    }

    for key, msg := range mapOfEncodedMessages {
        encodedMsg := make([]byte, hex.DecodedLen(len(msg)))

        _, err := hex.Decode(encodedMsg, msg)
        if err != nil {
            t.Fatalf("Decode: could not convert to hexadecimal")
        }

        decodedMsg := make([]byte, 0)
        decodedMsg = append(decodedMsg, Decode(encodedMsg)...)

        fmt.Printf("#%v -> %s \n", key, decodedMsg)
    }
}

func TestDecodedLength(t *testing.T) {
    mapOfDecodedStructs := map[*Decoded]string{
        &Decoded{0, 1, reflect.String}: "80",
        &Decoded{11, 1, reflect.String}: "8b68656c6c6f20776f726c64",
        &Decoded{12, 1, reflect.Slice}: "cc8568656c6c6f85776f726c64",
    }

    for value, str := range mapOfDecodedStructs {
        decodedStr, _ := hex.DecodeString(str)
        decodedLen := messageLength(decodedStr)
        if reflect.DeepEqual(value, decodedLen) != true {
            t.Fatalf("decodedLength: the result %+v of the function does not match the test case %+v", decodedLen, value)
        }
    }
}

func TestToInteger(t *testing.T) {
    mapOfHexValues := map[int][]byte{
        1024: {0x04, 0x00},
        2048: {0x08, 0x00},
        4096: {0x10, 0x00},
        6969: {0x1B, 0x39},
        8888: {0x22, 0xB8},
        9876: {0x26, 0x94},
    }

    for key, value := range mapOfHexValues {
        res := toInteger(value)
        if res != key {
            t.Fatalf("toInteger: the result %v does not match the expected value %v", res, key)
        }
    }
}