GO_BIN = go1.11

.PHONY: build
build:
	GOOS=js GOARCH=wasm $(GO_BIN) build -o ./error/index.wasm ./error/
	GOOS=js GOARCH=wasm $(GO_BIN) build -o ./success/index.wasm ./success/
