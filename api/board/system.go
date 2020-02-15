package board

import (
	"fmt"
	"sort"
	"time"

	"github.com/go-numb/board-trading-system/api/board/books"
)

type System struct {
	LTP int

	Ask *books.B
	Bid *books.B

	UpdatedAt time.Time
}

func New() *System {
	return &System{
		Ask:       books.New(),
		Bid:       books.New(),
		UpdatedAt: time.Now(),
	}
}

func (p *System) String(depth int) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("board print exec time: ", end.Sub(start))
	}()

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

	fmt.Printf("------------	%d\n", p.LTP)

	sort.Sort(p.Bid)
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
