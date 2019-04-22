//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package resources

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DataContractType struct {
	DocType          string                 `json:"docType"`
	ID               string                 `json:"id"`
	ThumbnailURL     string                 `json:"thumbnailURL"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Status           DataContractTypeStatus `json:"dataContractTypeStatus"`
	CreationDateTime time.Time              `json:"creationDateTime"`
	CategoryID       string                 `json:"categoryId"`
	DataType         DataType               `json:"dataType"`
	PriceType        PriceType              `json:"priceType"`
	Ownership        Ownership              `json:"ownership"`
	ProviderID       string                 `json:"provider"`
	Extras           ContractTypeExtras     `json:"extras"`
	DefinitionFormat string                 `json:"definition_format"`
	Reviews          []Review               `json:"reviews"`
	Score            float32                `json:"score"`
	NumberOfReviews  int                    `json:"numberOfReviews"`
}

func CreateDataContractType(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, dataContractType DataContractType) pb.Response {
	logger.Info("entering-create-dataContractType")
	defer logger.Info("exiting-create-dataContractType")

	dataContractType.Score = 3
	dataContractType.NumberOfReviews = 0

	// === Check that data is accurate
	err := dataContractType.checkAttributes(logger, stub)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	// ==== Check if dataContractType already exists ====
	dataContractTypeBytes, err := stub.GetState(dataContractType.ID)
	if err != nil {
		errMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", dataContractType.ID)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	} else if dataContractTypeBytes != nil {
		errMsg := fmt.Sprintf("error-dataContractType-already-exists-%s", dataContractType.ID)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	dataContractTypeJSONBytes, err := json.Marshal(dataContractType)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// === Save dataContractType to state ===
	err = stub.PutState(dataContractType.ID, dataContractTypeJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetDataContractType(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-dataContractType")
	defer logger.Info("exiting-get-dataContractType")

	dataContractTypeAsBytes, err := GetDataContractTypeState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(dataContractTypeAsBytes)
}

func GetDataContractTypeState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	logger.Info("entering-get-dataContractType-state")
	defer logger.Info("exiting-get-dataContractType-state")

	dataContractType, err := GetDataContractTypeStructState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	dataContractTypeAsbytes, err := json.Marshal(dataContractType)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return dataContractTypeAsbytes, nil
}

func GetDataContractTypeStructState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) (DataContractType, error) {
	logger.Info("entering-get-dataContractType-state")
	defer logger.Info("exiting-get-dataContractType-state")

	dataContractTypeAsbytes, err := stub.GetState(id) //get the dataContractTypeJSONBytes from chaincode state
	if err != nil {
		errMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", id)
		logger.Error(errMsg)
		return DataContractType{}, fmt.Errorf(errMsg)
	} else if dataContractTypeAsbytes == nil {
		errMsg := fmt.Sprintf("error-dataContractType-does-not-exist-%s", id)
		logger.Error(errMsg)
		return DataContractType{}, fmt.Errorf(errMsg)
	}
	dataContractType := DataContractType{}
	err = json.Unmarshal(dataContractTypeAsbytes, &dataContractType)
	if err != nil {
		errMsg := fmt.Sprintf("error-unmarshaling-%s", id)
		logger.Error(errMsg)
		return DataContractType{}, fmt.Errorf(errMsg)
	}

	if dataContractType.DocType != DATA_CONTRACT_TYPE_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", dataContractType.DocType, DATA_CONTRACT_TYPE_DOCTYPE)
		logger.Error(errorMsg)
		return DataContractType{}, fmt.Errorf(errorMsg)
	}
	return dataContractType, nil
}

func (d *DataContractType) checkAttributes(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) error {
	logger.Info("entering-checkAttributes-dataContractType")
	defer logger.Info("exiting-checkAttributes-dataContractType")

	if d.DocType != DATA_CONTRACT_TYPE_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", d.DocType, DATA_CONTRACT_TYPE_DOCTYPE)
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	if d.DataType == STREAM {
		if d.Extras.StreamType == PULL {
			if d.Extras.StreamSourceEndpoint == "" {
				errorMsg := fmt.Sprintf("error-stream-source-endpoint-is-required")
				logger.Error(errorMsg)
				return fmt.Errorf(errorMsg)
			}
		}

	}

	_, err := GetDataCategoryState(logger, stub, d.CategoryID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = GetBusinessState(logger, stub, d.ProviderID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (d *DataContractType) AddReview(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, review Review) error {
	d.Score = (d.Score*float32(d.NumberOfReviews) + float32(review.Score)) / float32(d.NumberOfReviews+1)
	d.NumberOfReviews++
	d.Reviews = append(d.Reviews, review)

	dataContractTypeJSONBytes, err := json.Marshal(d)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = stub.PutState(d.ID, dataContractTypeJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
