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

func (b *Button) OnClick(f func()) {
	ff := js.FuncOf(func(js.Value, []js.Value) interface{} {
		f()
		return nil
	})
	b.el.Call("addEventListener", "click", ff)
	b.listeners = append(b.listeners, ff)
}

func (b *Button) Release() {
	for _, f := range b.listeners {
		f.Release()
	}
	b.listeners = nil
}
