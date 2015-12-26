package main

import (
	"fmt"
)

type FlowData struct {
	N int
}


func setViaClosureWithoutReturn(data interface{}, getInt func(interface{})int, setInt func(interface{},int)) {
	setInt(data, getInt(data) + 1)
}

func setViaClosureWithReturn(data interface{}, getInt func(interface{})int,
		setInt func(interface{},int)interface{}) interface{} {
	return setInt(data, getInt(data) + 1)
}

// -------- single op with struct

type IncOpStruct struct {
	getInt func(interface{})int
	setInt func(interface{},int)interface{}
	outPort func(interface{})
}

func NewIncOpStruct(getInt func(interface{})int, setInt func(interface{},int)interface{}) *IncOpStruct {
	return &IncOpStruct{getInt, setInt, nil}
}

func (o *IncOpStruct) InPort(data interface{}) {
	o.outPort(o.setInt(data, o.getInt(data) + 1))
}

func (o *IncOpStruct) SetOutPort(outPort func(interface{})) {
	o.outPort = outPort
}


// -------- single op with pure int

type IncOpInt struct {
	getInt func(interface{})int
	setInt func(interface{},int)interface{}
	outPort func(interface{})
}

func NewIncOpInt(getInt func(interface{})int, setInt func(interface{},int)interface{}) *IncOpInt {
	return &IncOpInt{getInt, setInt, nil}
}

func (o *IncOpInt) InPort(data interface{}) {
	o.outPort(o.setInt(data, o.getInt(data) + 1))
}

func (o *IncOpInt) SetOutPort(outPort func(interface{})) {
	o.outPort = outPort
}



// -------- main

func main() {
	getN1 := func(data interface{})int{ return (data.(*FlowData)).N }
	setN1 := func(data interface{}, i int){ data.(*FlowData).N = i }
	flowData := &FlowData{}
	var data interface{}
	data = flowData
	setViaClosureWithoutReturn(data, getN1, setN1)
	fmt.Println("setViaClosureWithoutReturn:\t", flowData.N)

	getN2 := func(data interface{})int{ return data.(*FlowData).N }
	setN2 := func(data interface{}, i int)interface{}{ data.(*FlowData).N = i; return data }
	flowData = &FlowData{}
	data = flowData
	data = setViaClosureWithReturn(data, getN2, setN2)
	fmt.Println("setViaClosureWithReturn:\t", data.(*FlowData).N)

	var result int
	structOp := NewIncOpStruct(func(data interface{})int{ return data.(*FlowData).N },
		func(data interface{}, i int)interface{}{ data.(*FlowData).N = i; return data })
	structOp.SetOutPort(func(data interface{}){ result = data.(*FlowData).N })
	flowData = &FlowData{}
	data = flowData
	structOp.InPort(data)
	fmt.Println("OpWithStruct:\t", result)

	intOp := NewIncOpInt(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	intOp.SetOutPort(func(data interface{}){ result = data.(int) })
	flowData = &FlowData{}
	data = flowData
	intOp.InPort(0)
	fmt.Println("OpWithInt:\t", result)
}
