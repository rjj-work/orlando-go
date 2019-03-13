package main

import (
	"fmt"
)

type vr struct {
	v int
}

func (a vr) printMeValue() {
	fmt.Printf("printMeValue: a.v: %d\n", a.v)
	a.v++
}

func (a *vr) printMePtr() {
	fmt.Printf("printMeValue: b.v: %d\n", a.v)
	a.v++
}

func main() {
	vr1 := vr{v: 1}

	fmt.Printf("main A: vr1: %v\n", vr1)

	vr1.printMeValue()

	fmt.Printf("main B: vr1: %v\n", vr1)

	vr1.printMePtr()

	fmt.Printf("main C: vr1: %v\n", vr1)
}
