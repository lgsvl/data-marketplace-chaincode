package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Token struct {
	DocType           string  `json:"docType"`
	ID                string  `json:"id"`
	TotalTokensSupply float64 `json:"totalTokensSupply"`
	RemainingSupply   float64 `json:"remainingSupply"`
}

func NewToken(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) (Token, error) {
	return NewTokenWithSupply(logger, stub, id, 0.0)
}

func NewTokenWithSupply(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string, ts float64) (Token, error) {
	if ts < 0 {
		return Token{}, fmt.Errorf("error-token-totalTokensSupply-should-be-pausitive")
	}
	// ==== Check if token already exists ====
	tokenBytes, err := stub.GetState(id)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-check-token-%s", err.Error())
		logger.Error(errorMsg)
		return Token{}, fmt.Errorf(errorMsg)
	} else if tokenBytes != nil {
		errorMsg := fmt.Sprintf("this-token-already-exists-%s", id)
		logger.Error(errorMsg)
		return Token{}, fmt.Errorf(errorMsg)
	}

	token := Token{
		DocType:           TOKEN_DOCTYPE,
		ID:                id,
		TotalTokensSupply: ts,
		RemainingSupply:   ts,
	}
	err = SetTokenState(logger, stub, token)
	if err != nil {
		logger.Error(err.Error())
		return Token{}, err
	}

	return token, nil
}

func (t *Token) TotalSupply(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) (float64, error) {
	logger.Info("entering-token-TotalSupply")
	defer logger.Info("exiting-token-TotalSupply")
	token, err := GetTokenState(logger, stub, t.ID)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	return token.TotalTokensSupply, nil
}

func (t *Token) AvailableSupply(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface) (float64, error) {
	logger.Info("entering-token-AvailableSupply")
	defer logger.Info("exiting-token-AvailableSupply")
	token, err := GetTokenState(logger, stub, t.ID)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	return token.RemainingSupply, nil
}

func (t *Token) SetAccountBalance(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, account Account, tokens float64) (bool, error) {
	logger.Info("entering-token-SetAccountBalance")
	defer logger.Info("exiting-token-SetAccountBalance")
	token, err := GetTokenState(logger, stub, t.ID)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	if tokens > token.RemainingSupply {
		return false, fmt.Errorf("error-transfer-amount-greater-then-remaining-supply")
	}

	account, err = GetAccountState(logger, stub, account.ID)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	_, err = account.SetBalance(logger, stub, tokens)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	token.RemainingSupply = token.RemainingSupply - tokens

	err = SetTokenState(logger, stub, token)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	return true, nil
}

func GetTokenState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) (Token, error) {
	logger.Info("entering-get-tokenState")
	defer logger.Info("exiting-get-tokenState")

	tokenAsbytes, err := stub.GetState(id)
	if err != nil {
		respMsg := fmt.Sprintf("error-failed-to-get-token-state-for-%s", id)
		logger.Error(respMsg)
		return Token{}, fmt.Errorf(respMsg)
	} else if tokenAsbytes == nil {
		respMsg := fmt.Sprintf("error-token-does-not-exist-%s", id)
		logger.Error(respMsg)
		return Token{}, fmt.Errorf(respMsg)
	}
	token := Token{}
	err = json.Unmarshal(tokenAsbytes, &token)
	if err != nil {
		respMsg := fmt.Sprintf("error-unmarshalling-%s", id)
		logger.Error(respMsg)
		return Token{}, fmt.Errorf(respMsg)
	}

	err = token.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return Token{}, err
	}
	return token, nil
}

func SetTokenState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, token Token) error {
	logger.Info("entering-set-tokenState")
	defer logger.Info("exiting-set-tokenState")
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-marshal-token-%s", err.Error())
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	err = stub.PutState(token.ID, tokenBytes)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-update-token-%s", err.Error())
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}
	return nil
}

func (t *Token) checkAttributes(logger *shim.ChaincodeLogger) error {
	logger.Info("entering-checkAttributes-token")
	defer logger.Info("exiting-checkAttributes-token")

	if t.DocType != TOKEN_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", t.DocType, TOKEN_DOCTYPE)
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	return nil
}
