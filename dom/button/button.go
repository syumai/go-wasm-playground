package button

import (
	"syscall/js"

	"github.com/syumai/go-wasm-playground/dom/document"
)

type Button struct {
	el        js.Value
	listeners []js.Func
}

func New(id string) *Button {
	return &Button{
		el: document.GetElementByID(id),
	}
}

func (b *Button) OnClick(fn func()) {
	f := js.FuncOf(func(js.Value, []js.Value) interface{} {
		fn()
		return nil
	})
	b.el.Call("addEventListener", "click", f)
	b.listeners = append(b.listeners, f)
}

func (b *Button) Release() {
	for _, f := range b.listeners {
		b.el.Call("removeEventListener", "click", f)
		f.Release()
	}
	b.listeners = nil
}
