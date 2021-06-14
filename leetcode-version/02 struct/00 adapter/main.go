package main

import "fmt"

//适配器模式：将一个类的接口转换成客户希望的另外一个接口，是的原本由于接口不兼容二不能一起工作的那些类一起工作
//适配器模式并不推荐多用。因为未雨绸缪好过亡羊补牢

type Target interface {
	HelpCharge()
}

type MyPhone struct{}

func (m *MyPhone) Charge() {
	fmt.Println("手机充电")
}

//适配器给手机充电
//使用继承 调用Charge接口

type Adapter struct {
	MyPhone
}

func (a *Adapter) HelpCharge() {
	fmt.Println("适配器转换电源")
	a.Charge()
}

func main() {
	adapter := Adapter{}
	adapter.HelpCharge()
}
