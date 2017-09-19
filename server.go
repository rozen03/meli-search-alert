package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const melink = "https://api.mercadolibre.com/sites/MLA/search?limit=200&category="

// func ping(c *gin.Context) {
// c.JSON(200, gin.H{
// "message": "pong",
// })
// }
func prices(c *gin.Context, ch chan ArgsAndResult) {
	res := Suggest(c.Param("id"), ch, GetMeli)
	c.JSON(200, gin.H{
		"max":       strconv.FormatFloat(res.max, 'f', 2, 64),
		"suggested": strconv.FormatFloat(res.suggested, 'f', 2, 64),
		"min":       strconv.FormatFloat(res.min, 'f', 2, 64),
	})

}
func menores(c *gin.Context, ch chan ArgsAndResult) {
	res := getMenores(ch, GetMeli)
}

func GetMeli(args string) (*http.Response, error) { return http.Get(melink + args) }

func start() {
	var ch chan ArgsAndResult
	if len(os.Args) > 0 {
		fmt.Println(os.Args[1])
		workers, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		ch = startWorkers(workers)
	} else {
		ch = startWorkers(maxChanelsSched)
	}
	myfile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(myfile, os.Stdout)

	r := gin.Default()
	// r.GET("/ping", ping)
	// r.GET("/categories/:id/prices", func(c *gin.Context) { prices(c, ch) })
	r.GET("/menores", func(c *gin.Context) { menores(c, ch) })
	r.Run(":8081")

}
