package main

import _ "github.com/wasilibs/nottinygc"

/*
#cgo CFLAGS: -Ilib
#cgo LDFLAGS: ./target/wasm32-wasi/release/librust_lib.a

#include "rust_lib.h"
*/
import "C"

func main() {
	seed := []uint8{0, 1, 2, 3}
	C.init((*C.uchar)(&seed[0]), C.size_t(len(seed)))
}

//export "canister_update.run"
func run() {
	println(C.number())
}
