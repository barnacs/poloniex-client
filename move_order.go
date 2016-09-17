package poloniex

import (
	"strconv"
)

type MoveOrderResult struct {
	OrderNumber     int                         `json:"orderNumber,string"`
	ResultingTrades map[string][]ResultingTrade `json:"resultingTrades"`
	successResponse
}

// Cancels an order and places a new one of the same type in a single atomic transaction, meaning either both operations will succeed or both will fail.
func (c *Client) MoveOrder(orderNumber int, rate, amount float64) (MoveOrderResult, error) {
	var result MoveOrderResult
	args := buildArgMap(orderNumber, rate, amount)
	err := c.callTradingApi("moveOrder", args, &result)
	if err != nil {
		return result, err
	}
	if !result.success() {
		return result, &PoloniexError{"Poloniex says moving the order failed"}
	}
	return result, err
}

func buildArgMap(orderNumber int, rate, amount float64) map[string]string {
	args := map[string]string{
		"orderNumber": strconv.Itoa(orderNumber),
		"rate":        strconv.FormatFloat(rate, 'f', -1, 64),
	}
	if amount > 0 {
		args["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)
	}
	return args
}
