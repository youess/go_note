package main

import (
	"fmt"
)

type A struct {
	id    int
	attr1 string
	attr2 string
}

func NewA(id int, a1 string, a2 string) *A {
	return &A{
		id:    id,
		attr1: a1,
		attr2: a2,
	}
}

type B struct {
	a     *A
	bAttr string
}

func NewB(a *A, desc string) *B {
	return &B{
		a:     a,
		bAttr: desc,
	}
}

func main() {
	a := NewA(1, "attribute-1", "attribute-2")
	a.attr2 = "attribute-2-modified-outside"
	fmt.Printf("struct A: %s\n", a)
	b := NewB(a, "b desc")
	b.a.attr1 = "change from b instance"
	fmt.Printf("struct B: %s\n", b)
	/*
		struct A: &{%!s(int=1) attribute-1 attribute-2-modified-outside}
		struct B: &{%!s(*main.A=&{1 change from b instance attribute-2-modified-outside}) b desc}
	*/
}
