package main

import (
	"fmt"
	wordchecker "jtr/internal"
	"unsafe"
)

//export decryptHashWasm
func decryptHashWasm(hashPtr *uint8, hashLen uint32, guessPtr *uint8, guessLen uint32) *uint8 {
	hashSlice := unsafe.Slice(hashPtr, hashLen)
	hashStr := string(hashSlice)

	guessSlice := unsafe.Slice(guessPtr, guessLen)
	guessStr := string(guessSlice)

	res := wordchecker.Read(hashStr, guessStr)
	var result string
	if res == "" {
		result = fmt.Sprintf("No match found for %s & %s", guessStr, hashStr)
	} else {
		result = fmt.Sprintf("Match found for %s", res)
	}
	resultPtr, _ := stringToPtr(result)
	return resultPtr
}

//export bruteForceWasm
func bruteForceWasm(hashPtr *uint8, hashLen uint32, guessPtr *uint8, guessLen uint32, charLimit int) *uint8 {
	hashSlice := unsafe.Slice(hashPtr, hashLen)
	hashStr := string(hashSlice)

	guessSlice := unsafe.Slice(guessPtr, guessLen)
	guessStr := string(guessSlice)

	res := wordchecker.BruteForce(hashStr, guessStr, 100000, charLimit)
	var result string
	if res == "" {
		result = fmt.Sprintf("No match found for %s & %s", guessStr, hashStr)
	} else {
		result = fmt.Sprintf("Match found for %s", res)
	}
	resultPtr, _ := stringToPtr(result)
	return resultPtr
}

// stringToPtr converts a Go string to a *uint8 and length
func stringToPtr(s string) (*uint8, uint32) {
	b := []byte(s)
	if len(b) == 0 {
		return nil, 0
	}
	return &b[0], uint32(len(b))
}

//export freeMemory
func freeMemory(ptr *uint8, len uint32) {
	s := unsafe.Slice(ptr, len)
	_ = s
}

func main() {}

// This function is exported to JavaScript, so can be called using
// exports.add() in JavaScript.
//
//export add
func add(x, y int) int {
	return x + y
}
