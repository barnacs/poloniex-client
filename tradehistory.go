package poloniex

import (
	"encoding/json"
	"strconv"
	"time"
)

type tradeType string

const (
	TradeTypeBuy  tradeType = "buy"
	TradeTypeSell tradeType = "sell"
)

type Trade struct {
	Date   time.Time
	Type   tradeType `json:"type"`
	Rate   float64   `json:"rate,string"`
	Amount float64   `json:"amount,string"`
	Total  float64   `json:"total,string"`
}

func (t *Trade) UnmarshalJSON(b []byte) (err error) {
	type trade Trade
	sTrade := &struct {
		Date string `json:"date"`
		*trade
	}{
		trade: (*trade)(t),
	}

	if err := json.Unmarshal(b, sTrade); err != nil {
		return err
	}

	t.Date, err = time.Parse("2006-01-02 15:04:05", sTrade.Date)
	return err
}

// Returns the past 200 trades for a given market.
func (c *Client) TradeHistory(market string) ([]Trade, error) {
	var trades []Trade
	args := map[string]string{
		"currencyPair": market,
	}
	err := c.callPublicApi("returnTradeHistory", args, &trades)
	return trades, err
}

// Returns up to 50,000 trades between a range specified.
func (c *Client) TradeHistoryBetween(market string, start, end time.Time) ([]Trade, error) {
	var trades []Trade
	args := map[string]string{
		"currencyPair": market,
		"start":        strconv.FormatInt(start.Unix(), 10),
		"end":          strconv.FormatInt(end.Unix(), 10),
	}
	err := c.callPublicApi("returnTradeHistory", args, &trades)
	return trades, err
}
