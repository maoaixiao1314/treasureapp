#!/bin/bash

# echo "Allocate genesis accounts (cosmos formatted addresses)"
# ./ethermintd add-genesis-account $KEY 100000000000000000000000000aphoton --home /ethermint --keyring-backend $KEYRING
# echo "Sign genesis transaction"
# ethermintd gentx $KEY 1000000000000000000000aphoton --keyring-backend $KEYRING --chain-id $CHAINID
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./ethermintd validate-genesis --home /ethermint

echo "starting ethermint node $ID in background ..."
./ethermintd start \
--home /ethermint \
--keyring-backend test \
--json-rpc.api eth,txpool,personal,net,debug,web3,miner \
--pruning=nothing \
--trace


echo "started ethermint node"
tail -f /dev/null