<!--
order: 1
-->

# Concepts

## EVM

The Ethereum Virtual Machine [EVM](https://ethereum.org/en/developers/docs/evm/)  is a state machine
that provides the necessary tools to run or create a contract on a given state.

## State DB

The `StateDB` interface from geth represents an EVM database for full state querying of both
contracts and accounts. The concrete type that fulfills this interface on Treasurenet is the
`CommitStateDB`.

## Genesis State

The `x/evm` module `GenesisState` defines the state necessary for initializing the chain from a previous exported height.


### Genesis Accounts

The `GenesisAccount` type corresponds to an adaptation of the Ethereum `GenesisAccount` type. Its
main difference is that the one on Treasurenet uses a custom `Storage` type that uses a slice instead
of maps for the evm `State` (due to non-determinism), and that it doesn't contain the private key
field.

It is also important to note that since the `auth` and `bank` SDK modules manage the accounts and
balance state,  the `Address` must correspond to an `EthAccount` that is stored in the `auth`'s
module `AccountKeeper` and the balance must match the balance of the `EvmDenom` token denomination
defined on the `GenesisState`'s `Param`. The values for the address and the balance amount maintain
the same format as the ones from the SDK to make manual inspections easier on the genesis.json.


### Transaction Logs

On every Treasurenet transaction, its result contains the Ethereum `Log`s from the state machine
execution that are used by the JSON-RPC Web3 server for for filter querying. Since Cosmos upgrades
don't persist the transactions on the blockchain state, we need to persist the logs the EVM module
state to prevent the queries from failing.

`TxsLogs` is the field that contains all the transaction logs that need to be persisted after an
upgrade. It uses an array instead of a map to ensure determinism on the iteration.


### Chain Config

The `ChainConfig` is a custom type that contains the same fields as the go-ethereum `ChainConfig`
parameters, but using `sdk.Int` types instead of `*big.Int`. It also defines additional YAML tags
for pretty printing.

The `ChainConfig` type is not a configurable SDK `Param` since the SDK does not allow for validation
against a previous stored parameter values or `Context` fields. Since most of this type's fields
rely on the block height value, this limitation prevents the validation of of potential new
parameter values against the current block height (eg: to prevent updating the config block values
to a past block).

If you want to update the config values, use an software upgrade procedure.


### Params

See the [params](07_params.md) document for further information about parameters.
