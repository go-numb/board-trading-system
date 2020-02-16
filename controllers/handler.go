package controllers

import (
	"net/http"

	"github.com/go-numb/board-trading-system/api/board"
	"github.com/go-numb/board-trading-system/api/board/books/orders"

	"github.com/labstack/echo"
)

type Client struct {
	Board *board.System
}

func New() *Client {
	return &Client{
		Board: board.New(),
	}
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func (p *Client) Orderboard(c echo.Context) error {
	p.Board.LTP = 150
	return c.JSON(http.StatusOK, &Response{
		Code:   200,
		Status: "success",
		Data:   p.Board.Snap(0),
	})
}

func (p *Client) Order(c echo.Context) error {
	o := new(orders.Order)
	if err := c.Bind(o); err != nil {
		return err
	}

	if o.Side.IsAsk() {
		p.Board.Ask.Find(o.Price).Set(o)
	} else {
		p.Board.Bid.Find(o.Price).Set(o)
	}

	return c.JSON(http.StatusOK, &Response{
		Code:   200,
		Status: "success",
		Data:   o,
	})
}
