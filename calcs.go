package main

import "math"

/*
**returns min, max and the average
 */
func MinMaxAvgData(data obtainedData) Respuesta {
	return Respuesta{data.max, data.sum / data.total, data.min}
}

/*
**get two references to obtainedData
**and saves in the first the min between it's max valures,
**the max between it's max valures
**the sum between the two total counts and prices
 */
func MergeObainedData(res *obtainedData, res_i *obtainedData) {
	(*res).sum += (*res_i).sum
	(*res).total += (*res_i).total
	(*res).max = math.Max((*res).max, (*res_i).max)
	(*res).min = math.Min((*res).min, (*res_i).min)
}
