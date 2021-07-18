package main

import (
	"fmt"
	"sync"
	"time"
)

//事件
type Event struct {
	Data int
}

//观察者
type Observer interface {
	NotifyCallback(event Event)
}

type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

type eventObserver struct {
	ID   int
	Time time.Time
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Recieved:%d after %v\n", event.Data, time.Since(e.Time))
}

type eventSubject struct {
	Observers sync.Map
}

func (e *eventSubject) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}

func (e *eventSubject) RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}

func (e *eventSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

func main() {
	n := eventSubject{Observers: sync.Map{}}

	obs1 := eventObserver{ID: 1, Time: time.Now()}
	obs2 := eventObserver{ID: 2, Time: time.Now()}

	n.AddListener(obs1)
	n.AddListener(obs2)

	for x := range Fib(10) {
		n.Notify(Event{Data: x})
	}

	//time.Sleep(5)
	//n.RemoveListener(obs1)
	//for x := range Fib(10){
	//	n.Notify(Event{Data: x})
	//}
}
