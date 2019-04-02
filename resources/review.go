package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Review struct {
	DocType            string `json:"docType"`
	ID                 string `json:"id"`
	ReviewText         string `json:"reviewText"`
	Score              int    `json:"score"`
	DataContractID     string `json:"dataContract"`
	ReviewerID         string `json:"reviewer"`
	DataContractTypeID string `json:"dataContractType"`
}

func SubmitReview(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, review Review) pb.Response {
	logger.Info("entering-submit-review")
	defer logger.Info("exiting-submit-review")
	// === Check that data is accurate
	dataContractType, provider, err := review.checkAndGetAttributes(logger, stub)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	// ==== Check if review already exists ====
	reviewBytes, err := stub.GetState(review.ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + review.ID + "\"}"
		logger.Error(jsonResp)
		return shim.Error(jsonResp)
	} else if reviewBytes != nil {
		jsonResp := "{\"Error\":\"review already exists: " + review.ID + "\"}"
		logger.Error(jsonResp)
		return shim.Error(jsonResp)
	}

	reviewJSONBytes, err := json.Marshal(review)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// === Save review to state ===
	err = stub.PutState(review.ID, reviewJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	err = provider.AddReview(logger, stub, review)
	if err != nil {
		//TODO if anything fails undo
		//stub.DelState(review.ID)
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	err = dataContractType.AddReview(logger, stub, review)
	if err != nil {
		//provider.RemoveReview(logger, stub, review)
		//stub.DelState(review.ID)
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetReview(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-review")
	defer logger.Info("exiting-get-review")

	reviewAsBytes, err := GetReviewState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(reviewAsBytes)
}

func GetReviewState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	logger.Info("entering-get-review-state")
	defer logger.Info("exiting-get-review-state")
	reviewAsbytes, err := stub.GetState(id) //get the reviewAsbytes from chaincode state
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		logger.Error(jsonResp)
		return nil, fmt.Errorf(jsonResp)
	} else if reviewAsbytes == nil {
		jsonResp := "{\"Error\":\"review does not exist: " + id + "\"}"
		logger.Error(jsonResp)
		return nil, fmt.Errorf(jsonResp)
	}
	review := Review{}
	err = json.Unmarshal(reviewAsbytes, &review)
	if err != nil {
		jsonResp := "{\"Error\":\"unmarshaling: " + id + "\"}"
		logger.Error(jsonResp)
		return nil, fmt.Errorf(jsonResp)
	}

	if review.DocType != REVIEW_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", review.DocType, REVIEW_DOCTYPE)
		logger.Error(errorMsg)
		return nil, fmt.Errorf(errorMsg)
	}

	reviewAsbytes, err = json.Marshal(review)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return reviewAsbytes, nil
}

func (r *Review) checkAndGetAttributes(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) (DataContractType, Business, error) {
	logger.Info("entering-checkAttributes-review")
	defer logger.Info("exiting-checkAttributes-review")

	if r.DocType != REVIEW_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", r.DocType, REVIEW_DOCTYPE)
		logger.Error(errorMsg)
		return DataContractType{}, Business{}, fmt.Errorf(errorMsg)
	}
	if r.Score < 0 || r.Score > 5 {
		errorMsg := "error-score-should-be-between-0-and-5"
		logger.Error(errorMsg)
		return DataContractType{}, Business{}, fmt.Errorf(errorMsg)

	}

	_, err := GetBusinessState(logger, stub, r.ReviewerID)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}

	dataContractAsBytes, err := GetDataContractState(logger, stub, r.DataContractID)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}
	dataContract := DataContract{}
	err = json.Unmarshal(dataContractAsBytes, &dataContract)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}

	dataContractTypeAsBytes, err := GetDataContractTypeState(logger, stub, dataContract.DataContractTypeID)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}

	dataContractType := DataContractType{}
	err = json.Unmarshal(dataContractTypeAsBytes, &dataContractType)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}

	providerAsBytes, err := GetBusinessState(logger, stub, dataContractType.ProviderID)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}

	provider := Business{}
	err = json.Unmarshal(providerAsBytes, &provider)
	if err != nil {
		logger.Error(err.Error())
		return DataContractType{}, Business{}, err
	}
	return dataContractType, provider, nil
}
