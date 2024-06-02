package main

import (
	"syscall/js"
)

func add(this js.Value, p []js.Value) interface{} {
	a := p[0].Int()
	b := p[1].Int()
	result := a + b
	return js.ValueOf(result)
}

func sub(this js.Value, p []js.Value) interface{} {
	a := p[0].Int()
	b := p[1].Int()
	result := a - b
	return js.ValueOf(result)
}

func registerCallbacks() {
	js.Global().Set("sub", js.FuncOf(sub))
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
