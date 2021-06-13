package main

import "fmt"

func main() {

	//普通付款
	var b = NewBill()
	b.SetUnit(1.1)
	b.SetNumber(6)
	b.SetPrice()
	//measure(b)
	//fmt.Println(b.price)

	//七夕付款
	loves := LovesDayBill{
		NormalBill: *b,
	}
	measure(&loves)
	fmt.Println(loves.price)
	loves.SayLove()
}
