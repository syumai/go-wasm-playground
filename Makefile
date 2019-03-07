.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./time/index.wasm ./time/
