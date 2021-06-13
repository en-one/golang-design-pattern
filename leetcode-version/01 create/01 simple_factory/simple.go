package main

import "fmt"

type FruitSimpleFactory interface {
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

func createFruit(fruit string) FruitSimpleFactory {
	switch {
	case fruit == "Apple":
		return Apple{}
	case fruit == "Banner":
		return Banana{}
	}
	return nil
}

//我们只需要告诉工厂 我们需要什么水果 工厂就会给我们生产水果
func main() {
	fruitMark := "Banners"
	fruit := createFruit(fruitMark)
	fruit.Introduce()
}
