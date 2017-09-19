package main

import (
	"testing"
)

func Benchmark500Workers2Requests(b *testing.B) {
	workers := 500
	ch := startWorkers(workers)
	res1 := make(chan Respuesta)
	res2 := make(chan Respuesta)
	for i := 0; i < b.N; i++ {
		go func() {
			res1 <- Suggest(meliId, ch, GetMeli)
		}()
		go func() {
			res2 <- Suggest(meliId, ch, GetMeli)
		}()
		for c := 0; c < 2; c++ {
			select {
			case <-res1:
				c++
			case <-res2:
				c++
			}

		}
	}
}
func Benchmark500Workers5Requests(b *testing.B) {
	workers := 500
	ch := startWorkers(workers)
	res1 := make(chan Respuesta)
	res2 := make(chan Respuesta)
	res3 := make(chan Respuesta)
	res4 := make(chan Respuesta)
	res5 := make(chan Respuesta)
	for i := 0; i < b.N; i++ {
		go func() {res1 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res2 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res3 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res4 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res5 <- Suggest(meliId, ch, GetMeli)}()
		for c := 0; c < 5; c++ {
			select {
			case <-res1:
			case <-res2:
			case <-res1:
			case <-res3:
			case <-res4:
			case <-res5:
			}
			c++
		}
	}
}
func Benchmark500Workers10Requests(b *testing.B) {
	workers := 500
	ch := startWorkers(workers)
	res1 := make(chan Respuesta)
	res2 := make(chan Respuesta)
	res3 := make(chan Respuesta)
	res4 := make(chan Respuesta)
	res5 := make(chan Respuesta)
	res6 := make(chan Respuesta)
	res7 := make(chan Respuesta)
	res8 := make(chan Respuesta)
	res9 := make(chan Respuesta)
	res10 := make(chan Respuesta)
	for i := 0; i < b.N; i++ {
		go func() {res1 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res2 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res3 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res4 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res5 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res6 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res7 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res8 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res9 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res10 <- Suggest(meliId, ch,GetMeli)}()
		for c := 0; c < 10; c++ {
			select {
			case <-res1:
			case <-res2:
			case <-res3:
			case <-res4:
			case <-res5:
			case <-res6:
			case <-res7:
			case <-res8:
			case <-res9:
			case <-res10:
			}
			c++
		}
	}
}
func Benchmark500Workers20Requests(b *testing.B) {
	workers := 500
	ch := startWorkers(workers)
	res1 := make(chan Respuesta)
	res2 := make(chan Respuesta)
	res3 := make(chan Respuesta)
	res4 := make(chan Respuesta)
	res5 := make(chan Respuesta)
	res6 := make(chan Respuesta)
	res7 := make(chan Respuesta)
	res8 := make(chan Respuesta)
	res9 := make(chan Respuesta)
	res10 := make(chan Respuesta)
	res11 := make(chan Respuesta)
	res12 := make(chan Respuesta)
	res13 := make(chan Respuesta)
	res14 := make(chan Respuesta)
	res15 := make(chan Respuesta)
	res16 := make(chan Respuesta)
	res17 := make(chan Respuesta)
	res18 := make(chan Respuesta)
	res19 := make(chan Respuesta)
	res20 := make(chan Respuesta)
	for i := 0; i < b.N; i++ {
		go func() {res1 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res2 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res3 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res4 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res5 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res6 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res7 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res8 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res9 <- Suggest(meliId, ch, GetMeli)}()
		go func() {res10 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res11 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res12 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res13 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res14 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res15 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res16 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res17 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res18 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res19 <- Suggest(meliId, ch,GetMeli)}()
		go func() {res20 <- Suggest(meliId, ch,GetMeli)}()
		for c := 0; c < 20; c++ {
			select {
			case <-res1:
			case <-res2:
			case <-res3:
			case <-res4:
			case <-res5:
			case <-res6:
			case <-res7:
			case <-res8:
			case <-res9:
			case <-res10:
			case <-res11:
			case <-res12:
			case <-res13:
			case <-res14:
			case <-res15:
			case <-res16:
			case <-res17:
			case <-res18:
			case <-res19:
			case <-res20:
			}
			c++
		}
	}
}
