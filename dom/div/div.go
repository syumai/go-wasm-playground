package div

import (
	"syscall/js"

	"github.com/syumai/go-wasm-playground/dom/document"
)

type Div struct {
	el js.Value
}

func New(id string) *Div {
	return &Div{
		el: document.GetElementByID(id),
	}
}

func (div *Div) SetText(s string) {
	div.el.Set("textContent", s)
}
