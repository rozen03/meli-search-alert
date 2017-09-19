package main

import (
	"fmt"
	"testing"
)

const meliId = "MLA5726"

func Test001(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func Benchmark10Workers(b *testing.B) {
	ch := startWorkers(10)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}
}
func Benchmark20Workers(b *testing.B) {
	ch := startWorkers(20)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}
}
func Benchmark30Workers(b *testing.B) {
	ch := startWorkers(30)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}
}
func Benchmark40Workers(b *testing.B) {
	workers := 40
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}
}
func Benchmark50Workers(b *testing.B) {
	workers := 50
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark60Workers(b *testing.B) {
	workers := 60
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark70Workers(b *testing.B) {
	workers := 70
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark100Workers(b *testing.B) {
	workers := 100
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark200Workers(b *testing.B) {
	workers := 200
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark500Workers(b *testing.B) {
	workers := 500
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark1000Workers(b *testing.B) {
	workers := 1000
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark2000Workers(b *testing.B) {
	workers := 2000
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}

}
func Benchmark4000Workers(b *testing.B) {
	workers := 4000
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}
}
func Benchmark8000Workers(b *testing.B) {
	workers := 8000
	ch := startWorkers(workers)
	for i := 0; i < b.N; i++ {
		Suggest(meliId, ch, GetMeli)
	}
}
