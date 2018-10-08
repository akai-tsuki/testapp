package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var count int

func main() {
	e := echo.New()
	e.GET("/server/:id", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {

	id := c.Param("id")

	count++
	cntStr := strconv.Itoa(count)

	return c.String(http.StatusOK, "Hello, World! id: "+id+", count: "+cntStr)

	//	sum := 0
	//	for i := 0; i < 5; i++ {
	//		sum += i
	//
	//	}

	//	sums := strconv.Itoa(sum)

	//	return c.String(http.StatusOK, "Hello, World!"+id+sums)
}
