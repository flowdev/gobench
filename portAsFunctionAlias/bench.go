package main

import (
	"fmt"
)

// -------- Inc10MFlow

type Inc10MFlow struct {
	op01       *Inc1MFlow
	op02       *Inc1MFlow
	op03       *Inc1MFlow
	op04       *Inc1MFlow
	op05       *Inc1MFlow
	op06       *Inc1MFlow
	op07       *Inc1MFlow
	op08       *Inc1MFlow
	op09       *Inc1MFlow
	op10       *Inc1MFlow
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc10MFlow() *Inc10MFlow {
	f := &Inc10MFlow{NewInc1MFlow(), NewInc1MFlow(), NewInc1MFlow(),
		NewInc1MFlow(), NewInc1MFlow(), NewInc1MFlow(), NewInc1MFlow(),
		NewInc1MFlow(), NewInc1MFlow(), NewInc1MFlow(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- Inc1MFlow

type Inc1MFlow struct {
	op01       *Inc100KFlow
	op02       *Inc100KFlow
	op03       *Inc100KFlow
	op04       *Inc100KFlow
	op05       *Inc100KFlow
	op06       *Inc100KFlow
	op07       *Inc100KFlow
	op08       *Inc100KFlow
	op09       *Inc100KFlow
	op10       *Inc100KFlow
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc1MFlow() *Inc1MFlow {
	f := &Inc1MFlow{NewInc100KFlow(), NewInc100KFlow(), NewInc100KFlow(),
		NewInc100KFlow(), NewInc100KFlow(), NewInc100KFlow(), NewInc100KFlow(),
		NewInc100KFlow(), NewInc100KFlow(), NewInc100KFlow(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- Inc100KFlow

type Inc100KFlow struct {
	op01       *Inc10KFlow
	op02       *Inc10KFlow
	op03       *Inc10KFlow
	op04       *Inc10KFlow
	op05       *Inc10KFlow
	op06       *Inc10KFlow
	op07       *Inc10KFlow
	op08       *Inc10KFlow
	op09       *Inc10KFlow
	op10       *Inc10KFlow
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc100KFlow() *Inc100KFlow {
	f := &Inc100KFlow{NewInc10KFlow(), NewInc10KFlow(), NewInc10KFlow(),
		NewInc10KFlow(), NewInc10KFlow(), NewInc10KFlow(), NewInc10KFlow(),
		NewInc10KFlow(), NewInc10KFlow(), NewInc10KFlow(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- Inc10KFlow

type Inc10KFlow struct {
	op01       *Inc1KFlow
	op02       *Inc1KFlow
	op03       *Inc1KFlow
	op04       *Inc1KFlow
	op05       *Inc1KFlow
	op06       *Inc1KFlow
	op07       *Inc1KFlow
	op08       *Inc1KFlow
	op09       *Inc1KFlow
	op10       *Inc1KFlow
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc10KFlow() *Inc10KFlow {
	f := &Inc10KFlow{NewInc1KFlow(), NewInc1KFlow(), NewInc1KFlow(),
		NewInc1KFlow(), NewInc1KFlow(), NewInc1KFlow(), NewInc1KFlow(),
		NewInc1KFlow(), NewInc1KFlow(), NewInc1KFlow(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- Inc1KFlow

type Inc1KFlow struct {
	op01       *Inc100Flow
	op02       *Inc100Flow
	op03       *Inc100Flow
	op04       *Inc100Flow
	op05       *Inc100Flow
	op06       *Inc100Flow
	op07       *Inc100Flow
	op08       *Inc100Flow
	op09       *Inc100Flow
	op10       *Inc100Flow
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc1KFlow() *Inc1KFlow {
	f := &Inc1KFlow{NewInc100Flow(), NewInc100Flow(), NewInc100Flow(),
		NewInc100Flow(), NewInc100Flow(), NewInc100Flow(), NewInc100Flow(),
		NewInc100Flow(), NewInc100Flow(), NewInc100Flow(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- Inc100Flow

type Inc100Flow struct {
	op01       *Inc10Flow
	op02       *Inc10Flow
	op03       *Inc10Flow
	op04       *Inc10Flow
	op05       *Inc10Flow
	op06       *Inc10Flow
	op07       *Inc10Flow
	op08       *Inc10Flow
	op09       *Inc10Flow
	op10       *Inc10Flow
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc100Flow() *Inc100Flow {
	f := &Inc100Flow{NewInc10Flow(), NewInc10Flow(), NewInc10Flow(),
		NewInc10Flow(), NewInc10Flow(), NewInc10Flow(), NewInc10Flow(),
		NewInc10Flow(), NewInc10Flow(), NewInc10Flow(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- Inc10Flow

type Inc10Flow struct {
	op01       *IncOp
	op02       *IncOp
	op03       *IncOp
	op04       *IncOp
	op05       *IncOp
	op06       *IncOp
	op07       *IncOp
	op08       *IncOp
	op09       *IncOp
	op10       *IncOp
	InPort     func(int)
	SetOutPort func(func(int))
}

func NewInc10Flow() *Inc10Flow {
	f := &Inc10Flow{NewIncOp(), NewIncOp(), NewIncOp(), NewIncOp(), NewIncOp(),
		NewIncOp(), NewIncOp(), NewIncOp(), NewIncOp(), NewIncOp(), nil, nil}

	f.op01.SetOutPort(f.op02.InPort)
	f.op02.SetOutPort(f.op03.InPort)
	f.op03.SetOutPort(f.op04.InPort)
	f.op04.SetOutPort(f.op05.InPort)
	f.op05.SetOutPort(f.op06.InPort)
	f.op06.SetOutPort(f.op07.InPort)
	f.op07.SetOutPort(f.op08.InPort)
	f.op08.SetOutPort(f.op09.InPort)
	f.op09.SetOutPort(f.op10.InPort)
	f.InPort = f.op01.InPort
	f.SetOutPort = f.op10.SetOutPort

	return f
}

// -------- single op

type IncOp struct {
	outPort func(int)
}

func NewIncOp() *IncOp {
	return &IncOp{}
}

func (o *IncOp) InPort(i int) {
	o.outPort(i + 1)
}

func (o *IncOp) SetOutPort(outPort func(int)) {
	o.outPort = outPort
}

// -------- main

func main() {
	var result int
	flow10 := NewInc10Flow()
	flow10.SetOutPort(func(i int) { result = i })
	flow10.InPort(0)
	fmt.Println("10:\t", result)

	flow100 := NewInc100Flow()
	flow100.SetOutPort(func(i int) { result = i })
	flow100.InPort(0)
	fmt.Println("100:\t", result)

	flow1K := NewInc1KFlow()
	flow1K.SetOutPort(func(i int) { result = i })
	flow1K.InPort(0)
	fmt.Println("1K:\t", result)

	flow10K := NewInc10KFlow()
	flow10K.SetOutPort(func(i int) { result = i })
	flow10K.InPort(0)
	fmt.Println("10K:\t", result)

	flow100K := NewInc100KFlow()
	flow100K.SetOutPort(func(i int) { result = i })
	flow100K.InPort(0)
	fmt.Println("100K:\t", result)

	flow1M := NewInc1MFlow()
	flow1M.SetOutPort(func(i int) { result = i })
	flow1M.InPort(0)
	fmt.Println("1M:\t", result)

	flow10M := NewInc10MFlow()
	flow10M.SetOutPort(func(i int) { result = i })
	flow10M.InPort(0)
	fmt.Println("10M:\t", result)
}
