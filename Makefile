GOCMD   =  go
GOBUILD =  $(GOCMD)  build
GORUN = $(GOCMD) run
TINYGOCMD   =  tinygo
TINYGOBUILD = $(TINYGOCMD) build

compile-wasm:export GOOS=wasi
compile-wasm:export GOARCH=wasm
compile-wasm:
	$(GOBUILD) -o main.wasm && rm ./html/main.wasm && cp ./main.wasm ./html/main.wasm
compile-wasm:
	$(GOBUILD) -o main.wasm && rm ./html/main.wasm && cp ./main.wasm ./html/main.wasm

compile-wasmgo:export GOOS=wasip1
compile-wasmgo:export GOARCH=wasm
compile-wasmgo:
	$(TINYGOBUILD) -o main.wasm -no-debug main.go
run_server:
	cd html && $(GORUN) server.go	