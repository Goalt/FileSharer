package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.Static("/", "html/index.html")

	_ = e.Start(":8080")
}
