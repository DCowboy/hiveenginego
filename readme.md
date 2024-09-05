# HiveEngineGo - A client for the Hive Engine side chain on the Hive blockchain

This fork will be adding more functionality as it seems like the upstream repo might be abandoned. New functions will be added to the bottom, but above the warning.

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
// Numbers above are limit and offset
// Returns a struct
stake := balances.Stake
balance := balances.Balance
```

WARNING: It is not recommended to stream blocks from public APIs. They are provided as a service to users and saturating them with block requests may (rightfully) result in your IP getting banned
