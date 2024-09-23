# HiveEngineGo - A client for the Hive Engine side chain on the Hive blockchain

This fork will be adding more functionality as it seems like the upstream repo isn't getting the love it should. New functions will be added to the bottom, but above the warning.

At this time, there are only a few functions from the client. More will be added.

### Example usage:
create a client:
```
herpc := hiveenginego.NewHiveEngineRpc("http://MyHiveEngineApi")
```

Query latest block info:
```
latestBlockInfo, err := herpc.GetLatestBlockInfo()
//Returns a struct
latestBlockNum := latestBlockInfo.BlockNumber
```

Get All NFT of a given symbol (return rpc resonse as raw bytes):
```
rawNftBytes, err := herpc.GetSymbolAllNftFast("STAR")
```

Get block range as the raw response from the rpc (in bytes):
```
rpcResponsesBytes, err := herpc.GetBlockRangeFast(start, end)
```

Get an account's balances for a token:
```
balances, err :=  herpc.GetBalances("BEE", "alice", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive
// Returns a struct
stake := balances.Stake
balance := balances.Balance
etc.
```

Get buy/sell books for a token:
```
book, err :=  herpc.GetBook("buy", "BEE", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
// Returns a struct: book is still an array/slice
buyBook := book.Book
firstPrice := buyBook[0].Price
```

Get account's open orders for a token for a token:
```
orders, err :=  herpc.GetAccountOrders("BEE", "Alice", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
// Returns a struct of each book struct (still returned as a slice)
buyOrders := orders.Buy
firstPrice := buyOrders.Book[0].Price
```
Get trade history for a token:
```
history, err :=  herpc.GetHistory("BEE", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
// Returns a struct of an array of records
log := history.Log
firstRecord := log[0]
firstRecordTimestamp := log[0].Timestamp
```

Get metrics for a token:
```
metrics, err :=  herpc.GetMetrics("BEE", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
// Returns an array of a struct - Metrics returns as an array because of the query method it uses.
highest := (*response)[0].HighestBid
```

WARNING: It is not recommended to stream blocks from public APIs. They are provided as a service to users and saturating them with block requests may (rightfully) result in your IP getting banned
