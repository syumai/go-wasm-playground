## go-wasm-playground

* https://syumai.github.io/go-wasm-playground/success/
* https://syumai.github.io/go-wasm-playground/error/
* [https://syumai.github.io/go-wasm-playground/time/](https://syumai.github.io/go-wasm-playground/time/index.html)

### Success

```go
func main() {
	btn := js.Global().Get("document").Call("getElementById", "btn")
	btn.Set("onclick", js.NewCallback(func ([]js.Value) { showAlert() }))
	select {}
}
```

### Error

```go
func main() {
	btn := js.Global().Get("document").Call("getElementById", "btn")
	btn.Call("setAttribute", "onclick", js.NewCallback(func ([]js.Value) { showAlert() }))
	select {}
}
```

```
Uncaught SyntaxError: Unexpected token (
```

### TimeFormatter

* [https://syumai.github.io/go-wasm-playground/time/](https://syumai.github.io/go-wasm-playground/time/index.html)
