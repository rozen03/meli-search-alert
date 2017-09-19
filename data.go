package main

import "net/http"

type Respuesta struct {
	max       float64 `json:max`
	suggested float64 `json:suggested`
	min       float64 `json:min`
}
type httpInterface func(category string) (*http.Response, error)
type SearchGenerator func(text string) Search
type Search func(textt string) float64

/*
**Download the recieved search information
**returns a heap of prices and detailes of all prices in this search
 */
func getMenores(s Search, ch chan ArgsAndResult, get httpInterface) Respuesta {
	data := PreciosYVentas(s, category, ch, get)
	return MinMaxAvgData(data)
}
