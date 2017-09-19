package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

type obtainedData struct {
	min   float64
	max   float64
	sum   float64
	total float64
}

/*
**Download from MeLi with the arguments recieved.
**Returns map[string]interface{} parsing with the response from MeLi.
 */
func Download(args string, get httpInterface) (map[string]interface{}, error) {
	var resp *http.Response
	var err error
	var body map[string]interface{}
	//Download the json object from MeLi
	//Retry 100 times if error
	failed := 0
	for failed < 100 {
		// resp, err = http.Get(melink + args)
		resp, err = get(args)
		if err != nil {
			failed++
			continue
		}
		defer resp.Body.Close()
		//Decode response ioreader into map[string]interface{}
		//decode to any struct would likely cause errors
		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			failed++
			continue
		}

		break
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

/*
**Get the total count, sum, min and max of all prices.
 */
func GetPricesAndSold(items []interface{}) obtainedData {
	//Initialize return data
	prices := 0.0
	total := 0.0
	max := 0.0
	min := math.MaxFloat64
	/*
	**Loop over all items recieved
	**Sum the total and the prices
	**and evaluate the min and maximum values
	 */
	for i := range items {
		item := items[i].(map[string]interface{})
		price, ok := item["price"].(float64)
		if !ok {
			continue
		}
		sold, ok := item["sold_quantity"].(float64)
		if !ok {
			continue
		}
		max = math.Max(price, max)
		min = math.Min(price, min)
		prices += price * (sold + 1)
		total += (sold + 1)
	}
	return obtainedData{min, max, prices, total}
}

/*
**Download and returns min, max, total and prices
**From the recieved arguments
 */
func GetObtainedData(args string, c chan obtainedData, get httpInterface) {
	body, err := Download(args, get)
	if err != nil {
		c <- obtainedData{0.0, 0.0, 0.0, 0.0}
	}
	results, ok := body["results"].([]interface{})
	if !ok {
		c <- obtainedData{0.0, 0.0, 0.0, 0.0}
	}
	go func() {
		c <- GetPricesAndSold(results)
		if err := recover(); err != nil {
			c <- obtainedData{0.0, 0.0, 0.0, 0.0}
		}
	}()
}

/*
**Auxiliary method created for aesthetic code purposes
 */
func GetTotalCount(body *map[string]interface{}) int {
	return int((*body)["paging"].(map[string]interface{})["total"].(float64))
}

/*
**Send to the channel the arguments to download and process the items
**and also the channel where the task worker should send the prices
**processed and loop waiting for all prices while merging the information.
**Finally when all prices where processed
**returns the total count, the minimum, the maximum, the sum of all prices
 */
func PreciosYVentas(s Search, ch chan ArgsAndResult, get httpInterface) obtainedData {
	//Downloads first set of 200 items and total count
	body, err := Download(category, get)
	if err != nil {
		return obtainedData{0.0, 0.0, 0.0, 0.0}
	}
	results := body["results"].([]interface{})
	total := GetTotalCount(&body)
	res := GetPricesAndSold(results)
	chanels := (total / 200) - 1
	responses := make(chan obtainedData)
	//Start a Goroutine that would send in order all downloads waiting for any
	//Task worker free to download
	go func() {
		for c := 0; c < chanels; c++ {
			ch <- ArgsAndResult{responses, category + "&offset=" + strconv.Itoa(200*(c+1)), get}
		}
	}()
	/*
	**Wait for all channels to return and merge the prices information
	**then delete the channel key in the map to reduce the ammount of iterations
	 */
	done := 0
	for done < chanels {
		res_i := <-responses
		MergeObainedData(&res, &res_i)
		done++

	}
	return res
}
