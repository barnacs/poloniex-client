package poloniex

import (
	"encoding/json"
	"strconv"
	"time"
)

type period int

const (
	Period5Minutes  period = 300
	Period15Minutes period = 900
	Period30Minutes period = 1800
	Period2Hours    period = 7200
	Period4Hours    period = 14400
	Period24Hours   period = 86400
)

type ChartData struct {
	Date            time.Time
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Open            float64 `json:"open"`
	Close           float64 `json:"close"`
	Volume          float64 `json:"volume"`
	QuoteVolume     float64 `json:"quoteVolume"`
	WeightedAverage float64 `json:"weightedAverage"`
}

func (c *ChartData) UnmarshalJSON(b []byte) error {
	type chartData ChartData
	sChartData := &struct {
		Date int64 `json:"date"`
		*chartData
	}{
		chartData: (*chartData)(c),
	}

	if err := json.Unmarshal(b, sChartData); err != nil {
		return err
	}

	c.Date = time.Unix(sChartData.Date, 0)
	return nil
}

// Returns candlestick chart data.
func (c *Client) ChartData(market string, start, end time.Time, period period) ([]ChartData, error) {
	var chartData []ChartData
	args := map[string]string{
		"currencyPair": market,
		"start":        strconv.FormatInt(start.Unix(), 10),
		"end":          strconv.FormatInt(end.Unix(), 10),
		"period":       strconv.Itoa(int(period)),
	}
	err := c.callPublicApi("returnChartData", args, &chartData)
	return chartData, err
}
