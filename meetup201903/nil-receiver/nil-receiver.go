package main

import (
	"fmt"
)

// Rcvr is just an example receiver
type Rcvr struct {
	X, Y int
}

func (r *Rcvr) String() string {
	if r == nil {
		return "Rcvr::String invoked with nil receiver"
	}

	return fmt.Sprintf("<%d,%d>", r.X, r.Y) /* aRcvr.String: {%!s(int=14) %!s(int=22)} */

	// x, y := r.X, r.Y
	// return fmt.Sprintf("<%d,%d>", x, y)
}

func (r *Rcvr) whyDoThis(prefix string) {
	fmt.Printf("%s) Rcvr::WhyDoThis called.\n", prefix)
}

func main() {
	aRcvr := Rcvr{X: 14, Y: 22}

	fmt.Printf("A) aRcvr.String: %s\n", aRcvr) // golint complains about
	fmt.Printf("B) aRcvr.String: %s\n", &aRcvr)

	aRcvr.whyDoThis("C")

	var ptrRcvr *Rcvr
	fmt.Printf("ptrRcvr of type(%T), value: %v\n", ptrRcvr, ptrRcvr)

	ptrRcvr.whyDoThis("D")

}
