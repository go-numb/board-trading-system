package books

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-numb/board-trading-system/api/board/books/orders"
)

type B struct {
	Books []*Book
}

func New(isAsk bool, ltp, length int) *B {
	books := make([]*Book, length)
	price := ltp

	// 増減調整用
	adjust := 1
	if !isAsk {
		adjust = -1
	}

	for i := 0; i < length; i++ {
		books[i] = &Book{
			Price:     price + (i * adjust),
			Orders:    make([]orders.Order, 0),
			UpdatedAt: time.Now(),
		}
	}

	return &B{
		books,
	}
}

func (p *B) Find(price int) *Book {
	for i := range p.Books {
		if price != p.Books[i].Price {
			continue
		}
		return p.Books[i]
	}
	return nil
}

type Book struct {
	mux sync.RWMutex

	Price     int
	Orders    []orders.Order
	UpdatedAt time.Time
}

func (p *Book) String() string {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return fmt.Sprintf("%d - %d - %d", p.Price, len(p.Orders), p.Aggregate())
}

func (p *Book) Set(o *orders.Order) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.Orders = append(p.Orders, *o)
}

func (p *Book) Aggregate() (size int) {
	p.mux.RLock()
	defer p.mux.RUnlock()

	for i := range p.Orders {
		size += p.Orders[i].Size
	}
	return size
}
