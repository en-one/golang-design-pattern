package main

import (
	"fmt"
	"strconv"
)

type Handler interface {
	Handler(handlerID int) string
}

type handler struct {
	name      string
	next      Handler
	handlerID int
}

func NewHandler(name string, next Handler, handlerId int) *handler {
	return &handler{name: name, next: next, handlerID: handlerId}
}

func (h *handler) Handler(handlerId int) string {
	if h.handlerID == handlerId {
		return h.name + " handled " + strconv.Itoa(handlerId)
	}

	if h.next == nil {
		return ""
	}

	return h.next.Handler(handlerId)
}

func main() {
	wang := NewHandler("lao wang", nil, 1)
	zhang := NewHandler("lao zhang", nil, 2)

	r := wang.Handler(1)
	fmt.Println(r)

	r = zhang.Handler(2)
	fmt.Println(r)
}
