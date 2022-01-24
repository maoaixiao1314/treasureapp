#!/bin/bash

CHAINID="treasurenet_9000-1"
MONIKER="localtestnet"

# localKey address 0x7cb61d4117ae31a12e393a1cfa3bac666481d02e
VAL_KEY="localkey"
VAL_MNEMONIC="gesture inject test cycle original hollow east ridge hen combine junk child bacon zero hope comfort vacuum milk pitch cage oppose unhappy lunar seat"

# user1 address 0xc6fe5d33615a1c52c08018c47e8bc53646a0e101
USER1_KEY="user1"
USER1_MNEMONIC="copper push brief egg scan entry inform record adjust fossil boss egg comic alien upon aspect dry avoid interest fury window hint race symptom"

# user2 address 0x963ebdf2e1f8db8707d05fc75bfeffba1b5bac17
USER2_KEY="user2"
USER2_MNEMONIC="maximum display century economy unlock van census kite error heart snow filter midnight usage egg venture cash kick motor survey drastic edge muffin visual"

# remove existing daemon and client
rm -rf ~/.treasurenet*

# Import keys from mnemonics
echo $VAL_MNEMONIC | treasurenetd keys add $VAL_KEY --recover --keyring-backend test --algo "eth_secp256k1"
echo $USER1_MNEMONIC | treasurenetd keys add $USER1_KEY --recover --keyring-backend test --algo "eth_secp256k1"
echo $USER2_MNEMONIC | treasurenetd keys add $USER2_KEY --recover --keyring-backend test  --algo "eth_secp256k1"

treasurenetd init $MONIKER --chain-id $CHAINID

# Set gas limit in genesis
cat $HOME/.treasurenetd/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.treasurenetd/config/tmp_genesis.json && mv $HOME/.treasurenetd/config/tmp_genesis.json $HOME/.treasurenetd/config/genesis.json

# Allocate genesis accounts (cosmos formatted addresses)
treasurenetd add-genesis-account "$(treasurenetd keys show $VAL_KEY -a --keyring-backend test)" 1000000000000000000000aunit,1000000000000000000stake --keyring-backend test
treasurenetd add-genesis-account "$(treasurenetd keys show $USER1_KEY -a --keyring-backend test)" 1000000000000000000000aunit,1000000000000000000stake --keyring-backend test
treasurenetd add-genesis-account "$(treasurenetd keys show $USER2_KEY -a --keyring-backend test)" 1000000000000000000000aunit,1000000000000000000stake --keyring-backend test

# Sign genesis transaction
treasurenetd gentx $VAL_KEY 1000000000000000000stake --amount=1000000000000000000000aunit --chain-id $CHAINID --keyring-backend test

# Collect genesis tx
treasurenetd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
treasurenetd validate-genesis

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
treasurenetd start --pruning=nothing --rpc.unsafe --keyring-backend test --trace --log_level info --json-rpc.api eth,txpool,personal,net,debug,web3
