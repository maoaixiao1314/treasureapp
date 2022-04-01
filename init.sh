
KEY="mykey"
CHAINID="treasurenet_9000-1"
MONIKER="localtestnet"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# to trace evm
TRACE="--trace"
# TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon and client
rm -rf ~/.treasurenet*

make install

treasurenetd config keyring-backend $KEYRING
treasurenetd config chain-id $CHAINID

# if $KEY exists it should be deleted
treasurenetd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for Treasurenet (Moniker can be anything, chain-id must be an integer)
treasurenetd init $MONIKER --chain-id $CHAINID 

# Change parameter token denominations to aunit
cat $HOME/.treasurenetd/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aunit"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json
cat $HOME/.treasurenetd/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aunit"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json
cat $HOME/.treasurenetd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aunit"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json
cat $HOME/.treasurenetd/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="aunit"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json

# increase block time (?)
cat $HOME/.treasurenetd/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="30000"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json

# Set gas limit in genesis
cat $HOME/.treasurenetd/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.treasurenetd/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.treasurenetd/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.treasurenetd/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.treasurenetd/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.treasurenetd/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
treasurenetd add-genesis-account $KEY 100000000000000000000000000aunit --keyring-backend $KEYRING

# Sign genesis transaction
treasurenetd gentx $KEY 1000000000000000000000aunit --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
treasurenetd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
treasurenetd validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed) --json-rpc.ws-address localhost:8546
#./treasurenetd query staking validator ethvaloper1s8znfe26xwq8lm5dhzxh4rjr7n90wrl2hn4dwj -o json --chain-id treasurenet_4143179869527-1 --home /data/mytreasurenet/node8//treasurenetd
#./treasurenetd tx slashing unjail --from node8 --chain-id treasurenet_4143179869527-1 --home /data/mytreasurenet/node8/treasurenetd --keyring-backend test --gas-prices 20aunit
#treasurenetd start --pruning=nothing --trace --log_level info --minimum-gas-prices=0.0001aunit --json-rpc.api eth,txpool,personal,net,debug,web3,miner --p2p.persistent_peers=a2f4cd782091b4429efc0748109d326053c99a59@54.241.61.93:26656,56490415fafa170557e410069faf05e3077500fd@54.193.100.97:26656,a82f9a917b35c77c572f6b9139635cec3c5f0317@18.144.45.142:26656
treasurenetd start --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001aunit --json-rpc.api eth,txpool,personal,net,debug,web3,miner
