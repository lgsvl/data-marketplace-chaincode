package resources

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Business struct {
	DocType           string   `json:"docType"`
	ID                string   `json:"id"`
	EmailDomain       string   `json:"emailDomain"`
	OpenIDProviderURL string   `json:"openIdProviderURL"`
	Name              string   `json:"name"`
	Address           Address  `json:"address"`
	AccountBalance    float64  `json:"accountBalance"`
	PublicKey         string   `json:"publicKey"`
	Reviews           []Review `json:"reviews"`
	Score             float32  `json:"score"`
	NumberOfReviews   int      `json:"numberOfReviews"`
	Persons           []Person `json:"persons"`
}

func CreateBusiness(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, business Business) pb.Response {
	logger.Info("entering-create-business")
	defer logger.Info("exiting-create-business")
	business.Score = 3
	business.NumberOfReviews = 0
	// ==== Check business attributes
	err := business.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// ==== Check if business already exists ====
	businessBytes, err := stub.GetState(business.ID)
	if err != nil {
		errorMsg := fmt.Sprintf("failed-to-get-business-%s", err.Error())
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	} else if businessBytes != nil {
		errorMsg := fmt.Sprintf("this-business-already-exists-%s", business.ID)
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	businessJSONBytes, err := json.Marshal(business)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	// === Save business to state ===
	err = stub.PutState(business.ID, businessJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func GetBusiness(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) pb.Response {
	logger.Info("entering-get-business")
	defer logger.Info("exiting-get-business")
	business, err := GetBusinessState(logger, stub, id)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(business)
}

func GetBusinessState(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	logger.Info("entering-get-businessState")
	defer logger.Info("exiting-get-businessState")

	businessAsbytes, err := stub.GetState(id) //get the business from chaincode state
	if err != nil {
		respMsg := fmt.Sprintf("error-failed-to-get-state-for-%s", id)
		logger.Error(respMsg)
		return nil, fmt.Errorf(respMsg)
	} else if businessAsbytes == nil {
		respMsg := fmt.Sprintf("error-business-does-not-exist-%s", id)
		logger.Error(respMsg)
		return nil, fmt.Errorf(respMsg)
	}
	business := Business{}
	err = json.Unmarshal(businessAsbytes, &business)
	if err != nil {
		respMsg := fmt.Sprintf("error-unmarshalling-%s", id)
		logger.Error(respMsg)
		return nil, fmt.Errorf(respMsg)
	}

	err = business.checkAttributes(logger)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	businessAsbytes, err = json.Marshal(business)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return businessAsbytes, nil
}

func (b *Business) checkAttributes(logger *shim.ChaincodeLogger) error {
	logger.Info("entering-checkAttributes-business")
	defer logger.Info("exiting-checkAttributes-business")

	if b.DocType != BUSINESS_DOCTYPE {
		errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", b.DocType, BUSINESS_DOCTYPE)
		logger.Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	return nil
}

func (b *Business) AddReview(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, review Review) error {
	b.Score = (b.Score*float32(b.NumberOfReviews) + float32(review.Score)) / float32(b.NumberOfReviews+1)
	b.NumberOfReviews++
	b.Reviews = append(b.Reviews, review)
	businessJSONBytes, err := json.Marshal(b)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = stub.PutState(b.ID, businessJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (b *Business) AddPerson(logger *shim.ChaincodeLogger, stub shim.ChaincodeStubInterface, person Person) error {

	b.Persons = append(b.Persons, person)
	businessJSONBytes, err := json.Marshal(b)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = stub.PutState(b.ID, businessJSONBytes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
