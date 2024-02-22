# IC Canister in Go

## Prerequisites

1. Install Rust

2. Install Tinygo

3. Install `wasi2ic`

4. Install `dfx`

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