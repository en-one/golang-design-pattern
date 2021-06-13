package main

import "fmt"

type Fruit interface {
	Introduce()
}

type Apple struct{}

func (A Apple) Introduce() {
	fmt.Println("i was an apple")
}

type Banana struct{}

func (B Banana) Introduce() {
	fmt.Println("i want a banner")
}

type FruitAbstractFactory interface {
	CreateApple() Fruit
	CreateBanner() Fruit
}

type SimpleFruitFactory struct{}

func NewSimpleFruitFactory() FruitAbstractFactory {
	return &SimpleFruitFactory{}
}

func (s *SimpleFruitFactory) CreateApple() Fruit {
	return &Apple{}
}

func (s *SimpleFruitFactory) CreateBanner() Fruit {
	return &Banana{}
}

func main() {
	fFactory := NewSimpleFruitFactory()
	apple := fFactory.CreateApple()
	apple.Introduce()
}
