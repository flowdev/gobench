package main

import "testing"

func BenchmarkInc10Flow(b *testing.B) {
	var result int
	flow := NewInc10Flow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc100Flow(b *testing.B) {
	var result int
	flow := NewInc100Flow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc1KFlow(b *testing.B) {
	var result int
	flow := NewInc1KFlow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc10KFlow(b *testing.B) {
	var result int
	flow := NewInc10KFlow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc100KFlow(b *testing.B) {
	var result int
	flow := NewInc100KFlow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc1MFlow(b *testing.B) {
	var result int
	flow := NewInc1MFlow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc10MFlow(b *testing.B) {
	var result int
	flow := NewInc10MFlow()
	flow.SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}
