package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Account struct {
	DocType         string             `json:"docType"`
	ID              string             `json:"id"`
	Owner           string             `json:"owner"`
	Balance         float64            `json:"balance"`
	TotalAllowances float64            `json:"totalAllowances"`
	Allowances      map[string]float64 `json:"allowances"`
}

func CreateAccount(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, account Account) pb.Response {
	logger.Info("entering-create-account")
	defer logger.Info("exiting-create-account")

	// ==== Check account attributes
	err := account.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	// ==== Check if account already exists ====
	accountBytes, err := stub.GetState(account.ID)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-get-account-%s", err.Error())
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	} else if accountBytes != nil {
		errorMsg := fmt.Sprintf("this-account-already-exists-%s", account.ID)
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	account.Balance = 0.0
	account.TotalAllowances = 0.0
	err = SetAccountState(logger, stub, account)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	accountBytes, err = json.Marshal(account)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(accountBytes)
}

func GetAccount(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-account")
	defer logger.Info("exiting-get-account")
	account, err := GetAccountState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	accountAsbytes, err := json.Marshal(account)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(accountAsbytes)
}

func GetAccountState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) (Account, error) {
	logger.Info("entering-get-accountState")
	defer logger.Info("exiting-get-accountState")

	accountAsbytes, err := stub.GetState(id)
	if err != nil {
		respMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", id)
		logger.Error(respMsg)
		return Account{}, fmt.Errorf(respMsg)
	} else if accountAsbytes == nil {
		respMsg := fmt.Sprintf("error-account-does-not-exist-%s", id)
		logger.Error(respMsg)
		return Account{}, fmt.Errorf(respMsg)
	}
	account := Account{}
	err = json.Unmarshal(accountAsbytes, &account)
	if err != nil {
		respMsg := fmt.Sprintf("error-unmarshalling-%s", id)
		logger.Error(respMsg)
		return Account{}, fmt.Errorf(respMsg)
	}

	err = account.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return Account{}, err
	}

	return account, nil
}

func SetAccountState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, account Account) error {
	logger.Info("entering-set-accountState")
	defer logger.Info("exiting-set-accountState")
	accountBytes, err := json.Marshal(account)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-marshal-account-%s", err.Error())
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	err = stub.PutState(account.ID, accountBytes)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-update-account-%s", err.Error())
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}
	return nil
}

func (a *Account) checkAttributes(logger *shim.ChaincodeLogger) error {
	logger.Info("entering-checkAttributes-account")
	defer logger.Info("exiting-checkAttributes-account")

	if a.DocType != ACCOUNT_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", a.DocType, ACCOUNT_DOCTYPE)
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	return nil
}

func BalanceOf(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, account Account) (float64, error) {
	logger.Info("entering-token-BalanceOf")
	defer logger.Info("exiting-token-BalanceOf")

	account, err := GetAccountState(logger, stub, account.ID)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}
	return account.Balance, nil
}

func Approve(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, ownerID string, spenderID string, tokens float64) (bool, error) {
	logger.Info("entering-token-Approve")
	defer logger.Info("exiting-token-Approve")

	owner, err := GetAccountState(logger, stub, ownerID)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	if owner.Balance < owner.TotalAllowances+tokens {
		errorMsg := ("no-enough-funds-to-fulfill-allowances")
		logger.Error(errorMsg)
		return false, fmt.Errorf(errorMsg)
	}

	_, err = GetAccountState(logger, stub, spenderID)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	if owner.Allowances == nil {
		owner.Allowances = map[string]float64{}
	}

	if val, ok := owner.Allowances[spenderID]; !ok {
		owner.TotalAllowances = owner.TotalAllowances + tokens
	} else {
		owner.TotalAllowances = owner.TotalAllowances - val + tokens
	}

	owner.Allowances[spenderID] = tokens

	err = SetAccountState(logger, stub, owner)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	return true, nil
}

func Allowance(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, owner Account, spender Account) (float64, error) {
	logger.Info("entering-token-Allowance")
	defer logger.Info("exiting-token-Allowance")

	owner, err := GetAccountState(logger, stub, owner.ID)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	if allowance, ok := owner.Allowances[spender.ID]; ok {
		return allowance, nil
	}

	return 0, nil
}
func (a *Account) Transfer(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, to Account, tokens float64) (bool, error) {
	logger.Info("entering-token-Transfer")
	defer logger.Info("exiting-token-Transfer")
	return false, fmt.Errorf("method-not-implemented")
}

func TransferFrom(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, sourceID string, destinationID string, tokens float64) (bool, error) {
	logger.Info("entering-token-TransferFrom")
	defer logger.Info("exiting-token-TransferFrom")

	source, err := GetAccountState(logger, stub, sourceID)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	destination, err := GetAccountState(logger, stub, destinationID)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	allowance, ok := source.Allowances[destinationID]
	if !ok {
		allowance = 0
	}

	if tokens > allowance {
		return false, fmt.Errorf("no-enough-allowance-to-transfer")
	}

	_, err = source.RetrieveFunds(logger, stub, tokens)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	_, err = source.ReduceAllowance(logger, stub, destinationID, tokens)
	if err != nil {
		source.AddFunds(logger, stub, tokens)

		logger.Error(err.Error())
		return false, err
	}

	_, err = destination.AddFunds(logger, stub, tokens)
	if err != nil {
		source.AddFunds(logger, stub, tokens)

		logger.Error(err.Error())
		return false, err
	}

	return true, nil
}

func (a *Account) ReduceAllowance(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, spenderID string, tokens float64) (float64, error) {
	logger.Info("entering-account-ReduceAllowance")
	defer logger.Info("exiting-account-ReduceAllowance")

	allowance, ok := a.Allowances[spenderID]
	if !ok {
		return 0, nil
	}
	initialAllowance := allowance
	allowance -= tokens
	a.TotalAllowances -= tokens

	if allowance == 0 {
		delete(a.Allowances, spenderID)
	} else {
		a.Allowances[spenderID] = allowance
	}

	err := SetAccountState(logger, stub, *a)
	if err != nil {
		logger.Error(err.Error())
		return initialAllowance, err
	}

	return allowance, nil
}

func (a *Account) SetBalance(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, tokens float64) (float64, error) {
	logger.Info("entering-account-SetBalance")
	defer logger.Info("exiting-account-SetBalance")
	oldBalance := a.Balance
	a.Balance = tokens

	err := SetAccountState(logger, stub, *a)
	if err != nil {
		logger.Error(err.Error())
		return oldBalance, err
	}

	return a.Balance, nil
}

func (a *Account) AddFunds(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, tokens float64) (float64, error) {
	logger.Info("entering-account-AddAllowance")
	defer logger.Info("exiting-account-AddAllowance")

	oldBalance := a.Balance
	a.Balance += tokens

	err := SetAccountState(logger, stub, *a)
	if err != nil {
		logger.Error(err.Error())
		return oldBalance, err
	}

	return a.Balance, nil
}

func (a *Account) RetrieveFunds(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, tokens float64) (float64, error) {
	logger.Info("entering-account-RetrieveFunds")
	defer logger.Info("exiting-account-RetrieveFunds")
	oldBalance := a.Balance
	if tokens > a.Balance {
		return a.Balance, fmt.Errorf("no-enough-funds-to-retrieve")
	}

	a.Balance -= tokens
	err := SetAccountState(logger, stub, *a)
	if err != nil {
		logger.Error(err.Error())
		return oldBalance, err
	}

	return a.Balance, nil
}
