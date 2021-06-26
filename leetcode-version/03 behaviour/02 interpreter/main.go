package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(variables map[string]Expression) int
}

type Integer struct {
	integer int
}

func (i *Integer) Interpret(variables map[string]Expression) int {
	return i.integer
}

type Plus struct {
	LeftOperand  Expression
	rightOperand Expression
}

func (p *Plus) Interpret(variables map[string]Expression) int {
	return p.LeftOperand.Interpret(variables) + p.rightOperand.Interpret(variables)
}

type Variable struct {
	name string
}

func (v *Variable) Interpret(variables map[string]Expression) int {
	value, found := variables[v.name]
	if !found {
		return 0
	}
	return value.Interpret(variables)
}

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{value: value, next: s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

type Evaluator struct {
	syntaxTree Expression
}

func (e *Evaluator) Interpret(context map[string]Expression) int {
	return e.syntaxTree.Interpret(context)
}

func NewEvaluator(expression string) *Evaluator {
	expressionStack := new(Stack)
	for _, token := range strings.Split(expression, " ") {
		switch token {
		case "+":
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left, right}
			expressionStack.Push(subExpression)
		default:
			expressionStack.Push(&Variable{token})
		}
	}
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree: syntaxTree}
}

func main() {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Integer{6}
	variables["x"] = &Integer{10}
	variables["z"] = &Integer{41}
	result := sentence.Interpret(variables)
	fmt.Println(result)
}
