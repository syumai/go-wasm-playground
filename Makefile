.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./time/index.wasm ./time/
	GOOS=js GOARCH=wasm go build -o ./classify/index.wasm ./classify/
