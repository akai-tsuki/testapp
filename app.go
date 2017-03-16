package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/server/:id", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {

	sum := 0
	for i := 0; i < 5; i++ {
		sum += i

	}

	id := c.Param("id")

	sums := strconv.Itoa(sum)

	return c.String(http.StatusOK, "Hello, World!"+id+sums)
}
