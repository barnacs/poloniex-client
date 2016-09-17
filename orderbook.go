package poloniex

import (
	"encoding/json"
	"strconv"
)

type OrderBook struct {
	Asks []Order
	Bids []Order
}

type Order struct {
	Price  float64
	Amount float64
}

func (o *Order) UnmarshalJSON(b []byte) error {
	var order [2]interface{}
	err := json.Unmarshal(b, &order)
	if err != nil {
		return err
	}
	o.Price, err = strconv.ParseFloat(order[0].(string), 64)
	o.Amount = order[1].(float64)
	return err
}

// Returns the order book for all markets.
func (c *Client) OrderBooks(depth int) (map[string]OrderBook, error) {
	var orderBooks map[string]OrderBook
	args := map[string]string{
		"currencyPair": "all",
		"depth":        strconv.Itoa(depth),
	}
	err := c.callPublicApi("returnOrderBook", args, &orderBooks)
	return orderBooks, err
}

// Returns the order book for a given market.
func (c *Client) OrderBook(market string, depth int) (OrderBook, error) {
	var orderBook OrderBook
	args := map[string]string{
		"currencyPair": market,
		"depth":        strconv.Itoa(depth),
	}
	err := c.callPublicApi("returnOrderBook", args, &orderBook)
	return orderBook, err
}
