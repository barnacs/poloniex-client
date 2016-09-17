package poloniex

type Ticker struct {
	Last          float64 `json:"last,string"`
	LowestAsk     float64 `json:"lowestAsk,string"`
	HighestBid    float64 `json:"highestBid,string"`
	PercentChange float64 `json:"percentChange,string"`
	BaseVolume    float64 `json:"baseVolume,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
}

// Returns the ticker for all markets.
func (c *Client) Tickers() (map[string]Ticker, error) {
	var tickers map[string]Ticker
	err := c.callPublicApi("returnTicker", nil, &tickers)
	return tickers, err
}

// Returns the ticker for a specific market.
func (c *Client) Ticker(market string) (Ticker, error) {
	tickers, err := c.Tickers()
	if err != nil {
		return Ticker{}, err
	}
	ticker, ok := tickers[market]
	if !ok {
		return ticker, &PoloniexError{"Unknown market"}
	}
	return ticker, nil
}
