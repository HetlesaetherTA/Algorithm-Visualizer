package main

import (
	"fmt"
	"syscall/js"
	"time"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) append(val int) {
	h.array = append(h.array, val)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 {
		if h.array[i] > h.array[h.parent(i)] {
			h.swap(i, h.parent(i))
			i = h.parent(i)
		} else {
			break
		}
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	for h.right(i) < len(h.array) {
		if h.array[i] < h.array[h.left(i)] {
			h.swap(i, h.left(i))
			i = h.left(i)
		} else if h.array[i] < h.array[h.right(i)] {
			h.swap(i, h.right(i))
			i = h.right(i)
		} else {
			break
		}
	}
}

func (h *MaxHeap) swap(i int, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) left(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) right(i int) int {
	return 2*i + 2
}

func temp(this js.Value, p []js.Value) interface{} {
	jsArray := p[0]
	fmt.Println("HelloWorld")
	return jsArray
}

func sort(this js.Value, p []js.Value) interface{} {
	jsArray := p[0]
	h := &MaxHeap{}

	for i := 0; i < jsArray.Length(); i++ {
		h.append(jsArray.Index(i).Int())
	}

	itr := 0
	var returnValue []int

	for {
		if len(h.array) != 1 {
			h.swap(0, len(h.array)-1)
			returnValue = append([]int{h.array[len(h.array)-1]}, returnValue...)
			h.array = h.array[:len(h.array)-1]
			draw([]int{4, 5, 3, 2, 40, 60, 50, 70, 20, 90, 10, 40})

			h.heapifyDown(0)
			time.Sleep(500 * time.Millisecond)
			fmt.Println(itr, ": ", h.array, " + ", returnValue)
			itr++
		} else {
			break
		}
	}

	fmt.Println("done: ", append(h.array, returnValue...))

	draw(append(h.array, returnValue...))

	for i, val := range returnValue {
		jsArray.SetIndex(i, val)
	}

	return jsArray
}

func draw(arr []int) {
	val := js.Global().Get("Array").New(len(arr))
	for i, v := range arr {
		val.SetIndex(i, js.ValueOf(v))
	}
	js.Global().Call("draw", val)
}

func registerCallbacks() {
	js.Global().Set("sort", js.FuncOf(sort))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
