# Board trading system
This pkg is board trading system.  
※ Price is integer.

## Progress
■□□□□□□□□□□ 10%  

## TODO
- [ ] Order match & search & delete
- [ ] to Executions
- [ ] API (-> set -> match -> executions -> response)
- [ ] Frontend with Nuxt.js
- [ ] all kinds...


- [x] Improve registration speed  
- [x] Improve reading speed
- [x] Improve aggregation speed

# Usage
``` go
package main

func main() {
	centerPrice := 10
	maxLength := 10
	board := New(ltp, maxLength)


    // 注文の登録
    setPrice = 10
    board.Ask.Find(setPrice).Set(&orders.Order{
        ID:    i,
        UUID:  fmt.Sprintf("UUID:%d", i),
        Price: setPrice,
        Size:  <size>,
    })


    // セットした注文の集計と表示
    for i := range board.Ask.Books {
        fmt.Printf("Ask:	%s\n", board.Ask.Books[len(board.Ask.Books)-i-1].String())
    }
    for i := range board.Bid.Books {
        fmt.Printf("		%s	:Bid\n", board.Bid.Books[i].String())
    }
}

```

## Auther
[@_numbP](https://twitter.com/_numbP)