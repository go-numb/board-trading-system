package orders

import (
	"time"

	"github.com/go-numb/board-trading-system/api/models"
)

type Orders []Order

func (a Orders) Len() int { return len(a) }
func (a Orders) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Orders) Less(i, j int) bool { return a[i].CreatedAt.UnixNano() < a[j].CreatedAt.UnixNano() }

type Order struct {
	ID      int
	UUID    string
	Product string

	AcceptanceID string

	Side  models.OrderSide
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
