#!/bin/bash

# echo "Allocate genesis accounts (cosmos formatted addresses)"
# ./treasurenetd add-genesis-account $KEY 100000000000000000000000000aunit --home /treasurenet --keyring-backend $KEYRING
# echo "Sign genesis transaction"
# treasurenetd gentx $KEY 1000000000000000000000aunit --keyring-backend $KEYRING --chain-id $CHAINID
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./treasurenetd validate-genesis --home /treasurenet

echo "starting treasurenet node $ID in background ..."
./treasurenetd start \
--home /treasurenet \
--keyring-backend test \
--json-rpc.api eth,txpool,personal,net,debug,web3,miner \
--pruning=nothing \
--trace


echo "started treasurenet node"
tail -f /dev/null