package main

import "testing"

func BenchmarkIncOp(b *testing.B) {
	var result int
	flow := NewIncOp()
	(&flow).SetOutPort(func(i int){ result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}
