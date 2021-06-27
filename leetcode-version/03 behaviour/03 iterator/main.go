package main

import "fmt"

type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

type ArrayIterator struct {
	array []interface{}
	index *int
}

func (a *ArrayIterator) Index() *int {
	return a.index
}
func (a *ArrayIterator) Value() interface{} {
	return a.array[*a.index]
}
func (a *ArrayIterator) HasNext() bool {
	return *a.index+1 <= len(a.array)
}
func (a *ArrayIterator) Next() {
	if a.HasNext() {
		*a.index++
	}
}

func main() {
	array := []interface{}{1, 3, 9, 2, 8}
	a := 0
	iterator := ArrayIterator{array: array, index: &a}

	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(int)
		if value != array[*index] {
			fmt.Println("error ...")
		}
		fmt.Println(*index, value)
	}

}
