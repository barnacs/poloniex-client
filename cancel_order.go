package poloniex

import (
	"strconv"
)

type successResponse struct {
	Success int `json:"success"`
}

func (r *successResponse) success() bool {
	return r.Success == 1
}

// Cancels an order you have placed in a given market.
func (c *Client) CancelOrder(orderNumber int) error {
	var resp successResponse
	args := map[string]string{
		"orderNumber": strconv.Itoa(orderNumber),
	}
	err := c.callTradingApi("cancelOrder", args, &resp)
	if err != nil {
		return err
	}
	if !resp.success() {
		return &PoloniexError{"Poloniex says order cancellation failed"}
	}
	return nil
}
