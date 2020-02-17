package orders

import (
	"time"

	"github.com/go-numb/board-trading-system/api/models"
)

const PREFIX = "BOS-"

type Orders []Order

func (a Orders) Len() int { return len(a) }
func (a Orders) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Orders) Less(i, j int) bool { return a[i].CreatedAt.UnixNano() < a[j].CreatedAt.UnixNano() }

type Order struct {
	ID      int64  `json:"id"`
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
	now := time.Now()
	return &Order{
		ID:        now.UnixNano() / 1000,
		UUID:      models.CreateID(PREFIX, ""),
		CreatedAt: now,
	}
}

// Next is remain size, go to next price
func (p *Order) Next() bool {
	if 0 < p.Size {
		return true
	}
	return false
}
