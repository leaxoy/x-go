package xstrings

import (
	"unsafe"
)

func UnsafeStringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func UnsafeBytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
