<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [v0.1.0] - 2021-08-20

### State Machine Breaking

* (app, rpc) [tharsis#447](https://github.com/maoaixiao1314/treasureapp/pull/447) Chain ID format has been changed from `<identifier>-<epoch>` to `<identifier>_<EIP155_number>-<epoch>`
in order to clearly distinguish permanent vs impermanent components.
* (app, evm) [tharsis#434](https://github.com/maoaixiao1314/treasureapp/pull/434) EVM `Keeper` struct and `NewEVM` function now have a new `trace` field to define
the Tracer type used to collect execution traces from the EVM transaction execution.
* (evm) [tharsis#175](https://github.com/maoaixiao1314/treasureapp/issues/175) The msg `TxData` field is now represented as a `*proto.Any`.
* (evm) [tharsis#84](https://github.com/maoaixiao1314/treasureapp/pull/84) Remove `journal`, `CommitStateDB` and `stateObjects`.
* (rpc, evm) [tharsis#81](https://github.com/maoaixiao1314/treasureapp/pull/81) Remove tx `Receipt` from store and replace it with fields obtained from the Tendermint RPC client.
* (evm) [tharsis#72](https://github.com/maoaixiao1314/treasureapp/issues/72) Update `AccessList` to use `TransientStore` instead of map.
* (evm) [tharsis#68](https://github.com/maoaixiao1314/treasureapp/issues/68) Replace block hash storage map to use staking `HistoricalInfo`.
* (evm) [tharsis#276](https://github.com/maoaixiao1314/treasureapp/pull/276) Vm errors don't result in cosmos tx failure, just
  different tx state and events.
* (evm) [tharsis#342](https://github.com/maoaixiao1314/treasureapp/issues/342) Don't clear balance when resetting the account.
* (evm) [tharsis#334](https://github.com/maoaixiao1314/treasureapp/pull/334) Log index changed to the index in block rather than
  tx.
* (evm) [tharsis#399](https://github.com/maoaixiao1314/treasureapp/pull/399) Exception in sub-message call reverts the call if it's not propagated.

### API Breaking

* (proto) [tharsis#448](https://github.com/maoaixiao1314/treasureapp/pull/448) Bump version for all Treasurenet messages to `v1`
* (server) [tharsis#434](https://github.com/maoaixiao1314/treasureapp/pull/434) `evm-rpc` flags and app config have been renamed to `json-rpc`.
* (proto, evm) [tharsis#207](https://github.com/maoaixiao1314/treasureapp/issues/207) Replace `big.Int` in favor of `sdk.Int` for `TxData` fields
* (proto, evm) [tharsis#81](https://github.com/maoaixiao1314/treasureapp/pull/81) gRPC Query and Tx service changes:
  * The `TxReceipt`, `TxReceiptsByBlockHeight` endpoints have been removed from the Query service.
  * The `ContractAddress`, `Bloom` have been removed from the `MsgEthereumTxResponse` and the
    response now contains the ethereum-formatted `Hash` in hex format.
* (evm) [#202](https://github.com/maoaixiao1314/treasureapp/pull/202) Web3 api `SendTransaction`/`SendRawTransaction` returns ethereum compatible transaction hash, and query api `GetTransaction*` also accept that.
* (rpc) [tharsis#258](https://github.com/maoaixiao1314/treasureapp/pull/258) Return empty `BloomFilter` instead of throwing an error when it cannot be found (`nil` or empty).
* (rpc) [tharsis#277](https://github.com/maoaixiao1314/treasureapp/pull/321) Fix `BloomFilter` response.

### Improvements

* (client) [tharsis#450](https://github.com/maoaixiao1314/treasureapp/issues/450) Add EIP55 hex address support on `debug addr` command.
* (server) [tharsis#343](https://github.com/maoaixiao1314/treasureapp/pull/343) Define a wrap tendermint logger `Handler` go-ethereum's `root` logger.
* (rpc) [tharsis#457](https://github.com/maoaixiao1314/treasureapp/pull/457) Configure RPC gas cap through app config.
* (evm) [tharsis#434](https://github.com/maoaixiao1314/treasureapp/pull/434) Support different `Tracer` types for the EVM.
* (deps) [tharsis#427](https://github.com/maoaixiao1314/treasureapp/pull/427) Bump ibc-go to [`v1.0.0`](https://github.com/cosmos/ibc-go/releases/tag/v1.0.0)
* (gRPC) [tharsis#239](https://github.com/maoaixiao1314/treasureapp/pull/239) Query `ChainConfig` via gRPC.
* (rpc) [tharsis#181](https://github.com/maoaixiao1314/treasureapp/pull/181) Use evm denomination for params on tx fee.
* (deps) [tharsis#423](https://github.com/maoaixiao1314/treasureapp/pull/423) Bump Cosmos SDK and Tendermint versions to [v0.43.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.43.0) and [v0.34.11](https://github.com/tendermint/tendermint/releases/tag/v0.34.11), respectively.
* (evm) [tharsis#66](https://github.com/maoaixiao1314/treasureapp/issues/66) Support legacy transaction types for signing.
* (evm) [tharsis#24](https://github.com/maoaixiao1314/treasureapp/pull/24) Implement metrics for `MsgEthereumTx`, state transitions, `BeginBlock` and `EndBlock`.
* (rpc)  [#124](https://github.com/maoaixiao1314/treasureapp/issues/124) Implement `txpool_content`, `txpool_inspect` and `txpool_status` RPC methods
* (rpc) [tharsis#112](https://github.com/maoaixiao1314/treasureapp/pull/153) Fix `eth_coinbase` to return the ethereum address of the validator
* (rpc) [tharsis#176](https://github.com/maoaixiao1314/treasureapp/issues/176) Support fetching pending nonce
* (rpc) [tharsis#272](https://github.com/maoaixiao1314/treasureapp/pull/272) do binary search to estimate gas accurately
* (rpc) [#313](https://github.com/maoaixiao1314/treasureapp/pull/313) Implement internal debug namespace (Not including logger functions nor traces).
* (rpc) [#349](https://github.com/maoaixiao1314/treasureapp/pull/349) Implement configurable JSON-RPC APIs to manage enabled namespaces.
* (rpc) [#377](https://github.com/maoaixiao1314/treasureapp/pull/377) Implement `miner_` namespace. `miner_setEtherbase` and `miner_setGasPrice` are working as intended. All the other calls are not applicable and return `unsupported`.

### Bug Fixes

* (keys) [tharsis#346](https://github.com/maoaixiao1314/treasureapp/pull/346) Fix `keys add` command with `--ledger` flag for the `secp256k1` signing algorithm.
* (evm) [tharsis#291](https://github.com/maoaixiao1314/treasureapp/pull/291) Use block proposer address (validator operator) for `COINBASE` opcode.
* (rpc) [tharsis#81](https://github.com/maoaixiao1314/treasureapp/pull/81) Fix transaction hashing and decoding on `eth_sendTransaction`.
* (rpc) [tharsis#45](https://github.com/maoaixiao1314/treasureapp/pull/45) Use `EmptyUncleHash` and `EmptyRootHash` for empty ethereum `Header` fields.
