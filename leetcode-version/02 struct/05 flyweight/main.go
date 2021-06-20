package main

import "fmt"

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{
		Name: name,
	}
}

//享元：通过工厂模式，在工厂类中，通过一个map来缓存已经创建过的享元对象

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{
		pool: make(map[string]*FlyWeight),
	}
}

func (f *FlyWeightFactory) GetFlayWeight(name string) *FlyWeight {
	weight, ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		f.pool[name] = weight
	}
	return weight
}

func main() {
	factory := NewFlyWeightFactory()
	zhang := factory.GetFlayWeight("zhang is beauty")
	//chen := factory.GetFlayWeight("chen is better beauty")

	fmt.Println(factory.pool["zhang is beauty"], zhang)
}
