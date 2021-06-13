package main

import "fmt"

type FruitFactoryMethod interface {
	Introduce() FruitFactoryMethod
}

type AppleFactory struct{}

func (AF AppleFactory) Introduce() FruitFactoryMethod {
	fmt.Println("i was an apple factory")
	return AppleFactory{}
}

type BananaFactory struct{}

func (BF BananaFactory) Introduce() FruitFactoryMethod {
	fmt.Println("i was a banner factory")
	return BananaFactory{}
}

//每个产品都有一个专属工厂。比如苹果有专属的苹果工厂，梨子有专属的梨子工厂
func createFactory(fruit string) FruitFactoryMethod {
	switch {
	case fruit == "Apple":
		return AppleFactory{}
	case fruit == "Banner":
		return BananaFactory{}
	}
	return nil
}

//我们只需要告诉工厂 我们需要什么水果 得到什么水果工厂
func main() {
	fruitFactoryMark := "Banner"
	fFactory := createFactory(fruitFactoryMark)
	fFactory.Introduce()
}
