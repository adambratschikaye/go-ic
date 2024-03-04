# IC Canister in Go

## Warning

This is an experimental POC for deploying a canister written in Go on the IC.

## Prerequisites

1. Install [Tinygo](https://github.com/tinygo-org/tinygo) from source(https://tinygo.org/docs/guides/build/) with the following patch:
	```
	diff --git a/targets/wasm-unknown.json b/targets/wasm-unknown.json
    index 903afa49..55ca4c70 100644
    --- a/targets/wasm-unknown.json
    +++ b/targets/wasm-unknown.json
    @@ -1,7 +1,7 @@
     {
            "llvm-target":   "wasm32-unknown-unknown",
            "cpu":           "generic",
    -       "features":      "+mutable-globals,+nontrapping-fptoint,+sign-ext,-bulk-memory",
    +       "features":      "+mutable-globals,+nontrapping-fptoint,+sign-ext,+bulk-memory",
            "build-tags":    ["tinygo.wasm", "wasm_unknown"],
            "goos":          "linux",
            "goarch":        "arm",
    @@ -11,18 +11,17 @@
            "gc":            "leaking",
            "default-stack-size": 4096,
            "cflags": [
    -               "-mno-bulk-memory",
                    "-mnontrapping-fptoint",
                    "-msign-ext"
            ],
            "ldflags": [
                    "--stack-first",
                    "--no-demangle",
    -               "--no-entry",
    -               "--import-memory"
    +               "--no-entry"
            ],
            "extra-files": [
                    "src/runtime/asm_tinygowasm.S"
            ],
            "emulator":      "wasmtime --dir={tmpDir}::/tmp {}"
     }

	```
	(I used commit 377415a6c3158d5cefe8cd46c592b144053593aa)

2. Install `dfx` (recommended using [dfxvm](https://github.com/dfinity/dfxvm)).

## Steps to build and deploy

1. Start a local `dfx` subnet.
	```
	dfx start --clean --background
	```
1. Build, deploy and send a message to the canister by executing `run.sh`. This assumes `tinygo` was installed to `~/go/bin/tinygo`.