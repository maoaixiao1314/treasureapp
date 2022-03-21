package app

import (
	"context"
	"log"
	"strings"

	//"encoding/json"
	"fmt"
	"math/big"

	//store "./contract" // for demo

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/treasurenet/x/evm/types"
	store "github.com/treasurenet/app/contract"
)

type EventLog struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data [][]interface{} `json:"data"` //`json:"-"`No serialization
}

// func NewEventLog() *EventLog {
// 	return &EventLog
// }
var Even EventLog

type NewData [][]interface{}

func getLogs(Start, End int64) {
	data := make([][]interface{}, 0)
	//Even = new(EventLog)
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		fmt.Println("No listening")
		Even = EventLog{
			Code: 300,
			Msg:  "ethclient Listening error",
			Data: data,
		}
	} else {
		//eventSignature := []byte("ItemSet(bytes32,bytes32)")
		//dongqi(address indexed account, uint256 indexed amount, uint256 variety);
		eventSignature := []byte("bidList(address,uint256,uint256)")
		hash := crypto.Keccak256Hash(eventSignature)
		topic := hash.Hex()
		fmt.Println("Test get log title new", topic)
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(Start),
			ToBlock:   big.NewInt(End),
			Topics: [][]common.Hash{
				{
					hash,
				},
			},
		}
		fmt.Println("Listening start：", client)
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			Even = EventLog{
				Code: 400,
				Msg:  "FilterLogs Listening error",
				Data: data,
			}
		}
		fmt.Println("Listening end")
		fmt.Printf("logs:%+v\n", logs)
		contractAbi, err := abi.JSON(strings.NewReader(string(store.ContractABI)))
		if err != nil {
			Even = EventLog{
				Code: 500,
				Msg:  "contractAbi Listening error",
				Data: data,
			}
		}
		fmt.Printf("contractAbi:%+v\n", contractAbi)
		if len(logs) == 0 {
			Even = EventLog{
				Code: 600,
				Msg:  "Log is empty",
				Data: data,
			}
		}
		dst := make([][]interface{}, len(logs))
		for index, vLog := range logs {
			LogAbi, err := contractAbi.Unpack("bidList", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			dst[index] = LogAbi
		}
		Even = EventLog{
			Code: 200,
			Msg:  "successful",
			Data: dst,
		}
		fmt.Printf("logabi:%+v\n", dst)
	}
}

// func getLog() (string, error) {
// 	ret := new(EventLog)
// 	//client, err := ethclient.Dial("ws://127.0.0.1:8546")
// 	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
// 	client, err := ethclient.DialContext(context.Background(), "ws://127.0.0.1:8546")
// 	fmt.Printf("client:%v\n", client)
// 	if err != nil {
// 		ret.Code = 201
// 		ret.Msg = "error"
// 		ret.Data = "connect: connection refused"
// 		data, _ := json.Marshal(ret)
// 		//log.Fatal(err)
// 		return string(data), err
// 	}

// 	//contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
// 	// query := ethereum.FilterQuery{
// 	// 	FromBlock: big.NewInt(2394201),
// 	// 	ToBlock:   big.NewInt(2394201),
// 	// 	Addresses: []common.Address{
// 	// 		contractAddress,
// 	// 	},
// 	// }
// 	eventSignature := []byte("ItemSet(bytes32,bytes32)")
// 	hash := crypto.Keccak256Hash(eventSignature)
// 	topic := hash.Hex()
// 	fmt.Println("Test get log title", topic)
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(1),
// 		ToBlock:   big.NewInt(2),
// 		Topics: [][]common.Hash{
// 			{
// 				hash,
// 			},
// 		},
// 	}
// 	fmt.Printf("query :%+v\n", query)
// 	fmt.Printf("context.Background():%+v\n", context.Background())
// 	//logs1 := make(chan types.Log)
// 	//ctx, cancel := context.WithCancel(context.Background())
// 	//ctx1, cancel := context.WithTimeout(context.Background(), time.Second*3)
// 	//ctx, cancel := Contextxt.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
// 	ch := make(chan interface{})
// 	go func() {
// 		logs, _ := client.FilterLogs(context.Background(), query)
// 		ch <- logs
// 	}()
// 	//defer client.Close()
// 	// time.Sleep(4 * time.Second)
// 	// cancel()
// 	//sub, err := client.SubscribeFilterLogs(context.Background(), query, logs1)
// 	logs2 := <-ch
// 	fmt.Printf("logs:%+v\n", logs2)
// 	//client.Close()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	ret.Code = 202
// 	// 	ret.Msg = "error"
// 	// 	ret.Data = "FilterLogs executes a filter query is error"
// 	// 	data, _ := json.Marshal(ret)
// 	// 	return string(data), err
// 	// }

// 	// contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
// 	// fmt.Printf("contractAbi:%+v\n", contractAbi)
// 	if err != nil {
// 		ret.Code = 203
// 		ret.Msg = "error"
// 		ret.Data = "contractAbi is error"
// 		data, _ := json.Marshal(ret)
// 		return string(data), err
// 	}
// 	//vLog := <-logs1
// 	// fmt.Printf("vLog:%+v\n", vLog)
// 	//Event, _ := json.Marshal(vLog)
// 	//EventLog := string(Event)
// 	//return EventLog, nil
// 	//}

// for _, vLog := range logs {
// 	fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
// 	fmt.Println(vLog.BlockNumber)     // 2394201
// 	fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

// 	event := struct {
// 		Key   [32]byte
// 		Value [32]byte
// 	}{}
// 	//err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
// 	_, err := contractAbi.Unpack("ItemSet", vLog.Data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 	fmt.Println(string(event.Key[:]))   // foo
// 	// 	fmt.Println(string(event.Value[:])) // bar

// 	// 	var topics [4]string
// 	// 	for i := range vLog.Topics {
// 	// 		topics[i] = vLog.Topics[i].Hex()
// 	// 	}

// 	// 	fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
// 	// }
// 	//eventSignature := []byte("ItemSet(bytes32,bytes32)")
// 	//hash := crypto.Keccak256Hash(eventSignature)
// 	//fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
// 	Event, _ := json.Marshal(logs2)
// 	EventLog := string(Event)
// 	return EventLog, nil
// }
