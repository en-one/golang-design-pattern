package main

import "fmt"

type Fruit interface {
	Name() string
	Clone() Fruit
}

type ConcreteFruit struct {
	name string
}

func (c *ConcreteFruit) Name() string {
	return c.name
}

func (c *ConcreteFruit) Clone() Fruit {
	return &ConcreteFruit{name: c.name}
}

//原型模式：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。
func main() {
	name := "durian"
	fruitProto := ConcreteFruit{name: name}
	newFruitProto := fruitProto.Clone()

	fmt.Println(newFruitProto.Name(), fruitProto.name)
}
