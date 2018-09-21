package main

import "syscall/js"

func showAlert() {
	js.Global().Call("alert", "show alert!!!!")
}

func main() {
	btn := js.Global().Get("document").Call("getElementById", "btn")
	btn.Call("setAttribute", "onclick", js.NewCallback(func ([]js.Value) { showAlert() }))
	select {}
}

