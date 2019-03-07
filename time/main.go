package main

import (
	"github.com/syumai/go-wasm-playground/time/formatter"
)

func main() {
	f := formatter.New("app")
	f.Now()
	select {}
}
