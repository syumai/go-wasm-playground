.PHONY: build
build:
#	GOOS=js GOARCH=wasm go build -o ./error/index.wasm ./error/
#	GOOS=js GOARCH=wasm go build -o ./success/index.wasm ./success/
	GOOS=js GOARCH=wasm go build -o ./time/index.wasm ./time/
