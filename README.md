# IC Canister in Go

## Warning

This is an experimental POC for deploying a canister written in Go on the IC.
It isn't working yet.

## Prerequisites

1. Install Rust (recommended by using [rustup](https://www.rust-lang.org/tools/install)).

2. Install [Tinygo](https://tinygo.org/getting-started/install/).

3. Install [`wasi2ic`](https://github.com/wasm-forge/wasi2ic).

4. Install `dfx` (recommended using [dfxvm](https://github.com/dfinity/dfxvm)).

## Steps to build and deploy

1. Build the Rust library:
	
	```
	cargo build --target wasm32-wasi --release
	```

1. Build the Go Wasm module:

	```
	tinygo build -tags=custommalloc -gc=custom -target wasi -o main.wasm main.go
	```

1. Modify the Wasm to run on the IC:
	```
	wasmi2ic main.wasm go_canister.wasm
	```

1. Start a local `dfx` subnet and deploy the canister:
	```
	dfx start --clean --background
	dfx deploy --no-wallet
	```
	This step should try installing the canister and eventually hit an error because the canister calls `ic0.time` in the `start` function which isn't allowed.
