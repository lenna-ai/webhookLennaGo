package test

import (
	"fmt"
	"testing"
	"webhooklenna/simples"
)

func TestSimpleService(t *testing.T) {
	data := simples.InitializeService(false)
	fmt.Println(*data)
}

func TestInterfaceBinding(t *testing.T) {
	simpleService := simples.NewHelloService(&simples.SayHelloImpl{})

	fmt.Println(simpleService)
}
