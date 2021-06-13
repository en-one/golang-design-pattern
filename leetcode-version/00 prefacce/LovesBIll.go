package main

import "fmt"

//继承

type LovesDayBill struct {
	NormalBill
	FeatureA string
}

func (l *LovesDayBill) SayLove() {
	fmt.Println("i want say love")
}

func (l *LovesDayBill) disCount() float64 {
	return l.price * 0.7
}
