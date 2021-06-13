package main

type NormalBill struct {
	number int
	unit   float64
	price  float64
}

func NewBill() *NormalBill {
	return &NormalBill{}
}

func (b *NormalBill) SetNumber(number int) {
	b.number = number
}

func (b *NormalBill) SetUnit(unit float64) {
	b.unit = unit
}

func (b *NormalBill) SetPrice() {
	b.price = b.unit * float64(b.number)
}

func (b *NormalBill) disCount() float64 {
	return b.price
}
