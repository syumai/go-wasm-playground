## go-wasm-playground

* https://syumai.github.io/go-wasm-playground/success/
* https://syumai.github.io/go-wasm-playground/error/

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
