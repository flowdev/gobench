package main

import (
	"fmt"
)


// -------- single op

type IncOp struct {
	outPort func(int)
}

func NewIncOp() IncOp {
	return IncOp{}
}

func (o IncOp) InPort(i int) {
	o.outPort(i + 1)
}

func (o *IncOp) SetOutPort(outPort func(int)) {
	o.outPort = outPort
}


// -------- main

func main() {
	var result int
	op := NewIncOp()
	(&op).SetOutPort(func(i int){ result = i })
	op.InPort(0)
	fmt.Println("1:\t", result)
}
