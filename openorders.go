package poloniex

type OpenOrder struct {
	OrderNumber int       `json:"orderNumber,string"`
	Type        tradeType `json:"type"`
	Rate        float64   `json:"rate,string"`
	Amount      float64   `json:"amount,string"`
	Total       float64   `json:"total,string"`
}

// Returns your open orders for all markets.
func (c *Client) OpenOrders() (map[string][]OpenOrder, error) {
	var orders map[string][]OpenOrder
	args := map[string]string{
		"currencyPair": "all",
	}
	err := c.callTradingApi("returnOpenOrders", args, &orders)
	return orders, err
}

// Returns your open orders for a given market.
func (c *Client) OpenOrdersFor(market string) ([]OpenOrder, error) {
	var orders []OpenOrder
	args := map[string]string{
		"currencyPair": market,
	}
	err := c.callTradingApi("returnOpenOrders", args, &orders)
	return orders, err
}
