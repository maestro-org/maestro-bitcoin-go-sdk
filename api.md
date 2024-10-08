# Addresses

## Brc20

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBrc20Quantities">TimestampedBrc20Quantities</a>

Methods:

- <code title="get /addresses/{address}/brc20">client.Addresses.Brc20.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressBrc20Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBrc20Quantities">TimestampedBrc20Quantities</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedAddressRuneUtxo">PaginatedAddressRuneUtxo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedRuneQuantities">TimestampedRuneQuantities</a>

Methods:

- <code title="get /addresses/{address}/runes/{rune}">client.Addresses.Runes.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressRuneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressRuneGetParams">AddressRuneGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedAddressRuneUtxo">PaginatedAddressRuneUtxo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /addresses/{address}/runes">client.Addresses.Runes.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressRuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedRuneQuantities">TimestampedRuneQuantities</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Txs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedInvolvedTransaction">PaginatedInvolvedTransaction</a>

Methods:

- <code title="get /addresses/{address}/txs">client.Addresses.Txs.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressTxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressTxListParams">AddressTxListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedInvolvedTransaction">PaginatedInvolvedTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Utxos

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedUtxo">PaginatedUtxo</a>

Methods:

- <code title="get /addresses/{address}/utxos">client.Addresses.Utxos.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressUtxoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, address <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AddressUtxoListParams">AddressUtxoListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedUtxo">PaginatedUtxo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

## Brc20

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedBrc20Ticker">PaginatedBrc20Ticker</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBrc20Info">TimestampedBrc20Info</a>

Methods:

- <code title="get /assets/brc20/{ticker}">client.Assets.Brc20.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetBrc20Service.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ticker <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBrc20Info">TimestampedBrc20Info</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/brc20">client.Assets.Brc20.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetBrc20Service.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetBrc20ListParams">AssetBrc20ListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedBrc20Ticker">PaginatedBrc20Ticker</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Holders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedBrc20Holder">PaginatedBrc20Holder</a>

Methods:

- <code title="get /assets/brc20/{ticker}/holders">client.Assets.Brc20.Holders.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetBrc20HolderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ticker <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetBrc20HolderListParams">AssetBrc20HolderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedBrc20Holder">PaginatedBrc20Holder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedRuneIDAndName">PaginatedRuneIDAndName</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedRuneInfo">TimestampedRuneInfo</a>

Methods:

- <code title="get /assets/runes/{rune}">client.Assets.Runes.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedRuneInfo">TimestampedRuneInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/runes">client.Assets.Runes.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneListParams">AssetRuneListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedRuneIDAndName">PaginatedRuneIDAndName</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Holders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedRuneHolder">PaginatedRuneHolder</a>

Methods:

- <code title="get /assets/runes/{rune}/holders">client.Assets.Runes.Holders.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneHolderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneHolderListParams">AssetRuneHolderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedRuneHolder">PaginatedRuneHolder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Utxos

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedRuneUtxo">PaginatedRuneUtxo</a>

Methods:

- <code title="get /assets/runes/{rune}/utxos">client.Assets.Runes.Utxos.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneUtxoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rune <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#AssetRuneUtxoListParams">AssetRuneUtxoListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#PaginatedRuneUtxo">PaginatedRuneUtxo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Blocks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBlock">TimestampedBlock</a>

Methods:

- <code title="get /blocks/{block_hash}">client.Blocks.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#BlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, blockHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBlock">TimestampedBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Latest

Methods:

- <code title="get /blocks/latest">client.Blocks.Latest.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#BlockLatestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBlock">TimestampedBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# General

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBlockchainInfo">TimestampedBlockchainInfo</a>

Methods:

- <code title="get /general/info">client.General.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#GeneralInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedBlockchainInfo">TimestampedBlockchainInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rpc

## Mempool

### Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolInfo">TimestampedMempoolInfo</a>

Methods:

- <code title="get /rpc/mempool/info">client.Rpc.Mempool.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#RpcMempoolInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolInfo">TimestampedMempoolInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactionDetails">TimestampedMempoolTransactionDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactions">TimestampedMempoolTransactions</a>

Methods:

- <code title="get /rpc/mempool/transactions/{tx_hash}">client.Rpc.Mempool.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#RpcMempoolTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactionDetails">TimestampedMempoolTransactionDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rpc/mempool/transactions">client.Rpc.Mempool.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#RpcMempoolTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactions">TimestampedMempoolTransactions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Ancestors

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactionAncestors">TimestampedMempoolTransactionAncestors</a>

Methods:

- <code title="get /rpc/mempool/transactions/{tx_hash}/ancestors">client.Rpc.Mempool.Transactions.Ancestors.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#RpcMempoolTransactionAncestorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactionAncestors">TimestampedMempoolTransactionAncestors</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Descendants

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactionDescendants">TimestampedMempoolTransactionDescendants</a>

Methods:

- <code title="get /rpc/mempool/transactions/{tx_hash}/descendants">client.Rpc.Mempool.Transactions.Descendants.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#RpcMempoolTransactionDescendantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedMempoolTransactionDescendants">TimestampedMempoolTransactionDescendants</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RpcTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedTransactionDetails">TimestampedTransactionDetails</a>

Methods:

- <code title="get /rpc/transactions/{tx_hash}">client.RpcTransactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#RpcTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TimestampedTransactionDetails">TimestampedTransactionDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Methods:

- <code title="get /transactions/{tx_hash}">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, txHash <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /transactions/submit">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TransactionService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TransactionSubmitParams">TransactionSubmitParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Psbt

Methods:

- <code title="post /transactions/psbt/decode">client.Transactions.Psbt.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TransactionPsbtService.Decode">Decode</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go">maestrobitcoin</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/maestro-bitcoin-go#TransactionPsbtDecodeParams">TransactionPsbtDecodeParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
