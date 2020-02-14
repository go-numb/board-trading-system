package board

import (
	"time"

	"github.com/go-numb/board-trading-system/api/board/books"
)

type System struct {
	LTP int

	Ask *books.B
	Bid *books.B

	UpdatedAt time.Time
}

func New(ltp, length int) *System {
	return &System{
		LTP:       ltp,
		Ask:       books.New(true, ltp, length),
		Bid:       books.New(false, ltp, length),
		UpdatedAt: time.Now(),
	}
}
