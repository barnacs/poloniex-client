package poloniex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	publicUrl  = "https://poloniex.com/public"
	tradingUrl = "https://poloniex.com/tradingApi"
)

type PoloniexError struct {
	Message string `json:"error"`
}

func (e *PoloniexError) Error() string {
	return e.Message
}

type Client struct {
	key     string
	secret  string
	client  *http.Client
	reqLock sync.Mutex
}

func NewClient(key, secret string) *Client {
	return &Client{
		key:    key,
		secret: secret,
		client: &http.Client{},
	}
}

func (c *Client) callPublicApi(cmd string, args map[string]string, resp interface{}) error {
	req, err := http.NewRequest("GET", publicUrl, nil)
	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Set("command", cmd)
	for key, value := range args {
		query.Set(key, value)
	}

	req.URL.RawQuery = query.Encode()
	r, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer r.Body.Close()
	return parseResponse(r.Body, resp)
}

func (c *Client) callTradingApi(cmd string, args map[string]string, resp interface{}) error {
	form := url.Values{}
	form.Add("command", cmd)
	for key, value := range args {
		form.Add(key, value)
	}

	c.reqLock.Lock()

	form.Add("nonce", strconv.FormatInt(time.Now().UnixNano(), 10))

	body := form.Encode()
	req, err := http.NewRequest("POST", tradingUrl, strings.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Key", c.key)

	mac := hmac.New(sha512.New, []byte(c.secret))
	mac.Write([]byte(body))
	signature := hex.EncodeToString(mac.Sum(nil))

	req.Header.Add("Sign", signature)

	r, err := c.client.Do(req)
	c.reqLock.Unlock()

	if err != nil {
		return err
	}

	defer r.Body.Close()
	return parseResponse(r.Body, resp)
}

func parseResponse(resp io.Reader, target interface{}) error {
	body, err := ioutil.ReadAll(resp)
	if err != nil {
		return err
	}
	if err = tryParseError(body); err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func tryParseError(body []byte) error {
	var parsedError PoloniexError
	json.Unmarshal(body, &parsedError)
	if len(parsedError.Message) > 0 {
		return &parsedError
	}
	return nil
}
