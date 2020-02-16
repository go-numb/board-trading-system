package models

import "strings"

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
	UNDEFINED OrderSide = iota
	BUY
	SELL
)

func ToSide(in interface{}) OrderSide {
	switch v := in.(type) {
	case int:
		if v == 1 {
			return BUY
		} else if v == -1 {
			return SELL
		}

	case string:
		side := strings.ToLower(v)
		if side == "buy" {
			return BUY
		} else if side == "sell" {
			return SELL
		}

	}

	return UNDEFINED
}

func (side OrderSide) IsAsk() bool {
	return side == SELL
}
