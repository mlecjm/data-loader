package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Smart contract - Chaincode structure definition
// ===============================================
type IconTxDataLoaderChaincode struct {
}

type iconTxData struct {
	ObjectType string `json:"docType"`   // .
	Icx_id     string `json:"icx_id"`    // Icon transaction data index
	Version    string `json:"version"`   // Protocol version ("0x3" for V3)
	From       string `json:"from"`      // EOA address that created the transaction
	To         string `json:"to"`        // EOA address to receive coins, or SCORE address to execute the transaction
	Value      string `json:"value"`     // Amount of ICX coins in loop to transfer. When ommitted, assumes 0. (1 icx = 1 ^ 18 loop)
	StepLimit  string `json:"stepLimit"` // Maximum step allowance that can be used by the transaction
	Timestamp  string `timestamp"`       // Transaction creation time. timestamp is in microsecond
	Nid        string `nid"`             // Network ID
	Nonce      string `nonce"`           // An arbitrary number used to prevent transaction hash collision
	TxHash     string `txHash"`          // Transaction hash
	DataType   string `dataType"`        // Type of data. (call, deploy, message or deposit)
}

// =============================================================================
// 			       Smart Contract/Chaincode Main Function
// =============================================================================

func main() {
	err := shim.Start(new(IconTxDataLoaderChaincode))
	if err != nil {
		fmt.Printf("Error starting iconTxDataLoader chaincode: %s", err)
	}
}

// Init - initializes chaincode
// =============================

func (t *IconTxDataLoaderChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println(" IconTxDataloader Chaincode Init.")
	return shim.Success(nil)
}

// Invoke - the entry point for function invocations
// =====================================================

func (t *IconTxDataLoaderChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
func (t *IconTxDataLoaderChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("- start init icon transaction data")

	// Icx_id:"0001", Version: "0x3", From: "hxd48ab6a11220", To: "hxd48ab6a11220", Value: "0x1bc16d67", StepLimit: "0x2faf080", Timestamp: "0x5cdbea6426a20", Nid: "0x1", Nonce: "0x1", TxHash: "0xc454d40939728d652e38", Value: "call"
	// "0001", "0x3", "hxd48ab6a11220", "hxd48ab6a11220", "0x1bc16d67", "0x2faf080", "0x5cdbea6426a20","0x1", "0x1", "0xc454d40939728d652e38","call"

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	icx_id := args[0]
	version := args[1]
	from := args[2]
	to := args[3]
	value := args[4]
	stepLimit := args[5]
	timestamp := args[6]
	nid := args[7]
	nonce := args[8]
	txHash := args[9]
	dataType := args[10]

	objectType := "iconTxData"
	iconTxData := &iconTxData{objectType, icx_id, version, from, to, value, stepLimit, timestamp, nid, nonce, txHash, dataType}
	iconTxDataJSONasBytes, err := json.Marshal(iconTxData)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(icx_id, iconTxDataJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init icon Transaction data")
	return shim.Success(nil)
}

// ==============================================================
// query - read an Ethereum Data instance from chaincode state
// ==============================================================

func (t *IconTxDataLoaderChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var icx_id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting only from address to query")
	}

	icx_id = args[0]
	valAsBytes, err := stub.GetState(icx_id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state of: " + icx_id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsBytes == nil {
		jsonResp = "{\"Error: \" \"The following eth data instance does not exist: " + icx_id + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsBytes)
}
