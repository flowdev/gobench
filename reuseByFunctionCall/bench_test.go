package main

import "testing"

func BenchmarkIncOp(b *testing.B) {
	var result int
	op := NewIncOp(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	op.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		op.InPort(0)
	}
}

func BenchmarkIncOp2(b *testing.B) {
	op := NewIncOp(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	op.SetOutPort(func(data interface{}){ })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		op.InPort(0)
	}
}

func BenchmarkInc10Flow(b *testing.B) {
	var result int
	flow := NewInc10Flow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc100Flow(b *testing.B) {
	var result int
	flow := NewInc100Flow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc1KFlow(b *testing.B) {
	var result int
	flow := NewInc1KFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc10KFlow(b *testing.B) {
	var result int
	flow := NewInc10KFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc100KFlow(b *testing.B) {
	var result int
	flow := NewInc100KFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc1MFlow(b *testing.B) {
	var result int
	flow := NewInc1MFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

/*
func BenchmarkInc10MFlow(b *testing.B) {
	var result int
	flow := NewInc10MFlow(func(data interface{})int{ return data.(int) },
		func(data interface{}, i int)interface{}{ return i })
	flow.SetOutPort(func(data interface{}){ result = data.(int) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}
*/
