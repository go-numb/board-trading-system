package orders

import (
	"time"

	"github.com/go-numb/board-trading-system/api/models"
)

type Order struct {
	ID      int
	UUID    string
	Product string

	AcceptanceID string

	Price int
	// Size use buy+/sell-
	Size           int
	MinuteToExpire int

	Type      models.OrderType
	CreatedAt time.Time
}

func New() *Order {
	return &Order{
		CreatedAt: time.Now(),
	}
}
