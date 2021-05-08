package simplefactory

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T)  {
	api := NewAPI(1)
	s := api.Say("tom")
	if s != "Hi, tom" {
		t.Fatal("Type1 test fail")
	}
	fmt.Println(s)
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("tom")
	if s != "Hello, tom" {
		t.Fatal("Type2 test fail")
	}
	fmt.Println(s)
}
