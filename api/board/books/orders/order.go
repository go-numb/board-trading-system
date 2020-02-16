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
	ID      int    `json:"id"`
	UUID    string `json:"uuid"`
	Product string `json:"product"`

	AcceptanceID string `json:"acceptance_id,omitempty"`

	Side  models.OrderSide `json:"side,omitempty"`
	Price int              `json:"price,omitempty"`
	// Size use buy+/sell-
	Size           int `json:"size,omitempty"`
	MinuteToExpire int `json:"minute_to_expire,omitempty"`

	Type      models.OrderType `json:"type,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
}

func New() *Order {
	return &Order{
		CreatedAt: time.Now(),
	}
}
