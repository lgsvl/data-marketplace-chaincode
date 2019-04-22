//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Person struct {
	DocType    string `json:"docType"`
	ID         string `json:"id"`
	BusinessID string `json:"businessID"`
	Role       Role   `json:"role"`
}

func AddPerson(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, person Person) pb.Response {
	logger.Info("entering-create-person")
	defer logger.Info("exiting-create-person")

	// ==== Check person attributes
	business, err := person.checkAndGetAttributes(logger, stub)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// ==== Check if person already exists ====
	personBytes, err := stub.GetState(person.ID)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to get person: " + err.Error())
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	} else if personBytes != nil {
		errorMsg := fmt.Sprintf("This person already exists: " + person.ID)
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	personJSONBytes, err := json.Marshal(person)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// === Save person to state ===
	err = stub.PutState(person.ID, personJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	err = business.AddPerson(logger, stub, person)
	if err != nil {
		logger.Error(err.Error())
		stub.DelState(person.ID)
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetPerson(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-person")
	defer logger.Info("exiting-get-person")
	person, err := GetPersonState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(person)
}

func GetPersonState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	logger.Info("entering-get-personState")
	defer logger.Info("exiting-create-personState")

	personAsbytes, err := stub.GetState(id) //get the person from chaincode state
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		logger.Error(jsonResp)
		return nil, fmt.Errorf(jsonResp)
	} else if personAsbytes == nil {
		jsonResp := "{\"Error\":\"person does not exist: " + id + "\"}"
		logger.Error(jsonResp)
		return nil, fmt.Errorf(jsonResp)
	}
	person := Person{}
	err = json.Unmarshal(personAsbytes, &person)
	if err != nil {
		jsonResp := "{\"Error\":\"unmarshalling: " + id + "\"}"
		logger.Error(jsonResp)
		return nil, fmt.Errorf(jsonResp)
	}

	_, err = person.checkAndGetAttributes(logger, stub)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	personAsbytes, err = json.Marshal(person)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return personAsbytes, nil
}

func (p *Person) checkAndGetAttributes(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) (Business, error) {
	logger.Info("entering-checkAttributes-person")
	defer logger.Info("exiting-checkAttributes-person")

	if p.DocType != PERSON_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", p.DocType, PERSON_DOCTYPE)
		logger.Error(errorMsg)
		return Business{}, fmt.Errorf(errorMsg)
	}

	businessAsBytes, err := GetBusinessState(logger, stub, p.BusinessID)
	if err != nil {
		logger.Error(err.Error())
		return Business{}, err
	}

	business := Business{}
	err = json.Unmarshal(businessAsBytes, &business)
	if err != nil {
		logger.Error(err.Error())
		return Business{}, err
	}
	return business, nil
}
