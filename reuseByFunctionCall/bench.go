package main

import (
	"fmt"
)


// -------- Inc10MFlow

type Inc10MFlow struct {
	op01      *Inc1MFlow
	op02      *Inc1MFlow
	op03      *Inc1MFlow
	op04      *Inc1MFlow
	op05      *Inc1MFlow
	op06      *Inc1MFlow
	op07      *Inc1MFlow
	op08      *Inc1MFlow
	op09      *Inc1MFlow
	op10      *Inc1MFlow
	InPort     func(interface{})
}

func NewInc10MFlow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc10MFlow {
	f := &Inc10MFlow{NewInc1MFlow(getInt, setInt), NewInc1MFlow(getInt, setInt), NewInc1MFlow(getInt, setInt),
		NewInc1MFlow(getInt, setInt), NewInc1MFlow(getInt, setInt), NewInc1MFlow(getInt, setInt),
		NewInc1MFlow(getInt, setInt), NewInc1MFlow(getInt, setInt), NewInc1MFlow(getInt, setInt),
		NewInc1MFlow(getInt, setInt), nil}

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

	return f
}

func (f *Inc10MFlow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- Inc1MFlow

type Inc1MFlow struct {
	op01      *Inc100KFlow
	op02      *Inc100KFlow
	op03      *Inc100KFlow
	op04      *Inc100KFlow
	op05      *Inc100KFlow
	op06      *Inc100KFlow
	op07      *Inc100KFlow
	op08      *Inc100KFlow
	op09      *Inc100KFlow
	op10      *Inc100KFlow
	InPort     func(interface{})
}

func NewInc1MFlow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc1MFlow {
	f := &Inc1MFlow{NewInc100KFlow(getInt, setInt), NewInc100KFlow(getInt, setInt), NewInc100KFlow(getInt, setInt),
		NewInc100KFlow(getInt, setInt), NewInc100KFlow(getInt, setInt), NewInc100KFlow(getInt, setInt),
		NewInc100KFlow(getInt, setInt), NewInc100KFlow(getInt, setInt), NewInc100KFlow(getInt, setInt),
		NewInc100KFlow(getInt, setInt), nil}

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

	return f
}

func (f *Inc1MFlow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- Inc100KFlow

type Inc100KFlow struct {
	op01      *Inc10KFlow
	op02      *Inc10KFlow
	op03      *Inc10KFlow
	op04      *Inc10KFlow
	op05      *Inc10KFlow
	op06      *Inc10KFlow
	op07      *Inc10KFlow
	op08      *Inc10KFlow
	op09      *Inc10KFlow
	op10      *Inc10KFlow
	InPort     func(interface{})
}

func NewInc100KFlow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc100KFlow {
	f := &Inc100KFlow{NewInc10KFlow(getInt, setInt), NewInc10KFlow(getInt, setInt), NewInc10KFlow(getInt, setInt),
		NewInc10KFlow(getInt, setInt), NewInc10KFlow(getInt, setInt), NewInc10KFlow(getInt, setInt),
		NewInc10KFlow(getInt, setInt), NewInc10KFlow(getInt, setInt), NewInc10KFlow(getInt, setInt),
		NewInc10KFlow(getInt, setInt), nil}

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

	return f
}

func (f *Inc100KFlow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- Inc10KFlow

type Inc10KFlow struct {
	op01      *Inc1KFlow
	op02      *Inc1KFlow
	op03      *Inc1KFlow
	op04      *Inc1KFlow
	op05      *Inc1KFlow
	op06      *Inc1KFlow
	op07      *Inc1KFlow
	op08      *Inc1KFlow
	op09      *Inc1KFlow
	op10      *Inc1KFlow
	InPort     func(interface{})
}

func NewInc10KFlow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc10KFlow {
	f := &Inc10KFlow{NewInc1KFlow(getInt, setInt), NewInc1KFlow(getInt, setInt), NewInc1KFlow(getInt, setInt),
		NewInc1KFlow(getInt, setInt), NewInc1KFlow(getInt, setInt), NewInc1KFlow(getInt, setInt),
		NewInc1KFlow(getInt, setInt), NewInc1KFlow(getInt, setInt), NewInc1KFlow(getInt, setInt),
		NewInc1KFlow(getInt, setInt), nil}

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

	return f
}

func (f *Inc10KFlow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- Inc1KFlow

type Inc1KFlow struct {
	op01      *Inc100Flow
	op02      *Inc100Flow
	op03      *Inc100Flow
	op04      *Inc100Flow
	op05      *Inc100Flow
	op06      *Inc100Flow
	op07      *Inc100Flow
	op08      *Inc100Flow
	op09      *Inc100Flow
	op10      *Inc100Flow
	InPort     func(interface{})
}

func NewInc1KFlow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc1KFlow {
	f := &Inc1KFlow{NewInc100Flow(getInt, setInt), NewInc100Flow(getInt, setInt), NewInc100Flow(getInt, setInt),
		NewInc100Flow(getInt, setInt), NewInc100Flow(getInt, setInt), NewInc100Flow(getInt, setInt),
		NewInc100Flow(getInt, setInt), NewInc100Flow(getInt, setInt), NewInc100Flow(getInt, setInt),
		NewInc100Flow(getInt, setInt), nil}

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

	return f
}

func (f *Inc1KFlow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- Inc100Flow

type Inc100Flow struct {
	op01      *Inc10Flow
	op02      *Inc10Flow
	op03      *Inc10Flow
	op04      *Inc10Flow
	op05      *Inc10Flow
	op06      *Inc10Flow
	op07      *Inc10Flow
	op08      *Inc10Flow
	op09      *Inc10Flow
	op10      *Inc10Flow
	InPort     func(interface{})
}

func NewInc100Flow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc100Flow {
	f := &Inc100Flow{NewInc10Flow(getInt, setInt), NewInc10Flow(getInt, setInt), NewInc10Flow(getInt, setInt),
		NewInc10Flow(getInt, setInt), NewInc10Flow(getInt, setInt), NewInc10Flow(getInt, setInt),
		NewInc10Flow(getInt, setInt), NewInc10Flow(getInt, setInt), NewInc10Flow(getInt, setInt),
		NewInc10Flow(getInt, setInt), nil}

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

	return f
}

func (f *Inc100Flow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- Inc10Flow

type Inc10Flow struct {
	op01      *IncOp
	op02      *IncOp
	op03      *IncOp
	op04      *IncOp
	op05      *IncOp
	op06      *IncOp
	op07      *IncOp
	op08      *IncOp
	op09      *IncOp
	op10      *IncOp
	InPort     func(interface{})
}

func NewInc10Flow(getInt func(interface{})int, setInt func(interface{},int)interface{}) *Inc10Flow {
	f := &Inc10Flow{NewIncOp(getInt, setInt), NewIncOp(getInt, setInt), NewIncOp(getInt, setInt),
		NewIncOp(getInt, setInt), NewIncOp(getInt, setInt), NewIncOp(getInt, setInt),
		NewIncOp(getInt, setInt), NewIncOp(getInt, setInt), NewIncOp(getInt, setInt),
		NewIncOp(getInt, setInt), nil}

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

	return f
}

func (f *Inc10Flow) SetOutPort(outPort func(interface{})) {
	f.op10.SetOutPort(outPort)
}


// -------- single op

type IncOp struct {
	getInt func(interface{})int
	setInt func(interface{},int)interface{}
	outPort func(interface{})
}

func NewIncOp(getInt func(interface{})int, setInt func(interface{},int)interface{}) *IncOp {
	return &IncOp{getInt, setInt, nil}
}

func (o *IncOp) InPort(data interface{}) {
	o.outPort(o.setInt(data, o.getInt(data) + 1))
}

func (o *IncOp) SetOutPort(outPort func(interface{})) {
	o.outPort = outPort
}


// -------- main

func main() {
	var result int
	flow10 := NewInc10Flow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow10.SetOutPort(func(data interface{}){ result = data.(int) })
	flow10.InPort(0)
	fmt.Println("10:\t", result)

	flow100 := NewInc100Flow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow100.SetOutPort(func(data interface{}){ result = data.(int) })
	flow100.InPort(0)
	fmt.Println("100:\t", result)

	flow1K := NewInc1KFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow1K.SetOutPort(func(data interface{}){ result = data.(int) })
	flow1K.InPort(0)
	fmt.Println("1K:\t", result)

	flow10K := NewInc10KFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow10K.SetOutPort(func(data interface{}){ result = data.(int) })
	flow10K.InPort(0)
	fmt.Println("10K:\t", result)

	flow100K := NewInc100KFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow100K.SetOutPort(func(data interface{}){ result = data.(int) })
	flow100K.InPort(0)
	fmt.Println("100K:\t", result)

	flow1M := NewInc1MFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow1M.SetOutPort(func(data interface{}){ result = data.(int) })
	flow1M.InPort(0)
	fmt.Println("1M:\t", result)
}
