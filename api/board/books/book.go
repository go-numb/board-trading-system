package books

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/go-numb/board-trading-system/api/board/books/orders"
)

type B struct {
	Books []*Book
}

func (a B) Len() int { return len(a.Books) }
func (a B) Swap(i, j int) {
	a.Books[i], a.Books[j] = a.Books[j], a.Books[i]
}
func (a B) Less(i, j int) bool { return a.Books[i].Price < a.Books[j].Price }

func New() *B {
	books := make([]*Book, 0)

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
	p.Books = append(p.Books, NewBook(price))
	return p.Books[len(p.Books)-1]
}

type Book struct {
	mux sync.RWMutex

	Price     int
	Orders    []orders.Order
	UpdatedAt time.Time
}

func NewBook(price int) *Book {
	return &Book{
		Price:     price,
		Orders:    make([]orders.Order, 0),
		UpdatedAt: time.Now(),
	}
}

func (p *Book) String() string {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return fmt.Sprintf("%d	-	%d	-	%d	%s", p.Price, len(p.Orders), p.Aggregate(), p.UpdatedAt.Format("15:04:05"))
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

// Match is order matcher
// - 注文が合致すれば注文分板を削除し、Executions化
// - 注文残があれば、引数oを次の価格へ持ち越す
// - 約定履歴は喰う側である成り方向のSideになる
func (p *Book) Match(o *orders.Order) (isMatch bool, executions []orders.Order) {
	if p == nil {
		return false, executions
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	orderArray := orders.Orders(p.Orders)
	sort.Reverse(orderArray)

	// 古いものが先頭にあるため、新規注文分を削除していく
	remain := o.Size
	for i := range orderArray {
		diff := orderArray[i].Size - remain
		if diff <= 0 { // 待機注文完全約定 & 新規注文残
			executions = append(executions, orderArray[i])
			if diff == 0 { // 待機注文と新規注文が合致
				break
			}
			// 約定した分、注文枚数を減らし、次へ
			remain -= orderArray[i].Size
			continue
		}

		// 部分約定
		// 待機注文一部約定 & 新規完全約定
		ord := orderArray[i]
		ord.Size = remain
		executions = append(executions, ord)
		// 約定残はorderArray[i]に戻す
		orderArray[i].Size -= remain
		// 新規注文がなくなった
		break
	}

	// 待機注文を満たしたものを削除する
	for i := range executions {
		for j := range orderArray {
			if orderArray[j].UUID != executions[i].UUID {
				continue
			}
			orderArray = append(orderArray[:j], orderArray[j+1:]...)
			break
		}
	}

	return true, executions
}
