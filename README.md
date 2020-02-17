# Board trading system
This pkg is board trading system.  
※ Price is integer.

## Progress
■■□□□□□□□□□ 20%  

## TODO
- [ ] Database in program
- [ ] to Executions
- [ ] API (-> set -> match -> executions -> response)
- [ ] each Websocket channels
- [ ] all kinds...


- [x] Improve registration speed  
- [x] Improve reading speed
- [x] Improve aggregation speed
- [x] Order match & search & delete
- [x] UUID with xid
- [x] Frontend with Nuxt.js

## Usage
``` go
package main

func main() {
    board := New()


    // 注文の登録
    setPrice = 10
    o := orders.New()
    o.Price = <price>
    o.Size = <size>
    isMatch,executions := board.Set(o)
    if isMatch {
        // TODO:
        // 注文主へPrivate配信
        // Publicへ配信
        _ = executions
    }


    // セットした注文の集計と表示
    board.String()

    // to JSON
    depth = 10
    res := board.Snap(depth)
    JSON(http.StatusOK, res)
    // -> { "code": 200, "status": "success", "data": { "ltp": 150, "asks": [ { "price": 201, "size": 20 }, { "price": 150, "size": 20 } ], "bids": [ { "price": 150, "size": 20 }, { "price": 100, "size": 20 } ], "updated_at": "14:52:58" } }
}

```

## Frontend like a testnet
![frontend](https://github.com/go-numb/board-trading-system/blob/master/frontend/static/frontend.png?raw=true)

## Auther
[@_numbP](https://twitter.com/_numbP)