package controllers

import (
	"fmt"
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
	o := orders.New()
	if err := c.Bind(o); err != nil {
		return err
	}

	if o.Side.IsAsk() {
		p.Board.Ask.Find(o.Price).Set(o)
		isMatch, executions := p.Board.Bid.Find(o.Price).Match(o)
		if isMatch {
			fmt.Printf("sell executions: %+v\n", executions)
		}
	} else {
		p.Board.Bid.Find(o.Price).Set(o)
		isMatch, executions := p.Board.Ask.Find(o.Price).Match(o)
		if isMatch {
			fmt.Printf("buy executions: %+v\n", executions)
		}
	}

	return c.JSON(http.StatusOK, &Response{
		Code:   200,
		Status: "success",
		Data:   o,
	})
}
