#!/bin/bash

echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./ethermintd validate-genesis --home /ethermint

echo "starting ethermint node $ID in background ..."
./ethermintd start \
--home /ethermint \
--keyring-backend test \
--json-rpc.api eth,txpool,personal,net,debug,web3,miner

echo "started ethermint node"
tail -f /dev/null