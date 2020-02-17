package controllers

import (
	"net/http"

	"github.com/go-numb/board-trading-system/api/board"
	"github.com/go-numb/board-trading-system/api/board/books/orders"
	"github.com/go-numb/board-trading-system/api/models"

	"github.com/labstack/echo"
)

const (
	PREFIX = "REC"
	SUFFIX = "F"
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

// Order set & match in board
func (p *Client) Order(c echo.Context) error {
	o := orders.New()
	if err := c.Bind(o); err != nil {
		return err
	}

	// 要検討: UUIDとは異なるが親しい文字列になりそう
	o.AcceptanceID = models.CreateID(PREFIX, SUFFIX)

	executions := p.Board.Set(o)
	if 0 < len(executions) {
		// TODO:
		// 1. executionsを各発注者へ通知
		// 2. executionsをpublicに配信
		_ = executions

	}

	return c.JSON(http.StatusOK, &Response{
		Code:   200,
		Status: "success",
		Data:   o,
	})
}
