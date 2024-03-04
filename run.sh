set -euxo pipefail

~/go/bin/tinygo build -target wasm-unknown -o main.wasm main.go
# tiny go doesn't support spaces in the wasm export names, so we export the name
# as "canister_update.run" and convert it to "canister_update run" afterward.
sed -i 's/canister_update\./canister_update /g' main.wasm

# upstream tinygo automatically imports the memory instead of defining it. If
# tinygo has been modified to define the memory then this isn't needed.
# Otherwise this part can be used to change the imported memory to a defined
# memory (with initial size of 2 pages.).
#
# wasm2wat main.wasm -o main.wat
# sed -i 's/(import "env".*$/(memory 2)/' main.wat
# wat2wasm main.wat -o main.wasm

cp main.wasm go_canister.wasm
dfx deploy --no-wallet
# The 'raw' input and output are hex strings so we use xxd to convert to hex and back
dfx canister call --type raw --output raw go_canister run $(echo "IC User" | xxd -p) | xxd -r -p