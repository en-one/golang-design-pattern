package main

import (
	"fmt"
	"strconv"
)

type ITAsk interface {
	RentHouse(desc string, price int)
}

type Task struct{}

func (t *Task) RendHouse(desc string, price int) {
	fmt.Sprintln(fmt.Printf("租房子的地址%s,价钱%s,中介费%s.", desc, strconv.Itoa(price), strconv.Itoa(price)))
}

//代理模式：实现方式 继承（同装饰模式）

type AgentTask struct {
	task *Task
}

func NewAgentTask() AgentTask {
	return AgentTask{
		task: &Task{},
	}
}

func (a *AgentTask) RentHouse(desc string, price int) {
	a.task.RendHouse(desc, price)
}

func main() {
	agent := NewAgentTask()
	agent.RentHouse("北京", 8000)
}
