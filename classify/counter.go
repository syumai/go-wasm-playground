package main

import "syscall/js"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Decrement() {
	c.count--
}

func (c *Counter) Count() int {
	return c.count
}

func NewCounter(this js.Value, args []js.Value) interface{} {
	c := &Counter{}
	this.Set("increment", toFunc(c.Increment))
	this.Set("decrement", toFunc(c.Decrement))
	this.Set("count", toIntFunc(c.Count))
	return this
}
