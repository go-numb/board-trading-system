package board

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	v1 "github.com/go-numb/go-bitflyer/v1"
	"github.com/go-numb/go-bitflyer/v1/jsonrpc"

	"github.com/go-numb/board-trading-system/api/board/books/orders"
	"github.com/go-numb/board-trading-system/api/models"
)

func TestNew(t *testing.T) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	ltp := 10
	max := 20
	board := New()

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
					ID:        i,
					UUID:      fmt.Sprintf("UUID:%d", i),
					Price:     price,
					Size:      (i + price) * -1,
					CreatedAt: time.Now().Add(time.Duration(i) + time.Minute),
				})
			} else {
				board.Bid.Find(price).Set(&orders.Order{
					ID:        i,
					UUID:      fmt.Sprintf("UUID:%d", i),
					Price:     price,
					Size:      i + price,
					CreatedAt: time.Now().Add(time.Duration(i) + time.Minute),
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

	for i := 0; i < max; i++ {
		isMatch, executions := board.Bid.Find(i).Match(&orders.Order{
			Size: 100,
		})
		if isMatch {
			fmt.Printf("is match: %+v\n", executions)
			break
		}
	}
}

func TestDo(t *testing.T) {
	board := New()

	ch := make(chan jsonrpc.Response)

	product := "FX_BTC_JPY"
	channels := []string{
		"lightning_board_%s",
		"lightning_executions_%s",
	}
	for i := range channels {
		channels[i] = fmt.Sprintf(channels[i], product)
	}
	go jsonrpc.Get(channels, ch)

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			board.String(10)
			// sort.Sort(board.Ask)
			// temp := board.Ask.Books[:12]
			// for i := range temp {
			// 	fmt.Printf("Ask:	%s\n", temp[len(temp)-1-i].String())
			// 	if 10 < i {
			// 		break
			// 	}
			// }
			// fmt.Printf("LTP: %d\n", board.LTP)
			// sort.Sort(board.Bid)
			// for i := range board.Bid.Books {
			// 	fmt.Printf("		%s	:Bid\n", board.Bid.Books[i].String())
			// 	if 10 < i {
			// 		break
			// 	}
			// }

		case v := <-ch:
			switch v.Type {
			case jsonrpc.Board:
				for i := range v.Orderbook.Asks {
					board.Ask.Find(int(v.Orderbook.Asks[i].Price)).Set(
						&orders.Order{
							Side:      models.SELL,
							Price:     int(v.Orderbook.Asks[i].Price),
							Size:      int(v.Orderbook.Asks[i].Size * 100000000000),
							CreatedAt: time.Now(),
						},
					)
				}
				for i := range v.Orderbook.Bids {
					board.Bid.Find(int(v.Orderbook.Bids[i].Price)).Set(
						&orders.Order{
							Side:      models.BUY,
							Price:     int(v.Orderbook.Bids[i].Price),
							Size:      int(v.Orderbook.Bids[i].Size * 100000000000),
							CreatedAt: time.Now(),
						},
					)
				}
				// fmt.Printf("board: %+v\n", v.Orderbook)
			case jsonrpc.Executions:
				board.LTP = int(v.Executions[0].Price)
				for i := range v.Executions {
					if v.Executions[i].Side == v1.BUY {
						board.Ask.Find(int(v.Executions[i].Price)).Match(&orders.Order{
							UUID:  v.Executions[i].BuyChildOrderAcceptanceID,
							Side:  models.ToSide(v.Executions[i].Side),
							Price: int(v.Executions[i].Price),
							Size:  int(v.Executions[i].Size * 100000000000),
						})
					} else {
						board.Ask.Find(int(v.Executions[i].Price)).Match(&orders.Order{
							UUID:  v.Executions[i].BuyChildOrderAcceptanceID,
							Side:  models.ToSide(v.Executions[i].Side),
							Price: int(v.Executions[i].Price),
							Size:  int(v.Executions[i].Size * 100000000000),
						})
					}
				}
				// fmt.Printf("executions: %+v\n", v.Executions)
			}
		}
	}
}
