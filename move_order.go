package poloniex

import (
	"strconv"
)

type MoveOrderRequest struct {
	OrderNumber int
	Rate        float64
	// You may optionally specify "amount" if you wish to change the amount of the new order.
	Amount float64
	// An immediate-or-cancel order can be partially or completely filled, but any portion of the order that cannot be filled immediately will be canceled rather than left on the order book. May be specified for exchange orders, but will have no effect on margin orders.
	ImmediateOrCancel bool
	// A post-only order will only be placed if no portion of it fills immediately; this guarantees you will never pay the taker fee on any part of the order that fills. May be specified for exchange orders, but will have no effect on margin orders.
	PostOnly bool
}

func (r MoveOrderRequest) asArgMap() map[string]string {
	args := map[string]string{
		"orderNumber": strconv.Itoa(r.OrderNumber),
		"rate":        strconv.FormatFloat(r.Rate, 'f', -1, 64),
	}
	if r.Amount > 0 {
		args["amount"] = strconv.FormatFloat(r.Amount, 'f', -1, 64)
	}
	if r.ImmediateOrCancel {
		args["immediateOrCancel"] = "1"
	}
	if r.PostOnly {
		args["postOnly"] = "1"
	}
	return args
}

type MoveOrderResult struct {
	OrderNumber     int                         `json:"orderNumber,string"`
	ResultingTrades map[string][]ResultingTrade `json:"resultingTrades"`
	successResponse
}

// Cancels an order and places a new one of the same type in a single atomic transaction, meaning either both operations will succeed or both will fail.
func (c *Client) MoveOrder(req MoveOrderRequest) (MoveOrderResult, error) {
	var result MoveOrderResult
	args := req.asArgMap()
	err := c.callTradingApi("moveOrder", args, &result)
	if err != nil {
		return result, err
	}
	if !result.success() {
		return result, &PoloniexError{"Poloniex says moving the order failed"}
	}
	return result, err
}
