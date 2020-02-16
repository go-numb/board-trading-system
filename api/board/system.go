package board

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/go-numb/board-trading-system/api/board/books"
)

type System struct {
	LTP int `json:"ltp,omitempty"`

	Ask *books.B `json:"ask,omitempty"`
	Bid *books.B `json:"bid,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func New() *System {
	return &System{
		Ask:       books.New(),
		Bid:       books.New(),
		UpdatedAt: time.Now(),
	}
}

func (p *System) String(depth int) {
	sort.Sort(p.Ask)
	var str []string
	count := 0
	for i := range p.Ask.Books {
		if p.Ask.Books[i].Price >= p.LTP &&
			len(p.Ask.Books[i].Orders) != 0 {
			str = append(str, fmt.Sprintf("%s", p.Ask.Books[i].String()))
			count++
		}
		if depth < count {
			break
		}
	}

	for i := range str {
		fmt.Printf("%+v\n", str[len(str)-1-i])
	}

	sort.Sort(p.Bid)
	var spread int
	if len(p.Bid.Books) != 0 && len(p.Ask.Books) != 0 {
		spread = int(math.Max(0, float64(p.Ask.Books[0].Price-p.Bid.Books[len(p.Bid.Books)-1].Price)))
	}
	fmt.Printf("------------	%d	%d\n", p.LTP, spread)

	count = 0
	for i := range p.Bid.Books {
		l := len(p.Bid.Books) - 1 - i
		if p.Bid.Books[l].Price <= p.LTP &&
			len(p.Bid.Books[l].Orders) != 0 {
			fmt.Printf("			%s\n", p.Bid.Books[l].String())
			count++
		}
		if depth < count {
			break
		}
	}
}

type Response struct {
	LTP       int       `json:"ltp,omitempty"`
	Spread    int       `json:"spread,omitempty"`
	Asks      []Inner   `json:"asks,omitempty"`
	Bids      []Inner   `json:"bids,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Inner struct {
	Price int `json:"price,omitempty"`
	Size  int `json:"size,omitempty"`
}

func NewResponse() *Response {
	return &Response{
		Asks: make([]Inner, 0),
		Bids: make([]Inner, 0),
	}
}

func (p *System) Snap(depth int) *Response {
	if depth == 0 {
		depth = 100
	}
	res := NewResponse()

	sort.Sort(p.Ask)
	var books []Inner
	count := 0
	for i := range p.Ask.Books {
		if p.Ask.Books[i].Price >= p.LTP &&
			len(p.Ask.Books[i].Orders) != 0 {
			books = append(books, Inner{
				Price: p.Ask.Books[i].Price,
				Size:  p.Ask.Books[i].Aggregate(),
			})
			count++
		}
		if depth < count {
			break
		}
	}

	for i := range books {
		res.Asks = append(res.Asks, books[len(books)-1-i])
	}

	sort.Sort(p.Bid)
	if len(p.Bid.Books) != 0 && len(p.Ask.Books) != 0 {
		res.Spread = int(math.Max(0, float64(p.Ask.Books[0].Price-p.Bid.Books[len(p.Bid.Books)-1].Price)))
	}
	res.LTP = p.LTP
	res.UpdatedAt = p.UpdatedAt

	count = 0
	for i := range p.Bid.Books {
		l := len(p.Bid.Books) - 1 - i
		if p.Bid.Books[l].Price <= p.LTP &&
			len(p.Bid.Books[l].Orders) != 0 {
			res.Bids = append(res.Bids, Inner{
				Price: p.Bid.Books[l].Price,
				Size:  p.Bid.Books[l].Aggregate(),
			})
			count++
		}
		if depth < count {
			break
		}
	}

	return res
}
