# Addresses

## Brc20

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBrc20Quantities">TimestampedBrc20Quantities</a>

Methods:

- <code title="get /addresses/{address}/brc20">client.Addresses.Brc20.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressBrc20Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBrc20Quantities">TimestampedBrc20Quantities</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runes

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedAddressRuneUtxo">PaginatedAddressRuneUtxo</a>
- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedRuneQuantities">TimestampedRuneQuantities</a>

Methods:

- <code title="get /addresses/{address}/runes/{rune}">client.Addresses.Runes.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressRuneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressRuneGetParams">AddressRuneGetParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedAddressRuneUtxo">PaginatedAddressRuneUtxo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /addresses/{address}/runes">client.Addresses.Runes.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressRuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedRuneQuantities">TimestampedRuneQuantities</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Txs

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedInvolvedTransaction">PaginatedInvolvedTransaction</a>

Methods:

- <code title="get /addresses/{address}/txs">client.Addresses.Txs.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressTxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressTxListParams">AddressTxListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedInvolvedTransaction">PaginatedInvolvedTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Utxos

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedUtxo">PaginatedUtxo</a>

Methods:

- <code title="get /addresses/{address}/utxos">client.Addresses.Utxos.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressUtxoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AddressUtxoListParams">AddressUtxoListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedUtxo">PaginatedUtxo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

## Brc20

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedBrc20Ticker">PaginatedBrc20Ticker</a>
- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBrc20Info">TimestampedBrc20Info</a>

Methods:

- <code title="get /assets/brc20/{ticker}">client.Assets.Brc20.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetBrc20Service.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ticker <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBrc20Info">TimestampedBrc20Info</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/brc20">client.Assets.Brc20.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetBrc20Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetBrc20ListParams">AssetBrc20ListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedBrc20Ticker">PaginatedBrc20Ticker</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Holders

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedBrc20Holder">PaginatedBrc20Holder</a>

Methods:

- <code title="get /assets/brc20/{ticker}/holders">client.Assets.Brc20.Holders.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetBrc20HolderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ticker <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetBrc20HolderListParams">AssetBrc20HolderListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedBrc20Holder">PaginatedBrc20Holder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runes

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedRuneIDAndName">PaginatedRuneIDAndName</a>
- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedRuneInfo">TimestampedRuneInfo</a>

Methods:

- <code title="get /assets/runes/{rune}">client.Assets.Runes.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedRuneInfo">TimestampedRuneInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/runes">client.Assets.Runes.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneListParams">AssetRuneListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedRuneIDAndName">PaginatedRuneIDAndName</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Holders

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedRuneHolder">PaginatedRuneHolder</a>

Methods:

- <code title="get /assets/runes/{rune}/holders">client.Assets.Runes.Holders.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneHolderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneHolderListParams">AssetRuneHolderListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedRuneHolder">PaginatedRuneHolder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Utxos

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedRuneUtxo">PaginatedRuneUtxo</a>

Methods:

- <code title="get /assets/runes/{rune}/utxos">client.Assets.Runes.Utxos.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneUtxoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#AssetRuneUtxoListParams">AssetRuneUtxoListParams</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#PaginatedRuneUtxo">PaginatedRuneUtxo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Blocks

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBlock">TimestampedBlock</a>

Methods:

- <code title="get /blocks/{block_hash}">client.Blocks.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#BlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, blockHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBlock">TimestampedBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Latest

Methods:

- <code title="get /blocks/latest">client.Blocks.Latest.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#BlockLatestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBlock">TimestampedBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# General

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBlockchainInfo">TimestampedBlockchainInfo</a>

Methods:

- <code title="get /general/info">client.General.Info.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#GeneralInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedBlockchainInfo">TimestampedBlockchainInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rpc

## Mempool

### Info

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolInfo">TimestampedMempoolInfo</a>

Methods:

- <code title="get /rpc/mempool/info">client.Rpc.Mempool.Info.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#RpcMempoolInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolInfo">TimestampedMempoolInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactionDetails">TimestampedMempoolTransactionDetails</a>
- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactions">TimestampedMempoolTransactions</a>

Methods:

- <code title="get /rpc/mempool/transactions/{tx_hash}">client.Rpc.Mempool.Transactions.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#RpcMempoolTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactionDetails">TimestampedMempoolTransactionDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rpc/mempool/transactions">client.Rpc.Mempool.Transactions.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#RpcMempoolTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactions">TimestampedMempoolTransactions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Ancestors

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactionAncestors">TimestampedMempoolTransactionAncestors</a>

Methods:

- <code title="get /rpc/mempool/transactions/{tx_hash}/ancestors">client.Rpc.Mempool.Transactions.Ancestors.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#RpcMempoolTransactionAncestorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactionAncestors">TimestampedMempoolTransactionAncestors</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Descendants

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactionDescendants">TimestampedMempoolTransactionDescendants</a>

Methods:

- <code title="get /rpc/mempool/transactions/{tx_hash}/descendants">client.Rpc.Mempool.Transactions.Descendants.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#RpcMempoolTransactionDescendantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedMempoolTransactionDescendants">TimestampedMempoolTransactionDescendants</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RpcTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedTransactionDetails">TimestampedTransactionDetails</a>

Methods:

- <code title="get /rpc/transactions/{tx_hash}">client.RpcTransactions.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#RpcTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TimestampedTransactionDetails">TimestampedTransactionDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Methods:

- <code title="get /transactions/{tx_hash}">client.Transactions.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /transactions/submit">client.Transactions.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TransactionService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TransactionSubmitParams">TransactionSubmitParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Psbt

Methods:

- <code title="post /transactions/psbt/decode">client.Transactions.Psbt.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TransactionPsbtService.Decode">Decode</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk">maestrobitcoingosdk</a>.<a href="https://pkg.go.dev/github.com/maestro-org/maestro-bitcoin-go-sdk#TransactionPsbtDecodeParams">TransactionPsbtDecodeParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
