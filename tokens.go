package hiveenginego

import (
	//~ "fmt"
	"strings"
	"encoding/json"
)

type Balances struct {
	_id                      int     `json:"_id"`
	Account                  string  `json:"account"`
	Symbol                   string  `json:"symbol"`
	Balance                  string  `json:"balance"`
	Stake                    string  `json:"stake"`
	PendingUnstake           string  `json:"pendingUnstake"`
	DelegationsIn            string  `json:"delegationsIn"`
	DelegationsOut           string  `json:"delegationsOut"`
	PendingUndelegations     string  `json:"pendingUndelegations"`
}

func (h HiveEngineRpcNode) GetBalances (token, account string, limit, offset int) (*Balances, error) {
	//~ q := make(map[string]string)
	//~ q["symbol"]  = token
	//~ q["account"] = account
	params := ContractQueryParams {
		Contract: "tokens",
		Table: "balances",
		Query: map[string]string{"symbol": strings.ToUpper(token), "account": strings.ToLower(account)},
		Limit: limit,
		Offset: offset,
	}
	response, err := h.QueryContractByAcc(params)
	if err != nil {
		return nil, err
	}
	bals := &Balances{}
	if uErr := json.Unmarshal(response, &bals); uErr != nil {
		return nil, uErr
	}
	return bals, nil
}

//eventually add GetEveryonesBalances (find) - maybe

//eventually add GetMultiTokenBalances (findOne/batch) - maybe


