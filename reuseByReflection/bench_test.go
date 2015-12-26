package main

import "testing"

func BenchmarkIncOp(b *testing.B) {
	op := NewIncOp()
	op.SetOutPort(func(data interface{}){ })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		op.InPort(flowData)
	}
}

func BenchmarkIncOp(b *testing.B) {
	var result int
	op := NewIncOp()
	op.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		op.InPort(flowData)
	}
}

func BenchmarkInc10Flow(b *testing.B) {
	var result int
	flow := NewInc10Flow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}

func BenchmarkInc100Flow(b *testing.B) {
	var result int
	flow := NewInc100Flow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}

func BenchmarkInc1KFlow(b *testing.B) {
	var result int
	flow := NewInc1KFlow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}

func BenchmarkInc10KFlow(b *testing.B) {
	var result int
	flow := NewInc10KFlow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}

func BenchmarkInc100KFlow(b *testing.B) {
	var result int
	flow := NewInc100KFlow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}

func BenchmarkInc1MFlow(b *testing.B) {
	var result int
	flow := NewInc1MFlow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}

/*
func BenchmarkInc10MFlow(b *testing.B) {
	var result int
	flow := NewInc10MFlow()
	flow.SetOutPort(func(data interface{}){ result = (data.(*FlowData)).N })
	flowData := &FlowData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowData.N = 0
		flow.InPort(flowData)
	}
}
*/
