package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DataCategory struct {
	DocType          string         `json:"docType"`
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	DefinitionFormat string         `json:"definition_format"`
	Children         []DataCategory `json:"children"`
}

func CreateDataCategory(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, category DataCategory) pb.Response {
	logger.Info("entering-create-datacategory")
	defer logger.Info("exiting-create-datacategory")

	err := category.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	// ==== Check if category already exists ====
	categoryBytes, err := stub.GetState(category.ID)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	} else if categoryBytes != nil {
		errorMsg := fmt.Sprintf("this-category-already-exists-%s", category.ID)
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	categoryJSONBytes, err := json.Marshal(category)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// === Save category to state ===
	err = stub.PutState(category.ID, categoryJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetDataCategory(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-datacategory")
	defer logger.Info("exiting-get-datacategory")
	categoryAsBytes, err := GetDataCategoryState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(categoryAsBytes)
}

func GetDataCategoryState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	logger.Info("entering-get-datacategoryState")
	defer logger.Info("exiting-get-datacategoryState")
	categoryAsbytes, err := stub.GetState(id) //get the category from chaincode state
	if err != nil {
		respMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", id)
		logger.Error(respMsg)
		return nil, fmt.Errorf(respMsg)
	} else if categoryAsbytes == nil {
		respMsg := fmt.Sprintf("error-category-does-not-exist-%s", id)
		logger.Error(respMsg)
		return nil, fmt.Errorf(respMsg)
	}

	category := DataCategory{}
	err = json.Unmarshal(categoryAsbytes, &category)
	if err != nil {
		respMsg := fmt.Sprintf("error-unmarshaling-category-%s", id)
		logger.Error(respMsg)
		return nil, fmt.Errorf(respMsg)
	}

	err = category.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	categoryAsbytes, err = json.Marshal(category)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return categoryAsbytes, nil
}

func (d *DataCategory) checkAttributes(logger *shim.ChaincodeLogger) error {
	logger.Info("entering-checkAttributes-dataCategory")
	defer logger.Info("exiting-checkAttributes-dataCategory")

	if d.DocType != DATA_CATEGORY_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", d.DocType, DATA_CATEGORY_DOCTYPE)
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	return nil
}
