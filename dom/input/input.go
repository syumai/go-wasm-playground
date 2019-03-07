package input

import (
	"syscall/js"

	"github.com/syumai/go-wasm-playground/dom/document"
)

type TextField struct {
	el js.Value
}

func NewTextField(id string) *TextField {
	return &TextField{
		el: document.GetElementByID(id),
	}
}

func (tf *TextField) Value() string {
	return tf.el.Get("value").String()
}

func (tf *TextField) SetValue(v string) {
	tf.el.Set("value", v)
}
