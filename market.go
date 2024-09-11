package hiveenginego

import (
	//~ "fmt"
	//~ "errors"
	"bytes"
	"strings"
	"encoding/json"
)

type PersonalOrders struct {
	Buy                       OrderBook
	Sell                      OrderBook
}

type OrderBook struct {
	Book                       []Order       `json:""`
}

type Order struct {
	_id                        int           `json:"_id"`
	TxId                       string        `json:"txId"`
	Timestamp                  int           `json:"timestamp"`
	Account                    string        `json:"account"`
	Symbol                     string        `json:"symbol"`
	Quantity                   float64       `json:",string"`
	Price                      float64       `json:",string"`
	PriceDec                   interface{}   `json:"priceDec"`
	TokensLocked               float64       `json:",string, omitempty"`
	Expiration                 int           `json:"expiration"`
}

type History struct {
	Log                        []Record
}

type Record struct {
	_id                        int            `json:"_id"`
	Type                       string         `json:"type"`
	Buyer                      string         `json:"buyer"`
	Seller                     string         `json:"seller"`
	Symbol                     string         `json:"symbol"`
	Quantity                   float64        `json:",string"`
	Price                      float64        `json:",string"`
	Timestamp                  int            `json:"timestamp"`
	Volume                     float64        `json:",string"`
	BuyTxId                    string         `json:"buyTxId"`
	SellTxId                   string         `json:"sellTxId"`
}

type Metrics struct {
	_id                        int           `json:"_id"`
	Symbol                     string        `json:"symbol"`
	Volume                     float64       `json:",string"`
	VolumeExpiration           int           `json:"volumeExpiration"`
	LastPrice                  float64       `json:",string"`
	LowestAsk                  float64       `json:",string"`
	HighestBid                 float64       `json:",string"`
	LastDayPrice               float64       `json:",string"`
	LastDayPriceExpiration     int           `json:"lastDayPruceExpiration"`
	PriceChangeHive            float64       `json:",string"`
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
	c := bytes.TrimLeft(response, " \t\r\n")
	if len(c) > 0 && c[0] == '[' {
		if uErr := json.Unmarshal(response, &book.Book); uErr != nil {
			return nil, uErr
		}
	} else if len(c) > 0 && c[0] == '{' {
		order := &Order{}
		if uErr := json.Unmarshal(response, &order); uErr != nil {
			return nil, uErr
		}
		book.Book = append(book.Book, *order)
	} else {
		book.Book = make([]Order, 0)
	}

	return book, nil
}

func (h HiveEngineRpcNode) GetAccountOrders (token, account string, limit, offset int) (*PersonalOrders, error) {
	orders := &PersonalOrders{}
	actions := []string{"buy", "sell"}
	for _, action := range actions {
		params := ContractQueryParams {
			Contract: "market",
			Table: string(strings.ToLower(action) + "Book"),
			Query: map[string]string{"symbol": strings.ToUpper(token), "account": strings.ToLower(account)},
			Limit: limit,
			Offset: offset,
		}
		response, err := h.QueryContractByAcc(params)
		if err != nil {
			return nil, err
		}
		book := &OrderBook{}
		c := bytes.TrimLeft(response, " \t\r\n")
		if len(c) > 0 && c[0] == '[' {
			if uErr := json.Unmarshal(response, &book.Book); uErr != nil {
				return nil, uErr
			}
		} else if len(c) > 0 && c[0] == '{' {
			order := &Order{}
			if uErr := json.Unmarshal(response, &order); uErr != nil {
				return nil, uErr
			}
			book.Book = append(book.Book, *order)
		} else {
			book.Book = make([]Order, 0)
		}
		

		if action == "buy" {
			orders.Buy = *book
		} else {
			orders.Sell = *book
		}
	}
	return orders, nil
}

//TODO: add other book functions like sort function

func (h HiveEngineRpcNode) GetHistory (token string, limit, offset int) (*History, error) {
	params := ContractQueryParams {
		Contract: "market",
		Table: "tradesHistory",
		Query: map[string]string{"symbol": strings.ToUpper(token)},
		Limit: limit,
		Offset: offset,
	}
	response, err := h.QueryContract(params)
	if err != nil {
		return nil, err
	}
	history := &History{}
	c := bytes.TrimLeft(response, " \t\r\n")
	if len(c) > 0 && c[0] == '[' {
		if uErr := json.Unmarshal(response, &history.Log); uErr != nil {
			return nil, uErr
		}
	} else if len(c) > 0 && c[0] == '{' {
		record := &Record{}
		if uErr := json.Unmarshal(response, &record); uErr != nil {
			return nil, uErr
		}
		history.Log = append(history.Log, *record)
	} else {
		history.Log = make([]Record, 0)
	}

	return history, nil
}

//Note: findOne does not work for tradesHistory or metrics tables. Will have to sort tradesHistory at a high limit to get account specific history

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
