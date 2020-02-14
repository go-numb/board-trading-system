package board

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-numb/board-trading-system/api/board/books/orders"
)

func TestNew(t *testing.T) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	ltp := 10
	max := 10
	board := New(ltp, max)

	func() {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Println("set exec time: ", end.Sub(start))
		}()

		// 注文を擬似挿入
		count := 10000
		for i := 0; i < count; i++ {
			price := r.Intn(max*2-1) + 1

			if ltp <= price {
				board.Ask.Find(price).Set(&orders.Order{
					ID:    i,
					UUID:  fmt.Sprintf("UUID:%d", i),
					Price: price,
					Size:  (i + price) * -1,
				})
			} else {
				board.Bid.Find(price).Set(&orders.Order{
					ID:    i,
					UUID:  fmt.Sprintf("UUID:%d", i),
					Price: price,
					Size:  i + price,
				})
			}

			// fmt.Printf("%+v\n", board.Ask.Find(15).Orders)
		}
	}()

	func() {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Println("load exec time: ", end.Sub(start))
		}()

		for i := range board.Ask.Books {
			fmt.Printf("Ask:	%s\n", board.Ask.Books[len(board.Ask.Books)-i-1].String())
		}
		for i := range board.Bid.Books {
			fmt.Printf("		%s	:Bid\n", board.Bid.Books[i].String())
		}
	}()

}
