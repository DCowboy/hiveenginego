package hiveenginego

import (
	//~ "fmt"
	"strings"
	"encoding/json"
)

type OrderBook struct {
	BookType                   string
	Book                       []Order       `json:""`
}

type Order struct {
	_id                        int           `json:"_id"`
	TxId                       string        `json:"txId"`
	Timestamp                  int           `json:"timestamp"`
	Account                    string        `json:"account"`
	Symbol                     string        `json:"symbol"`
	Quantity                   float32       `json:",string"`
	Price                      float32       `json:",string"`
	PriceDec                   interface{}   `json:"priceDec"`
	TokensLocked               float32       `json:",string, omitempty"`
	Expiration                 int           `json:"expiration"`
}

type Metrics struct {
	_id                        int           `json:"_id"`
	Symbol                     string        `json:"symbol"`
	Volume                     float32       `json:",string"`
	VolumeExpiration           int           `json:"volumeExpiration"`
	LastPrice                  float32       `json:",string"`
	LowestAsk                  float32       `json:",string"`
	HighestBid                 float32       `json:",string"`
	LastDayPrice               float32       `json:",string"`
	LastDayPriceExpiration     int           `json:"lastDayPruceExpiration"`
	PriceChangeHive            float32       `json:",string"`
	PriceChangePercent         string        `json:"priceChangePercent"`
}

func (h HiveEngineRpcNode) GetBook (bookType, token string, limit, offset int) (*OrderBook, error) {
	params := ContractQueryParams {
		Contract: "market",
		Table: string(strings.ToLower(bookType) + "Book"),
		Query: map[string]string{"symbol": strings.ToUpper(token)},
		Limit: limit,
		Offset: offset,
	}
	response, err := h.QueryContract(params)
	if err != nil {
		return nil, err
	}
	book := &OrderBook{}
	book.BookType = string(strings.ToTitle(bookType) + "Book")
	if uErr := json.Unmarshal(response, &book.Book); uErr != nil {
		return nil, uErr
	}

	return book, nil
}


//TODO: add other book functions like personal orders and sort function

//TODO: add trades history

func (h HiveEngineRpcNode) GetMetrics (token string, limit, offset int) (*Metrics, error) {
	params := ContractQueryParams {
		Contract: "market",
		Table: "metrics",
		Query: map[string]string{"symbol": strings.ToUpper(token)},
		Limit: limit,
		Offset: offset,
	}
	response, err := h.QueryContract(params)
	if err != nil {
		return nil, err
		//~ return "", err
	}
	metricsData := &[]Metrics{}
	if uErr := json.Unmarshal(response, &metricsData); uErr != nil {
		return nil, uErr
	}
	metrics := &(*metricsData)[0]

	return metrics, nil
}
