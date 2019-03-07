package document

import "syscall/js"

var document = js.Global().Get("document")

func GetElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func CreateElement(tagName string) js.Value {
	return document.Call("createElement", tagName)
}

func QuerySelector(query string) js.Value {
	return document.Call("querySelector", query)
}

func QuerySelectorAll(query string) js.Value {
	return document.Call("querySelectorAll", query)
}
