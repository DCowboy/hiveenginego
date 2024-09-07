package hiveenginego

import (
	"fmt"
	"strings"
	"encoding/json"
)

type OrderBook struct {
	BookType   string
	Book       []Order     `json:""`
}

type Order struct {
	_id           int             `json:"_id"`
	TxId          string          `json:"txId"`
	Timestamp     int             `json:"timestamp"`
	Account       string          `json:"account"`
	Symbol        string          `json:"symbol"`
	Quantity      float32         `json:",string"`
	Price         float32         `json:",string"`
	PriceDec      interface{}     `json:"priceDec"`
	TokensLocked  float32          `json:",string, omitempty"`
	Expiration    int             `json:"expiration"`
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
		fmt.Println("Actually tried to unmarshall but...")
		return nil, uErr
	}

	return book, nil
}
