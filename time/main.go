package main

import (
	"github.com/syumai/go-wasm-playground/time/timeformatter"
)

func main() {
	timeformatter.NewTimeFormatter("app")
	select {}
}
