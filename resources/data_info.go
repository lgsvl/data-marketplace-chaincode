package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/lgsvl/data-marketplace-chaincode/utils"
)

type DataInfoSentToConsumer struct {
	Hash           Hash   `json:"hash"`
	DataContractID string `json:"dataContract"`
}

type DataReceivedByConsumer struct {
	DataContractID string `json:"dataContract"`
}

func SetDataInfoSentToConsumer(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, dataInfo DataInfoSentToConsumer) pb.Response {
	logger.Info("entering-SetDataInfoSentToConsumer")
	defer logger.Info("exiting-SetDataInfoSentToConsumer")

	// ==== Check data attributes
	dataContract, err := checkAndGetAttributes(logger, stub, dataInfo.DataContractID)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	// ==== Check if business already exists ====
	if dataContract.Extras.FileStatus != PROPOSAL {
		errorMsg := "data is either shipped or received or a stream"
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err = dataContract.SetFileStatus(logger, stub, DATASHIPPED, dataInfo.Hash, true)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	consumerID, err := utils.GetAccountIDFromToken(logger, fmt.Sprintf("%s-%s", ACCOUNT_DOCTYPE, dataContract.ConsumerID))
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	providerID, err := utils.GetAccountIDFromToken(logger, fmt.Sprintf("%s-%s", ACCOUNT_DOCTYPE, dataContract.ProviderID))
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	dataContarctType, err := GetDataContractTypeStructState(logger, stub, dataContract.DataContractTypeID)

	_, err = TransferFrom(logger, stub, consumerID, providerID, dataContarctType.PriceType.Amount)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func SetDataReceivedByConsumer(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, dataReceived DataReceivedByConsumer) pb.Response {
	logger.Info("entering-SetDataReceivedByConsumer")
	defer logger.Info("exiting-SetDataReceivedByConsumer")

	// ==== Check data attributes
	dataContract, err := checkAndGetAttributes(logger, stub, dataReceived.DataContractID)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	if dataContract.Extras.FileStatus != DATASHIPPED {
		errorMsg := "data is either proposal, received or a stream"
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err = dataContract.SetFileStatus(logger, stub, DATARECEIVED, Hash{}, false)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func checkAndGetAttributes(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, dataContractID string) (DataContract, error) {
	logger.Info("entering-checkAttribute-dataTransaction")
	defer logger.Info("exiting-checkAttributes-dataTransaction")

	dataContractAsBytes, err := GetDataContractState(logger, stub, dataContractID)
	if err != nil {
		logger.Error(err.Error())
		return DataContract{}, err
	}
	dataContract := DataContract{}
	err = json.Unmarshal(dataContractAsBytes, &dataContract)
	if err != nil {
		logger.Error(err.Error())
		return DataContract{}, err
	}

	return dataContract, nil
}
