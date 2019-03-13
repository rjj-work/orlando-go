package main

import (
	"fmt"
)

// Rcvr is just an example receiver
type Rcvr struct {
	X, Y int
}

func (r Rcvr) String() string {
	return fmt.Sprintf("<%d,%d>", r.X, r.Y)
}

func (r Rcvr) whyDoThis(prefix string) {
	fmt.Printf("%s) Rcvr::whyDoThis called.\n", prefix)
}

func main() {
	aRcvr := Rcvr{X: 1, Y: 2}

	fmt.Printf("A) aRcvr.String: %s\n", aRcvr)
	fmt.Printf("B) aRcvr.String: %s\n", &aRcvr)

	aRcvr.whyDoThis("C")

	var ptrRcvr *Rcvr

	// This following line would core
	/*
	   panic: runtime error: invalid memory address or nil pointer dereference
	   [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1093616]
	*/
	//ptrRcvr.WhyDoThis()

	ptrRcvr = &Rcvr{X: -1, Y: -2}
	fmt.Printf("aRcvr.String: %s\n", ptrRcvr)

	ptrRcvr.whyDoThis("D")

}
