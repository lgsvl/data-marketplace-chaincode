//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
	"github.com/lgsvl/data-marketplace-chaincode/utils"
)

// DataMarketplaceChaincode datamarketplace Chaincode implementation
type DataMarketplaceChaincode struct {
	logger *shim.ChaincodeLogger
}

func NewDataMarketplaceChaincode(l *shim.ChaincodeLogger) *DataMarketplaceChaincode {
	return &DataMarketplaceChaincode{logger: l}
}

func (d *DataMarketplaceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()
	fmt.Printf("args %#v\n", args)
	d.logger.Info(fmt.Sprintf("args: %#v", args))
	return d.createToken(stub, args)
}

func (d *DataMarketplaceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	// POST and GET Operations
	case "createBusiness":
		return d.createBusiness(stub, args)
	case "getBusiness":
		return d.getBusiness(stub, args)
	case "createDataCategory":
		return d.createCategory(stub, args)
	case "getDataCategory":
		return d.getCategory(stub, args)
	case "createDataContractType":
		return d.createDataContractType(stub, args)
	case "getDataContractType":
		return d.getDataContractType(stub, args)
	case "submitDataContractProposal":
		return d.submitDataContractProposal(stub, args)
	case "getDataContract":
		return d.getDataContract(stub, args)
	case "submitReview":
		return d.submitReview(stub, args)
	case "getReview":
		return d.getReview(stub, args)
	case "addPerson":
		return d.addPerson(stub, args)
	case "getPerson":
		return d.getPerson(stub, args)
	case "setDataInfoSentToConsumer":
		return d.setDataInfoSentToConsumer(stub, args)
	case "setDataReceivedByConsumer":
		return d.setDataReceivedByConsumer(stub, args)
		// QUERIES
	case "getBusinesses":
		return d.getBusinesses(stub, args)
	case "getBusinessesWithPagination":
		return d.getBusinessesWithPagination(stub, args)
	case "getDataCategories":
		return d.getDataCategories(stub, args)
	case "getDataCategoriesWithPagination":
		return d.getDataCategoriesWithPagination(stub, args)
	case "getPopularDataCategories":
		return d.getPopularDataCategories(stub, args)
	case "getDataContractTypes":
		return d.getDataContractTypes(stub, args)
	case "getDataContractTypesAfterTimeStamp":
		return d.getDataContractTypesAfterTimeStamp(stub, args)
	case "getRecommendedDataContractType":
		return d.getRecommendedDataContractType(stub, args)
	case "getDataContractTypesWithPagination":
		return d.getDataContractTypesWithPagination(stub, args)
	case "getPopularDataContractTypes":
		return d.getPopularDataContractTypes(stub, args)
	case "getDataContractTypesByCategory":
		return d.getDataContractTypesByCategory(stub, args)
	case "getDataContractTypesByCategoryWithPagination":
		return d.getDataContractTypesByCategoryWithPagination(stub, args)
	case "getDataContractTypesByProvider":
		return d.getDataContractTypesByProvider(stub, args)
	case "getDataContractTypesByProviderWithPagination":
		return d.getDataContractTypesByProviderWithPagination(stub, args)
	case "selectNumberOfBusinessDataSetsToUpload":
		return d.selectNumberOfBusinessDataSetsToUpload(stub, args)
	case "getDataContracts":
		return d.getDataContracts(stub, args)
	case "getDataContractsWithPagination":
		return d.getDataContractsWithPagination(stub, args)
	case "getDataContractsByProvider":
		return d.getDataContractsByProvider(stub, args)
	case "getDataContractsByProviderWithPagination":
		return d.getDataContractsByProviderWithPagination(stub, args)
	case "getDataContractsByConsumer":
		return d.getDataContractsByConsumer(stub, args)
	case "getDataContractsByConsumerWithPagination":
		return d.getDataContractsByConsumerWithPagination(stub, args)
	case "selectDataSetContractsToUpload":
		return d.selectDataSetContractsToUpload(stub, args)
	case "selectDataSetContractsToUploadWithPagination":
		return d.selectDataSetContractsToUploadWithPagination(stub, args)
	case "selectDataContractsByDataContractType":
		return d.selectDataContractsByDataContractType(stub, args)
	case "selectDataContractsByDataContractTypeWithPagination":
		return d.selectDataContractsByDataContractTypeWithPagination(stub, args)
	case "selectBusinessDataSetsToUpload":
		return d.selectBusinessDataSetsToUpload(stub, args)
	case "selectBusinessDataSetsToUploadWithPagination":
		return d.selectBusinessDataSetsToUploadWithPagination(stub, args)
	case "selectBusinessDataSetsToUploadByDataContractType":
		return d.selectBusinessDataSetsToUploadByDataContractType(stub, args)
	case "selectBusinessDataSetsToUploadByDataContractTypeWithPagination":
		return d.selectBusinessDataSetsToUploadByDataContractTypeWithPagination(stub, args)
	case "selectBusinessDataSetsSoldShippedNotDownloaded":
		return d.selectBusinessDataSetsSoldShippedNotDownloaded(stub, args)
	case "selectBusinessDataSetsSoldShippedNotDownloadedWithPagination":
		return d.selectBusinessDataSetsSoldShippedNotDownloadedWithPagination(stub, args)
	case "selectBusinessDataSetsSoldAndDownloaded":
		return d.selectBusinessDataSetsSoldAndDownloaded(stub, args)
	case "selectBusinessDataSetsSoldAndDownloadedWithPagination":
		return d.selectBusinessDataSetsSoldAndDownloadedWithPagination(stub, args)
	case "selectBusinessDataSetsPurchasedNotUploaded":
		return d.selectBusinessDataSetsPurchasedNotUploaded(stub, args)
	case "selectBusinessDataSetsPurchasedNotUploadedWithPagination":
		return d.selectBusinessDataSetsPurchasedNotUploadedWithPagination(stub, args)
	case "selectBusinessDataSetsPurchasedUploadedNotDownloaded":
		return d.selectBusinessDataSetsPurchasedUploadedNotDownloaded(stub, args)
	case "selectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination":
		return d.selectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination(stub, args)
	case "selectBusinessDataSetsPurchasedDownloaded":
		return d.selectBusinessDataSetsPurchasedDownloaded(stub, args)
	case "selectBusinessDataSetsPurchasedDownloadedWithPagination":
		return d.selectBusinessDataSetsPurchasedDownloadedWithPagination(stub, args)
	case "cleanUp":
		return d.cleanUp(stub, args)
	case "deleteDoc":
		return d.deleteDoc(stub, args)
		// Token
	case "createToken":
		return d.createToken(stub, args)
	case "createAccount":
		return d.createAccount(stub, args)
	case "setAccountBalance":
		return d.setAccountBalance(stub, args)
	case "totalSupply":
		return d.totalSupply(stub, args)
	case "availableSupply":
		return d.availableSupply(stub, args)
	case "balanceOf":
		return d.balanceOf(stub, args)
	case "allowances":
		return d.allowances(stub, args)
	case "transfer":
		return d.transfer(stub, args)
	case "approve":
		return d.approve(stub, args)
	case "transferFrom":
		return d.transferFrom(stub, args)
	default:
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}

}

//////####################################

func (d *DataMarketplaceChaincode) createToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		errorMsg := fmt.Sprintf("incorrect-number-of-arguments-expecting-2-got%d", len(args))
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	tokenTotalSupply, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	_, err = resources.NewTokenWithSupply(d.logger, stub, args[0], tokenTotalSupply)
	if err != nil {
		d.logger.Error(err.Error())
		if err.Error() != "this-token-already-exists-dmpoken" {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)

}
func (d *DataMarketplaceChaincode) createAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		d.logger.Error("incorrect-number-of-arguments-expecting-1")
		return shim.Error("incorrect-number-of-arguments-expecting-1")
	}
	account := resources.Account{}
	err := json.Unmarshal([]byte(args[0]), &account)
	if err != nil {
		errorMsg := "error-unmarshalling-business-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.CreateAccount(d.logger, stub, account)
}

func (d *DataMarketplaceChaincode) setAccountBalance(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		d.logger.Error("incorrect-number-of-arguments-expecting-3")
		return shim.Error("incorrect-number-of-arguments-expecting-3")
	}
	token := resources.Token{ID: args[0], DocType: resources.TOKEN_DOCTYPE}
	account := resources.Account{ID: args[1], DocType: resources.ACCOUNT_DOCTYPE}
	tokens, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		errorMsg := fmt.Sprintf("error-parsing-tokens-%s", args[2])
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	_, err = token.SetAccountBalance(d.logger, stub, account, tokens)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	transfer := resources.Transfer{}

	transfer.To = account.ID
	transfer.Value = tokens
	evtData, err := json.Marshal(transfer)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	stub.SetEvent("SetAccountBalance", evtData)
	return shim.Success(nil)
}

func (d *DataMarketplaceChaincode) totalSupply(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		d.logger.Error("incorrect-number-of-arguments-expecting-1")
		return shim.Error("incorrect-number-of-arguments-expecting-1")
	}
	token := resources.Token{ID: args[0], DocType: resources.TOKEN_DOCTYPE}

	total, err := token.TotalSupply(d.logger, stub)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	resString := fmt.Sprintf("%f", total)
	return shim.Success([]byte(resString))
}

func (d *DataMarketplaceChaincode) availableSupply(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		d.logger.Error("incorrect-number-of-arguments-expecting-1")
		return shim.Error("incorrect-number-of-arguments-expecting-1")
	}
	token := resources.Token{ID: args[0], DocType: resources.TOKEN_DOCTYPE}

	available, err := token.AvailableSupply(d.logger, stub)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	resString := fmt.Sprintf("%f", available)
	return shim.Success([]byte(resString))
}

func (d *DataMarketplaceChaincode) balanceOf(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		d.logger.Error("incorrect-number-of-arguments-expecting-2")
		return shim.Error("incorrect-number-of-arguments-expecting-2")
	}

	account := resources.Account{DocType: resources.ACCOUNT_DOCTYPE, ID: args[1]}
	balance, err := resources.BalanceOf(d.logger, stub, account)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	resString := fmt.Sprintf("%f", balance)
	return shim.Success([]byte(resString))
}

func (d *DataMarketplaceChaincode) allowances(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		d.logger.Error("incorrect-number-of-arguments-expecting-3")
		return shim.Error("incorrect-number-of-arguments-expecting-3")
	}

	owner := resources.Account{ID: args[1], DocType: resources.ACCOUNT_DOCTYPE}
	spender := resources.Account{ID: args[2], DocType: resources.ACCOUNT_DOCTYPE}

	allowances, err := resources.Allowance(d.logger, stub, owner, spender)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	resString := fmt.Sprintf("%f", allowances)
	return shim.Success([]byte(resString))
}

func (d *DataMarketplaceChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("not-implemented"))
}

func (d *DataMarketplaceChaincode) approve(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	d.logger.Debugf("entering-chaincode-approve")
	defer d.logger.Debugf("exiting-chaincode-approve")
	if len(args) != 4 {
		d.logger.Error("incorrect-number-of-arguments-expecting-4")
		return shim.Error("incorrect-number-of-arguments-expecting-4")
	}

	tokens, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		errorMsg := fmt.Sprintf("error-parsing-tokens-%s", args[3])
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	d.logger.Debugf("approve-arguments-token-%s-owner-%s-spender-%s-amount-%f", args[0], args[1], args[2], tokens)

	_, err = resources.Approve(d.logger, stub, args[1], args[2], tokens)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (d *DataMarketplaceChaincode) transferFrom(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		d.logger.Error("incorrect-number-of-arguments-expecting-4")
		return shim.Error("incorrect-number-of-arguments-expecting-4")
	}

	tokens, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		errorMsg := fmt.Sprintf("error-parsing-tokens-%s", args[3])
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	_, err = resources.TransferFrom(d.logger, stub, args[1], args[2], tokens)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	transfer := resources.Transfer{}
	transfer.From = args[1]
	transfer.To = args[2]
	transfer.Value = tokens
	evtData, err := json.Marshal(transfer)
	if err != nil {
		d.logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	stub.SetEvent("TransferFrom", evtData)
	return shim.Success(nil)
}

//////####################################

func (d *DataMarketplaceChaincode) createBusiness(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		d.logger.Error("incorrect-number-of-arguments-expecting-1")
		return shim.Error("incorrect-number-of-arguments-expecting-1")
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	business := resources.Business{}
	err = json.Unmarshal([]byte(args[0]), &business)
	if err != nil {
		errorMsg := "error-unmarshalling-business-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.CreateBusiness(d.logger, stub, business)
}

func (d *DataMarketplaceChaincode) getBusiness(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetBusiness(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) createCategory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	category := resources.DataCategory{}
	err = json.Unmarshal([]byte(args[0]), &category)
	if err != nil {
		errorMsg := "error-unmarshalling-category-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.CreateDataCategory(d.logger, stub, category)
}

func (d *DataMarketplaceChaincode) getCategory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataCategory(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) createDataContractType(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	dataContractType := resources.DataContractType{}
	err = json.Unmarshal([]byte(args[0]), &dataContractType)
	if err != nil {
		errorMsg := "error-unmarshalling-dataContractType-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.CreateDataContractType(d.logger, stub, dataContractType)
}

func (d *DataMarketplaceChaincode) getDataContractType(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractType(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) submitDataContractProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	dataContractProposal := resources.DataContractProposal{}
	err = json.Unmarshal([]byte(args[0]), &dataContractProposal)
	if err != nil {
		errorMsg := "error-unmarshalling-submitDataContractProposal-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SubmitDataContractProposal(d.logger, stub, dataContractProposal)
}

func (d *DataMarketplaceChaincode) getDataContract(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContract(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) submitReview(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	review := resources.Review{}
	err = json.Unmarshal([]byte(args[0]), &review)
	if err != nil {
		errorMsg := "error-unmarshalling-review-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SubmitReview(d.logger, stub, review)
}

func (d *DataMarketplaceChaincode) getReview(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetReview(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) addPerson(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	person := resources.Person{}
	err = json.Unmarshal([]byte(args[0]), &person)
	if err != nil {
		errorMsg := "error-unmarshalling-person-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.AddPerson(d.logger, stub, person)
}

func (d *DataMarketplaceChaincode) getPerson(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetPerson(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) setDataInfoSentToConsumer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	dataInfoSentToConsumer := resources.DataInfoSentToConsumer{}
	err = json.Unmarshal([]byte(args[0]), &dataInfoSentToConsumer)
	if err != nil {
		errorMsg := "error-unmarshalling-dataInfoSentToConsumer-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SetDataInfoSentToConsumer(d.logger, stub, dataInfoSentToConsumer)
}

func (d *DataMarketplaceChaincode) setDataReceivedByConsumer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	dataReceivedByConsumer := resources.DataReceivedByConsumer{}
	err = json.Unmarshal([]byte(args[0]), &dataReceivedByConsumer)
	if err != nil {
		errorMsg := "error-unmarshalling-dataReceivedByConsumer-infos"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SetDataReceivedByConsumer(d.logger, stub, dataReceivedByConsumer)
}

// =========================================================================================
// Business related queries
// =========================================================================================

func (d *DataMarketplaceChaincode) getBusinesses(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive only the authorization token
	if len(args) != 1 {
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[0])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetBusinesses(d.logger, stub)
}

func (d *DataMarketplaceChaincode) getBusinessesWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[0])
	if err != nil {
		errorMsg := fmt.Sprintf("incorrect-page-size-format-%#v", args[0])
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetBusinessesWithPagination(d.logger, stub, int32(pageSize), args[1])
}

// =========================================================================================
// DataCategory related queries
// =========================================================================================

func (d *DataMarketplaceChaincode) getDataCategories(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive only the authorization token
	if len(args) != 1 {
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[0])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataCategories(d.logger, stub)
}

func (d *DataMarketplaceChaincode) getDataCategoriesWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[0])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataCategoriesWithPagination(d.logger, stub, int32(pageSize), args[1])
}

func (d *DataMarketplaceChaincode) getPopularDataCategories(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 0 or 1 arg and the authorization token
	switch len(args) {
	case 1:
		err := utils.CheckAuth(d.logger, args[0])
		if err != nil {
			errorMsg := "operation-not-authorized"
			d.logger.Error(errorMsg)
			return shim.Error(errorMsg)
		}
		return resources.GetPopularDataCategories(d.logger, stub, 8)
	case 2:
		err := utils.CheckAuth(d.logger, args[1])
		if err != nil {
			errorMsg := "operation-not-authorized"
			d.logger.Error(errorMsg)
			return shim.Error(errorMsg)
		}
		size, err := strconv.Atoi(args[0])
		if err != nil {
			errorMsg := "incorrect-size-format"
			d.logger.Error(errorMsg)
			return shim.Error(errorMsg)
		}
		return resources.GetPopularDataCategories(d.logger, stub, int32(size))
	default:
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
}

// =========================================================================================
// DataContractType related queries
// =========================================================================================

func (d *DataMarketplaceChaincode) getRecommendedDataContractType(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive only the authorization token
	if len(args) != 1 {
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[0])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetRecommendedDataContractType(d.logger, stub)
}

func (d *DataMarketplaceChaincode) getDataContractTypesAfterTimeStamp(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive only the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypesAfterTimeStamp(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) getDataContractTypes(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive only the authorization token
	if len(args) != 1 {
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[0])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypes(d.logger, stub)
}

func (d *DataMarketplaceChaincode) getDataContractTypesWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[0])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypesWithPagination(d.logger, stub, int32(pageSize), args[1])
}

func (d *DataMarketplaceChaincode) getPopularDataContractTypes(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 0 or 1 arg and the authorization token
	switch len(args) {
	case 1:
		err := utils.CheckAuth(d.logger, args[0])
		if err != nil {
			errorMsg := "operation-not-authorized"
			d.logger.Error(errorMsg)
			return shim.Error(errorMsg)
		}
		return resources.GetPopularDataContractTypes(d.logger, stub, 8)
	case 2:
		err := utils.CheckAuth(d.logger, args[1])
		if err != nil {
			errorMsg := "operation-not-authorized"
			d.logger.Error(errorMsg)
			return shim.Error(errorMsg)
		}
		size, err := strconv.Atoi(args[0])
		if err != nil {
			errorMsg := "incorrect-size-format"
			d.logger.Error(errorMsg)
			return shim.Error(errorMsg)
		}
		return resources.GetPopularDataContractTypes(d.logger, stub, int32(size))
	default:
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
}

func (d *DataMarketplaceChaincode) getDataContractTypesByCategory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypesByCategory(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) getDataContractTypesByCategoryWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypesByCategoryWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) getDataContractTypesByProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypesByProvider(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) getDataContractTypesByProviderWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractTypesByProviderWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) selectNumberOfBusinessDataSetsToUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectNumberOfBusinessDataSetsToUpload(d.logger, stub, args[0])
}

// =========================================================================================
// DataContract related queries
// =========================================================================================

func (d *DataMarketplaceChaincode) getDataContracts(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive only the authorization token
	if len(args) != 1 {
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[0])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContracts(d.logger, stub)
}

func (d *DataMarketplaceChaincode) getDataContractsWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[0])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractsWithPagination(d.logger, stub, int32(pageSize), args[1])
}

func (d *DataMarketplaceChaincode) getDataContractsByProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractsByProvider(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) getDataContractsByProviderWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractsByProviderWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) getDataContractsByConsumer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.GetDataContractsByConsumer(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) getDataContractsByConsumerWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.GetDataContractsByConsumerWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) selectDataSetContractsToUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectDataSetContractsToUpload(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) selectDataSetContractsToUploadWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectDataSetContractsToUploadWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsToUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsToUpload(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsToUploadWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectBusinessDataSetsToUploadWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) selectDataContractsByDataContractType(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectDataContractsByDataContractType(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) selectDataContractsByDataContractTypeWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectDataContractsByDataContractTypeWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}
func (d *DataMarketplaceChaincode) selectBusinessDataSetsToUploadByDataContractType(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsToUploadByContractType(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsToUploadByDataContractTypeWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectBusinessDataSetsToUploadByContractTypeWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}
func (d *DataMarketplaceChaincode) selectBusinessDataSetsSoldShippedNotDownloaded(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsSoldShippedNotDownloaded(d.logger, stub, args[0], args[1])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsSoldShippedNotDownloadedWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 4 args and the authorization token
	if len(args) != 5 {
		errorMsg := "incorrect-number-of-arguments-expecting-4"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[4])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[2])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsSoldShippedNotDownloadedWithPagination(d.logger, stub, args[0], args[1], int32(pageSize), args[3])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsSoldAndDownloaded(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsSoldAndDownloaded(d.logger, stub, args[0], args[1])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsSoldAndDownloadedWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 4 args and the authorization token
	if len(args) != 5 {
		errorMsg := "incorrect-number-of-arguments-expecting-4"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[4])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[2])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsSoldAndDownloadedWithPagination(d.logger, stub, args[0], args[1], int32(pageSize), args[3])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsPurchasedNotUploaded(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 1 arg and the authorization token
	if len(args) != 2 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[1])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsPurchasedNotUploaded(d.logger, stub, args[0])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsPurchasedNotUploadedWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-3"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectBusinessDataSetsPurchasedNotUploadedWithPagination(d.logger, stub, args[0], int32(pageSize), args[2])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsPurchasedUploadedNotDownloaded(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 2 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsPurchasedUploadedNotDownloaded(d.logger, stub, args[0], args[1])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-4"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[2])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination(d.logger, stub, args[0], args[1], int32(pageSize), args[3])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsPurchasedDownloaded(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 3 {
		errorMsg := "incorrect-number-of-arguments-expecting-2"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[2])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	return resources.SelectBusinessDataSetsPurchasedDownloaded(d.logger, stub, args[0], args[1])
}

func (d *DataMarketplaceChaincode) selectBusinessDataSetsPurchasedDownloadedWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// should receive 3 args and the authorization token
	if len(args) != 4 {
		errorMsg := "incorrect-number-of-arguments-expecting-4"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	err := utils.CheckAuth(d.logger, args[3])
	if err != nil {
		errorMsg := "operation-not-authorized"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	pageSize, err := strconv.Atoi(args[2])
	if err != nil {
		errorMsg := "incorrect-page-size-format"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.SelectBusinessDataSetsPurchasedDownloadedWithPagination(d.logger, stub, args[0], args[1], int32(pageSize), args[3])
}

// =========================================================================================
// cleanUp Function
// =========================================================================================

func (d *DataMarketplaceChaincode) cleanUp(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		errorMsg := "incorrect-number-of-arguments-expecting-0"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.CleanUp(d.logger, stub)
}

func (d *DataMarketplaceChaincode) deleteDoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		errorMsg := "incorrect-number-of-arguments-expecting-1"
		d.logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	return resources.DeleteDoc(d.logger, stub, args[0])
}

// =========================================================================================
// Main Function
// =========================================================================================

func main() {
	logger := shim.NewLogger("data-marketplace-chaincode-logger")

	err := shim.Start(NewDataMarketplaceChaincode(logger))
	if err != nil {
		logger.Error(err.Error())
		panic("error-starting-data-marketplace-chaincode:-" + err.Error())
	}

}

// func main() {

// 	err := utils.CheckAuth(nil, "eyJraWQiOiJjNmdCQWhydDBPMmplOTI2RWVqaFwvaHdVXC9ha2dhc2JOT3puVnR0OXdsc0k9IiwiYWxnIjoiUlMyNTYifQ.eyJhdF9oYXNoIjoiWV9menhtX3EtQTBUU1pSNzd4QXU3USIsInN1YiI6ImI1MGE0YjYyLTRmMmQtNDI3NC1iNzljLTdhMzA4MmEwMTllOSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtd2VzdC0yLmFtYXpvbmF3cy5jb21cL3VzLXdlc3QtMl9zdUdKeWNaNXciLCJjb2duaXRvOnVzZXJuYW1lIjoiamltIiwibm9uY2UiOiJmb29iYXJiYXoiLCJhdWQiOiI3Z21ucXAyNzIzNGFha25xdDRkMmd0MWI1ciIsImV2ZW50X2lkIjoiMDMyYjU2NjItY2NjNy0xMWU4LTljYzQtOTc1OTIxN2EwNTlkIiwidG9rZW5fdXNlIjoiaWQiLCJhdXRoX3RpbWUiOjE1MzkyMDE1NzYsImV4cCI6MTUzOTIwNTE3NiwiaWF0IjoxNTM5MjAxNTc2LCJlbWFpbCI6ImppbUBjb21wYW55Mi5jb20ifQ.ccada-wPb9loOHLuKqnms_hIhoFB-jvD4IcrmT1Y72XjjpT-T_rmSK7ya8ZBK86S5O3GHYo8a6tPPNSoOxjLeFJa_6EW54ZLFUY4mrlqyl1kLOpq5JFNSRUGPith_DpWaM38NKgnmeTBEAhixhAcCtMn0u7LjHJ34zLNrPWk95tcTMRXXo40Pb5uPZENGsouHC_kVxdcbjbSMBrI0GgKRo-WROY1HLsS4fb2MXI4tKUevOFCTn1Rx6Z0Gdz1wA4TeAyRYiXTVg5K6t11IjQ9cq9sRIkAnOzCvyiNKXFQOiPh-Fm8iqQPBkbk5wF3JwHMXmnCA0und-DhF0MPEpg7Qg")
// 	if err != nil {
// 		panic("error")
// 	}
// 	fmt.Printf("ok")
// }
