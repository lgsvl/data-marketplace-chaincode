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
	"github.com/lgsvl/data-marketplace-chaincode/utils"
)

type DataContract struct {
	DocType            string         `json:"docType"`
	ID                 string         `json:"id"`
	ProviderID         string         `json:"provider"`
	ConsumerID         string         `json:"consumer"`
	CreationDateTime   time.Time      `json:"creationDateTime"`
	Extras             ContractExtras `json:"extras"`
	DataContractTypeID string         `json:"dataContractType"`
}

type DataContractProposal struct {
	DataContractID        string         `json:"dataContractId"`
	ConsumerID            string         `json:"consumer"`
	DataContractTypeID    string         `json:"dataContractType"`
	DataContractTimestamp time.Time      `json:"dataContractTimestamp"`
	Extras                ContractExtras `json:"extras"`
}

func SubmitDataContractProposal(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, dataContractProposal DataContractProposal) pb.Response {
	logger.Info("entering-create-dataContract")
	defer logger.Info("exiting-create-dataContract")

	// === Check that data is accurate
	dataContractType, err := dataContractProposal.checkAndGetAttributes(logger, stub)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	// === approve account transfer
	consumerAccount, err := utils.GetAccountIDFromToken(logger, fmt.Sprintf("%s-%s", ACCOUNT_DOCTYPE, dataContractProposal.ConsumerID))
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	providerAccount, err := utils.GetAccountIDFromToken(logger, fmt.Sprintf("%s-%s", ACCOUNT_DOCTYPE, dataContractType.ProviderID))
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	approval, err := Approve(logger, stub, consumerAccount, providerAccount, dataContractType.PriceType.Amount)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	if !approval {
		errMsg := "payment-not-approved"
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}
	// ==== Check if dataContract already exists ====
	dataContractBytes, err := stub.GetState(dataContractProposal.DataContractID)
	if err != nil {
		errMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", dataContractProposal.DataContractID)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	} else if dataContractBytes != nil {
		errMsg := fmt.Sprintf("error-resource-already-exists-%s", dataContractProposal.DataContractID)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	if dataContractProposal.DataContractTimestamp.After(dataContractType.Extras.EndTime) || dataContractProposal.DataContractTimestamp.Before(dataContractType.Extras.StartTime) {
		errMsg := fmt.Sprintf("error-contract-creation-time-should-be-between-%s-and-%s", dataContractType.Extras.StartTime, dataContractType.Extras.EndTime)
		logger.Error(errMsg)
		return shim.Error(errMsg)
	}

	if dataContractType.DataType == "STREAM" {
		if dataContractProposal.Extras.EndDateTime.After(dataContractType.Extras.EndTime) || dataContractProposal.Extras.EndDateTime.Before(dataContractType.Extras.StartTime) {
			errMsg := fmt.Sprintf("error-contract-EndTime-should-be-between-%s-and-%s", dataContractType.Extras.StartTime, dataContractType.Extras.EndTime)
			logger.Error(errMsg)
			return shim.Error(errMsg)
		}
	}
	var extras ContractExtras
	if dataContractType.DataType == FILE {
		extras = ContractExtras{
			FileStatus:  PROPOSAL,
			EndDateTime: dataContractProposal.Extras.EndDateTime,
		}
	} else {
		if dataContractType.DataType == STREAM {
			extras = ContractExtras{EndDateTime: dataContractProposal.Extras.EndDateTime}

		}
	}

	dataContract := DataContract{
		DocType:            DATA_CONTRACT_DOCTYPE,
		ID:                 dataContractProposal.DataContractID,
		DataContractTypeID: dataContractProposal.DataContractTypeID,
		ConsumerID:         dataContractProposal.ConsumerID,
		CreationDateTime:   dataContractProposal.DataContractTimestamp,
		ProviderID:         dataContractType.ProviderID,
		Extras:             extras,
	}

	dataContractJSONBytes, err := json.Marshal(dataContract)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// === Save dataContract to state ===
	err = stub.PutState(dataContract.ID, dataContractJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetDataContract(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-dataContract")
	defer logger.Info("exiting-get-dataContract")

	dataContractAsBytes, err := GetDataContractState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(dataContractAsBytes)
}

func GetDataContractState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	logger.Info("entering-get-dataContract-state")
	defer logger.Info("exiting-get-dataContract-state")
	dataContractAsBytes, err := stub.GetState(id) //get the dataContractBytes from chaincode state
	if err != nil {
		errMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", id)
		logger.Error(errMsg)
		return nil, fmt.Errorf(errMsg)
	} else if dataContractAsBytes == nil {
		errMsg := fmt.Sprintf("error-dataContract-does-not-exist-%s", id)
		logger.Error(errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	dataContract := DataContract{}
	err = json.Unmarshal(dataContractAsBytes, &dataContract)
	if err != nil {
		errMsg := fmt.Sprintf("error-unmarshaling-%s", id)
		logger.Error(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if dataContract.DocType != DATA_CONTRACT_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", dataContract.DocType, DATA_CONTRACT_DOCTYPE)
		logger.Error(errorMsg)
		return nil, fmt.Errorf(errorMsg)
	}

	dataContractAsBytes, err = json.Marshal(dataContract)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return dataContractAsBytes, nil
}

func (d *DataContract) checkAttributes(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) error {
	logger.Info("entering-checkAttributes-dataContract")
	defer logger.Info("exiting-checkAttributes-dataContract")

	if d.DocType != DATA_CONTRACT_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", d.DocType, DATA_CONTRACT_DOCTYPE)
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	_, err := GetBusinessState(logger, stub, d.ConsumerID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (d *DataContractProposal) checkAndGetAttributes(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) (DataContractType, error) {
	logger.Info("entering-checkAttributes-SubmitDataContractProposal")
	defer logger.Info("exiting-checkAttributes-SubmitDataContractProposal")

	_, err := GetBusinessState(logger, stub, d.ConsumerID)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, err
	}

	dataContractTypeBytes, err := GetDataContractTypeState(logger, stub, d.DataContractTypeID)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, err
	}

	dataContractType := DataContractType{}

	err = json.Unmarshal(dataContractTypeBytes, &dataContractType)
	if err != nil {
		errMsg := fmt.Sprintf("error-unmarshaling-%s", d.DataContractTypeID)
		logger.Error(errMsg)
		return DataContractType{}, err
	}
	return dataContractType, nil
}

func (d *DataContract) SetFileStatus(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, status DataContractStatus, hash Hash, setHash bool) error {
	logger.Info("entering-dataContract-setfilestatus")
	defer logger.Info("exiting-dataContract-setfilestatus")

	d.Extras.FileStatus = status
	if setHash {
		d.Extras.FileHash = hash
	}
	dataContractJSONBytes, err := json.Marshal(d)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = stub.PutState(d.ID, dataContractJSONBytes)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-update-dataContract-%s", err.Error())
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}
	return nil

}
