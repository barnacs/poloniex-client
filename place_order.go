package poloniex

import (
	"strconv"
)

type OrderResult struct {
	OrderNumber     int `json:"orderNumber,string"`
	ResultingTrades []ResultingTrade
}

type ResultingTrade struct {
	TradeId int `json:"tradeID,string"`
	Trade
}

type OrderRequest struct {
	Market string
	Rate   float64
	Amount float64
	// A fill-or-kill order will either fill in its entirety or be completely aborted.
	FillOrKill bool
	// An immediate-or-cancel order can be partially or completely filled, but any portion of the order that cannot be filled immediately will be canceled rather than left on the order book.
	ImmediateOrCancel bool
	// A post-only order will only be placed if no portion of it fills immediately; this guarantees you will never pay the taker fee on any part of the order that fills.
	PostOnly bool
}

func (o OrderRequest) asArgMap() map[string]string {
	args := map[string]string{
		"currencyPair": o.Market,
		"rate":         strconv.FormatFloat(o.Rate, 'f', -1, 64),
		"amount":       strconv.FormatFloat(o.Amount, 'f', -1, 64),
	}
	if o.FillOrKill {
		args["fillOrKill"] = "1"
	}
	if o.ImmediateOrCancel {
		args["immediateOrCancel"] = "1"
	}
	if o.PostOnly {
		args["postOnly"] = "1"
	}
	return args
}

// Places a limit buy order in a given market.
func (c *Client) Buy(order OrderRequest) (OrderResult, error) {
	var result OrderResult
	args := order.asArgMap()
	err := c.callTradingApi("buy", args, &result)
	return result, err
}

// Places a sell order in a given market.
func (c *Client) Sell(order OrderRequest) (OrderResult, error) {
	var result OrderResult
	args := order.asArgMap()
	err := c.callTradingApi("sell", args, &result)
	return result, err
}
