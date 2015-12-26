package main

import (
	"testing"
	"reflect"
)

func BenchmarkClosureOpWithInt(b *testing.B) {
	getN := func(data interface{})int{ return data.(int) }
	setN := func(data interface{}, i int)interface{}{ return i }
	op := NewIncOpInt(getN, setN)
	op.SetOutPort(func(data interface{}){ })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		op.InPort(0)
	}
}

func BenchmarkNakedReflect(b *testing.B) {
	flowData := &FlowData{}
	index := []int{ 0 }
	var data interface{}
	data = flowData
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iv := reflect.ValueOf(data).Elem().FieldByIndex(index)
		var j int
		j = int(iv.Int()) + 1
		iv.SetInt(int64(j))
	}
}

func BenchmarkNakedIntClosuresWithReturn(b *testing.B) {
	getN := func(data interface{})int{ return data.(int) }
	setN := func(data interface{},i int)interface{}{ return i }
	var data interface{}
	data = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setN(data, getN(data) + 1)
	}
}
func BenchmarkNakedIntClosuresNoReturn(b *testing.B) {
	getN := func(data interface{})int{ return data.(int) }
	setN := func(data interface{}, i int){ }
	var data interface{}
	data = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setN(data, getN(data) + 1)
	}
}

func BenchmarkNakedStructClosuresNoReturn(b *testing.B) {
	getN := func(data interface{})int{ return (data.(*FlowData)).N }
	setN := func(data interface{}, i int){ data.(*FlowData).N = i }
	flowData := &FlowData{}
	var data interface{}
	data = flowData
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setN(data, getN(data) + 1)
	}
}

func BenchmarkNakedClosuresWithReturn(b *testing.B) {
	getN := func(data interface{})int{ return data.(*FlowData).N }
	setN := func(data interface{}, i int)interface{}{ data.(*FlowData).N = i; return data }
	flowData := &FlowData{}
	var data interface{}
	data = flowData
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setN(data, getN(data) + 1)
	}
}

func BenchmarkFunctionCallWithClosuresNoReturn(b *testing.B) {
	getN := func(data interface{})int{ return (data.(*FlowData)).N }
	setN := func(data interface{}, i int){ data.(*FlowData).N = i }
	flowData := &FlowData{}
	var data interface{}
	data = flowData
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setViaClosureWithoutReturn(data, getN, setN)
	}
}

func BenchmarkFunctionCallWithClosuresWithReturn(b *testing.B) {
	getN := func(data interface{})int{ return data.(*FlowData).N }
	setN := func(data interface{}, i int)interface{}{ data.(*FlowData).N = i; return data }
	flowData := &FlowData{}
	var data interface{}
	data = flowData
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setViaClosureWithReturn(data, getN, setN)
	}
}

func BenchmarkClosureOpWithStruct(b *testing.B) {
	getN := func(data interface{})int{ return data.(*FlowData).N }
	setN := func(data interface{}, i int)interface{}{ data.(*FlowData).N = i; return data }
	flowData := &FlowData{}
	var data interface{}
	data = flowData
	op := NewIncOpStruct(getN, setN)
	op.SetOutPort(func(data interface{}){ })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		op.InPort(data)
	}
}
