package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func CleanUp(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) pb.Response {
	docTypes := []string{BUSINESS_DOCTYPE, DATA_CATEGORY_DOCTYPE, DATA_CONTRACT_DOCTYPE, REVIEW_DOCTYPE, DATA_CONTRACT_TYPE_DOCTYPE, PERSON_DOCTYPE, ACCOUNT_DOCTYPE, TOKEN_DOCTYPE}
	//docTypes := []string{BUSINESS_DOCTYPE}

	for _, docType := range docTypes {
		err := deleteState(logger, stub, docType)
		if err != nil {
			logger.Error(err.Error())
			return shim.Error(err.Error())
		}
	}
	return shim.Success([]byte(""))
}

func DeleteDoc(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, doc string) pb.Response {
	logger.Info("entering-delete-doc")
	defer logger.Info("exit-delete-doc")

	if doc == "" {
		logger.Debug("fcannot-delete-empty-ID")
		return shim.Error("cannot-delete-empty-ID")
	}

	err := stub.DelState(doc)
	if err != nil {
		logger.Debug("failed-deleting-doc-%s", err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func deleteState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, docType string) error {
	queryResults, err := getDocumentsByDocType(logger, stub, docType)
	if err != nil {
		return err
	}
	var documents []interface{}
	err = json.Unmarshal(queryResults, &documents)
	if err != nil {
		return err
	}
	for _, document := range documents {
		mapDocument := document.(map[string]interface{})
		id, ok := mapDocument["id"]
		if ok {
			err = stub.DelState(fmt.Sprintf("%v", id))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
