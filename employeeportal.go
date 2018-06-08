package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// EmployeeChaincode example simple Chaincode implementation
type EmployeeChaincode struct {
}

// Init initializes the chaincode
func (t *EmployeeChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("employee Init")
	return shim.Success(nil)
}

//Invoke the chaincode
func (t *EmployeeChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Employee Chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *EmployeeChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

// Deletes an entity from state
func (t *EmployeeChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *EmployeeChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(EmployeeChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
