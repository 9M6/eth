package rlp

import (
    "fmt"
    "reflect"
    "strconv"
)

// Decoding of rlp Encoding
type Decoded struct {
    Len, Off int
    Type     reflect.Kind
}

// Decode takes a dst of type []byte and src of type []byte
// and decodes the RLP String
func Decode(msg []byte) []byte {
    dst := make([]byte, 0)
    dst = decode(dst, msg)
    return dst
}

// decode takes a dst of type []byte and src of type []byte
// and decodes the RLP String. This function is the internal
// implementation for decoding RLP data types.
func decode(dst, msg []byte) []byte {
    if len(msg) == 0 {
        return dst
    }

    message := messageLength(msg)

    switch message.Type {
    case reflect.String:
        dst = append(dst, toString(msg[message.Off:message.Off+message.Len])...)
        return decode(dst, msg[message.Off+message.Len:])
    case reflect.Slice:
        dst = toSlice(decode(dst, msg[message.Off:message.Off+message.Len]))
    }

    return decode(dst, msg[message.Off+message.Len:])
}

// messageLength are the Rule-Cases for the RLP decoding function
// the rules are as following:
//
// 1. the data is a string if the range of the first byte(i.e. prefix) is [0x00, 0x7f],
//    and the string is the first byte itself exactly;
//
// 2. the data is a string if the range of the first byte is [0x80, 0xb7], and the string
//    whose length is equal to the first byte minus 0x80 follows the first byte;
//
// 3. the data is a string if the range of the first byte is [0xb8, 0xbf], and the length
//    of the string whose length in bytes is equal to the first byte minus 0xb7 follows the
//    first byte, and the string follows the length of the string;
//
// 4. the data is a list if the range of the first byte is [0xc0, 0xf7], and the concatenation
//    of the RLP encodings of all items of the list which the total payload is equal to the first
//    byte minus 0xc0 follows the first byte;
//
// 5. the data is a list if the range of the first byte is [0xf8, 0xff], and the total payload
//    of the list whose length is equal to the first byte minus 0xf7 follows the first byte,
//    and the concatenation of the RLP encodings of all items of the list follows the total payload
//    of the list;
func messageLength(msg []byte) *Decoded {
    length := len(msg)
    prefix := int(msg[0])

    switch {
    case prefix <= 0x7F:
        return &Decoded{
            Off:  0,
            Len:  1,
            Type: reflect.String,
        }
    case (prefix <= 0xB7) && (length > prefix-0x80):
        return &Decoded{
            Off:  1,
            Len:  prefix - 0x80,
            Type: reflect.String,
        }
    case (prefix <= 0xBF) && length > prefix-0xB7 && length > hLen(prefix, 0xB7, msg):
        return &Decoded{
            Off:  prefix - 0xB7 + 1,
            Len:  toInteger(msg[1 : prefix-0xB7+1]),
            Type: reflect.String,
        }
    case (prefix <= 0xF7) && length > prefix-0xC0:
        return &Decoded{
            Off:  1,
            Len:  prefix - 0xC0,
            Type: reflect.Slice,
        }
    case (prefix <= 0xFF) && length > prefix-0xF7 && length > hLen(prefix, 0xF7, msg):
        return &Decoded{
            Off:  prefix - 0xF7 + 1,
            Len:  toInteger(msg[1 : prefix-0xF7+1]),
            Type: reflect.Slice,
        }
    }
    return nil
}

// toInteger takes a []byte of hexadecimal values and concatenates
// them and finally converts it to decimal representation.
func toInteger(hex []byte) int {
    length := len(hex)
    switch {
    case length == 0:
        panic("toInteger: hex []byte parameter is empty")
    case length == 1:
        return int(hex[0])
    case length > 1:
        i, err := strconv.ParseInt(fmt.Sprintf("%x", hex), 16, 0)
        if err != nil {
            panic("toInteger: can't convert hex to int type")
        }
        return int(i)
    }
    return 0
}

// toString converts a an internal representation of []bytes
// to a formatted []bytes of strings
func toString(msg []byte) []byte {
    return []byte(fmt.Sprintf(" String %s ", msg))
}

// toSlice converts a an internal representation of []bytes
// to a formatted []bytes of strings
func toSlice(msg []byte) []byte {
    return []byte(fmt.Sprintf("List { %8s }", msg))
}

// hLen is an utility function to abstract away a rule-case
// for the RLP length decoding.
// Keep in mind that in Go that slicing a slice with the
// operator msg[low:high] the low parameter is inclusive, the high
// parameter is exclusive, meaning its not going to be included,
// thus we need to bump the range by 1: msg[1:(prefix-ref+1)]
func hLen(prefix, ref int, msg []byte) int {
    return prefix - ref + toInteger(msg[1:(prefix-ref+1)])
}