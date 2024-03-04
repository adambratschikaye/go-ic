package main

import (
	"unsafe"
)

//go:wasmimport ic0 msg_arg_data_size
func arg_size() int32

//go:wasmimport ic0 msg_arg_data_copy
func arg_copy(dst int32, offset int32, size int32)

//go:wasmimport ic0 msg_reply_data_append
func reply_append(src int32, size int32)

//go:wasmimport ic0 msg_reply
func reply()

func main() {}

//go:export canister_update.run
func run() {
	size := arg_size()
	arg := make([]byte, size)
	arg_copy(int32(uintptr(unsafe.Pointer(&arg[0]))), 0, size)

	response := []byte("Hello ")
	response = append(response, arg[:]...)
	response = append(response, '!')
	reply_append(int32(uintptr(unsafe.Pointer(&response[0]))), int32(len(response)))
	reply()
}
