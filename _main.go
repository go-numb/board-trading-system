package main

import (
	"os"
	"path/filepath"

	"github.com/go-numb/board-system/api/controllers"
	"github.com/labstack/echo"
)

const PORT = ":8080"

var f *os.File

func init() {

}

func main() {
	e := echo.New()
	dir, _ := os.Getwd()
	c := controllers.New(filepath.Join(dir, "ldb"))

	e.GET(filepath.Join(API, VERSION, ":product_code", "orderboard"), c.Orderboard)
	e.POST(filepath.Join(API, VERSION, "orders"), c.Orders)
	e.Logger.Fatal(e.Start(PORT))
}
