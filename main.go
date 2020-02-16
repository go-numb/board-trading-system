package main

import (
	"github.com/go-numb/board-trading-system/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"path/filepath"
)

const (
	PORT    = ":8080"
	API     = "api"
	VERSION = "v1"
)

var f *os.File

func init() {

}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	// dir, _ := os.Getwd()
	// dbFilename := filepath.Join(dir, "ldb")
	c := controllers.New()

	e.GET(filepath.Join(API, VERSION, "board"), c.Orderboard)
	// e.GET(filepath.Join(API, VERSION, ":product_code", "orderboard"), c.Orderboard)
	e.POST(filepath.Join(API, VERSION, "private", "order"), c.Order)
	e.Logger.Fatal(e.Start(PORT))
}
