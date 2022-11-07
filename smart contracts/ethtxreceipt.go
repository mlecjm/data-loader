package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Smart contract - Chaincode structure definition
// ===============================================
type EthDataLoaderChaincode struct {
}

type ethData struct {
	ObjectType        string `json:"docType"`           // .
	Eth_id            string `json:"eth_id"`            // ethData index.
	TransactionHash   string `json:"transactionHash"`   // hash of the transaction.
	From              string `json:"from"`              // address of the sender.
	To                string `json:"to"`                // address of the receiver. null when its a contract creation transaction.
	CumulativeGasUsed string `json:"cumulativeGasUsed"` // The total amount of gas used when this transaction was executed in the block.
	gasUsed           string `json:"gasUsed"`           // The amount of gas used by this specific transaction alone.
	ContractAddress   string `json:"contractAddress"`   // The contract address created, if the transaction was a contract creation, otherwise null.
}

// =============================================================================
// 			       Smart Contract/Chaincode Main Function
// =============================================================================

func main() {
	err := shim.Start(new(EthDataLoaderChaincode))
	if err != nil {
		fmt.Printf("Error starting ethDataLoader chaincode: %s", err)
	}
}

// Init - initializes chaincode
// =============================

func (t *EthDataLoaderChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println(" EtheDataloader Chaincode Initilized.")
	return shim.Success(nil)
}

// Invoke - the entry point for function invocations
// =====================================================

func (t *EthDataLoaderChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke method gets called ")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" { // create an eth data instance
		return t.invoke(stub, args)
	} else if function == "query" { //read an eth data instance
		return t.query(stub, args)
	}

	fmt.Println("invoke did not find function: " + function)
	return shim.Error("Unknown function name. Expecting \"invoke\"\"query\"")
}

// ===================================================================
// invoke - set an EthTransaction Data instance from chaincode state
// ===================================================================
func (t *EthDataLoaderChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("- start init EthTranction Data")

	// Eth_id:"0001", TransactionHash: "0xabcd1245", From: "0x122", To: "0x234", CumulativeGasUsed: "9000000", gasUsed: "9000", ContractAddress: "0xakbkb1245"
	// "0001", "0xabcd1245", "0x122",  "0x234", "9000000", "9000", "0xakbkb1245"

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	eth_id := args[0]
	transactionHash := args[1]
	from := args[2]
	to := args[3]
	cumulativeGasUsed := args[4]
	gasUsed := args[5]
	contractAddress := args[6]

	objectType := "ethData"
	ethData := &ethData{objectType, eth_id, transactionHash, from, to, cumulativeGasUsed, gasUsed, contractAddress}
	ethDataJSONasBytes, err := json.Marshal(ethData)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(eth_id, ethDataJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init ethTransaction data")
	return shim.Success(nil)
}

// ==============================================================
// query - read an Ethereum Data instance from chaincode state
// ==============================================================

func (t *EthDataLoaderChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var eth_id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting only from address to query")
	}

	eth_id = args[0]
	valAsBytes, err := stub.GetState(eth_id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state of: " + eth_id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsBytes == nil {
		jsonResp = "{\"Error: \" \"The following eth data instance does not exist: " + eth_id + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsBytes)
}
