package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	header := c.Request().Header
	// fmt.Println(header.Get("user-agent"))
	headers := fmt.Sprintf("%#v", header)

	output := strings.Replace(headers, ", ", ", \n", -1)

	return c.String(http.StatusOK, "Hello, World! id: "+id+", count: "+cntStr+"\n"+output)

	//	sum := 0
	//	for i := 0; i < 5; i++ {
	//		sum += i
	//
	//	}

	//	sums := strconv.Itoa(sum)

	//	return c.String(http.StatusOK, "Hello, World!"+id+sums)
}
