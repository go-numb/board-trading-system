package models

type OrderType int

const (
	MARKET OrderType = iota
	LIMIT
	IFD
	OCO
	IFDOCO
)

type OrderSide int

const (
	BUY OrderSide = iota
	SELL
)
