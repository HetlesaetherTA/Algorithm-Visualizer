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

// Return value 0 = no change made, already sorted
// Return vale 1 = change made, array not sorted
func sort(this js.Value, p []js.Value) interface{} {
	jsArray := p[0]
	h := &MaxHeap{}

	for i := 0; i < jsArray.Length(); i++ {
		h.append(jsArray.Index(i).Int())
	}

	itr := 0
	var returnValue []int

	if len(h.array) != 1 {
		h.swap(0, len(h.array)-1)
		returnValue = append([]int{h.array[len(h.array)-1]}, returnValue...)
		h.array = h.array[:len(h.array)-1]

		h.heapifyDown(0)
		time.Sleep(500 * time.Millisecond)
		fmt.Println(itr, ": ", h.array, " + ", returnValue)
		itr++
	} else {
		return 1
	}

	fmt.Println("done: ", append(h.array, returnValue...))

	for i, val := range returnValue {
		jsArray.SetIndex(i, val)
	}

	return jsArray
}

func (h *MaxHeap) appendJs(this js.Value, p []js.Value) interface{} {
	h.append(p[0].Int())
	return nil
}

func (h *MaxHeap) heapifyDownJs(this js.Value, p []js.Value) interface{} {
	h.heapifyDown(p[0].Int())
	return nil
}

func (h *MaxHeap) swapJs(this js.Value, p []js.Value) interface{} {
	fmt.Println(p[0].Int(), " swapped with ", p[1].Int())
	h.swap(p[0].Int(), p[1].Int())
	return nil
}

func (h *MaxHeap) removeLastJs(this js.Value, p []js.Value) interface{} {
	h.array = h.array[:len(h.array)-1]
	return nil
}

func (h *MaxHeap) getArrayJs(this js.Value, p []js.Value) interface{} {
	arr := js.Global().Get("Array").New(len(h.array))
	for i, v := range h.array {
		arr.SetIndex(i, js.ValueOf(v))
	}
	return arr
}

func newHeap(this js.Value, p []js.Value) interface{} {
	h := &MaxHeap{}
	return js.ValueOf(map[string]interface{}{
		"append":      js.FuncOf(h.appendJs),
		"heapifyDown": js.FuncOf(h.heapifyDownJs),
		"swap":        js.FuncOf(h.swapJs),
		"getArray":    js.FuncOf(h.getArrayJs),
		"removeLast":  js.FuncOf(h.removeLastJs),
	})
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("newHeap", js.FuncOf(newHeap))
	<-c
}
