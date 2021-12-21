package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	//store "./contract" // for demo
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/tharsis/ethermint/app/contract"
)

// type EventLog struct {
// 	Code int    `json:"code"`
// 	Msg  string `json:"msg"`
// 	Data string `json:"data"` //`json:"-"`不进行序列化
// }

func main() {
	client, err := ethclient.Dial("ws://localhost:8546")
	fmt.Printf("client:%v\n", client)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}

	//contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	// query := ethereum.FilterQuery{
	// 	FromBlock: big.NewInt(2394201),
	// 	ToBlock:   big.NewInt(2394201),
	// 	Addresses: []common.Address{
	// 		contractAddress,
	// 	},
	// }
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	topic := hash.Hex()
	fmt.Println("测试获取日志标题", topic)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1),
		ToBlock:   big.NewInt(100),
		Topics: [][]common.Hash{
			{
				hash,
			},
		},
	}
	fmt.Printf("query:%+v\n", query)
	fmt.Printf("context.Background():%+v\n", context.Background())
	logs, err := client.FilterLogs(context.Background(), query)
	fmt.Printf("logs:%+v\n", logs)
	if err != nil {
		fmt.Println("error")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	fmt.Printf("contractAbi:%+v\n", contractAbi)
	if err != nil {
		fmt.Println("error")
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
		//err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		sub, err := contractAbi.Unpack("ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sub:%+v\n", sub)
		for _, log := range sub {
			Log1, _ := log.([32]byte)
			//Event, _ := json.Marshal(Log1)
			EventLog := string(Log1[:])
			fmt.Printf("log:%+v\n", EventLog)
		}
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	//eventSignature := []byte("ItemSet(bytes32,bytes32)")
	//hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	// Event, _ := json.Marshal(contractAbi)
	// EventLog := string(Event)
	// return EventLog, nil
}
