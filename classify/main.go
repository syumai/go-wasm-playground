package main

import "syscall/js"

func main() {
	window := js.Global()
	window.Set("Counter", js.FuncOf(NewCounter))
	select {}
}

func toFunc(f func()) js.Func {
	return js.FuncOf(func(js.Value, []js.Value) interface{} {
		f()
		return nil
	})
}

func toIntFunc(f func() int) js.Func {
	return js.FuncOf(func(js.Value, []js.Value) interface{} {
		return f()
	})
}
