package poloniex

type Balance struct {
	Available float64 `json:"available,string"`
	OnOrders  float64 `json:"onOrders,string"`
	BtcValue  float64 `json:"btcValue,string"`
}

// Returns all of your balances, including available balance, balance on orders, and the estimated BTC value of your balance.
// This call is limited to your exchange account.
func (c *Client) Balances() (map[string]Balance, error) {
	var balances map[string]Balance
	err := c.callTradingApi("returnCompleteBalances", nil, &balances)
	return balances, err
}

// Returns exchange account balances for a given currency.
func (c *Client) Balance(currency string) (Balance, error) {
	balances, err := c.Balances()
	if err != nil {
		return Balance{}, err
	}
	balance, ok := balances[currency]
	if !ok {
		return balance, &PoloniexError{"Unknown currency"}
	}
	return balance, err
}
