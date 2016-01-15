package main

import (
	"testing"
)

type Input interface {
	InPort(i int)
}
type Output interface {
	SetOutPort(func(i int))
}

func TestInterfaces(t *testing.T) {
	var result int
	op := NewIncOp()

	flow := NewInc10Flow()
	flow.SetOutPort(func(i int) { result = i })

	var value interface{}
	value = op
	switch val := value.(type) {
	case Output:
		val.SetOutPort(func(i int) { result = i })
	default:
		t.Errorf("Unable to convert operation (type %T) to interface 'Output'.\n", value)
	}
	switch val := value.(type) {
	case Input:
		val.InPort(0)
	default:
		t.Errorf("Unable to convert operation (type %T) to interface 'Input'.\n", value)
	}

	value = flow
	switch value.(type) {
	case Output:
		t.Errorf("Able to convert flow (type %T) to interface 'Output'.\n", value)
	default:
		t.Logf("The flow (type %T) doesn't implement the interface 'Output' because the functions belong to the struct itself and not a pointer to it.\n", value)
	}
	switch value.(type) {
	case Input:
		t.Errorf("Able to convert flow (type %T) to interface 'Input'.\n", value)
	default:
		t.Logf("The flow (type %T) doesn't implement the interface 'Input' because the functions belong to the struct itself and not a pointer to it.\n", value)
	}

	value = *flow
	switch value.(type) {
	case Output:
		t.Errorf("Able to convert flow (type %T) to interface 'Output'.\n", value)
	default:
		t.Logf("The flow (type %T) still doesn't implement the interface 'Output' because the functions are fields and not methods. :-(\n", value)
	}
	switch value.(type) {
	case Input:
		t.Errorf("Able to convert flow (type %T) to interface 'Input'.\n", value)
	default:
		t.Logf("The flow (type %T) still doesn't implement the interface 'Input' because the functions are fields and not methods. :-(\n", value)
	}
}

func BenchmarkInc10Flow(b *testing.B) {
	var result int
	flow := NewInc10Flow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc100Flow(b *testing.B) {
	var result int
	flow := NewInc100Flow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc1KFlow(b *testing.B) {
	var result int
	flow := NewInc1KFlow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc10KFlow(b *testing.B) {
	var result int
	flow := NewInc10KFlow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc100KFlow(b *testing.B) {
	var result int
	flow := NewInc100KFlow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc1MFlow(b *testing.B) {
	var result int
	flow := NewInc1MFlow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}

func BenchmarkInc10MFlow(b *testing.B) {
	var result int
	flow := NewInc10MFlow()
	flow.SetOutPort(func(i int) { result = i })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flow.InPort(0)
	}
}
