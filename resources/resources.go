package resources

import "time"

//Address this is the address for users
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Street  string `json:"street"`
	Zip     string `json:"zip"`
}

type Hash struct {
	Method string `json:"method"`
	Value  string `json:"value"`
}
type OwnershipType string
type DataContractStatus string
type StreamContractStatus string
type StreamType string
type DataType string
type DataContractTypeStatus string
type Role string

const (
	SHARED                        OwnershipType          = "SHARED"
	HOLD_BY_SELLER                OwnershipType          = "HOLD_BY_SELLER"
	TRANSFERRED_TO_BUYER          OwnershipType          = "TRANSFERRED_TO_BUYER"
	UNDEFINED                     OwnershipType          = "undefined"
	PROPOSAL                      DataContractStatus     = "PROPOSAL"
	DATASHIPPED                   DataContractStatus     = "DATASHIPPED"
	DATARECEIVED                  DataContractStatus     = "DATARECEIVED"
	CURRENT                       StreamContractStatus   = "CURRENT"
	PAST                          StreamContractStatus   = "PAST"
	PULL                          StreamType             = "PULL"
	PUSH                          StreamType             = "PUSH"
	NOTIFICATION                  StreamType             = "NOTIFICATION"
	FILE                          DataType               = "FILE"
	STREAM                        DataType               = "STREAM"
	FILEQUERYABLE                 DataType               = "FILEQUERYABLE"
	STREAMQUERYABLE               DataType               = "STREAMQUERYABLE"
	ACTIVE                        DataContractTypeStatus = "ACTIVE"
	INACTIVE                      DataContractTypeStatus = "INACTIVE"
	CONSORITUM_ADMIN              Role                   = "CONSORTIUM_ADMIN"
	GROUP_ADMIN                   Role                   = "GROUP_ADMIN"
	DATASET_ADMIN                 Role                   = "DATASET_ADMIN"
	USER                          Role                   = "USER"
	BUSINESS_DOCTYPE              string                 = "com.lge.svl.datamarketplace.contract.Business"
	DATA_CATEGORY_DOCTYPE         string                 = "com.lge.svl.datamarketplace.contract.DataCategory"
	DATA_CONTRACT_DOCTYPE         string                 = "com.lge.svl.datamarketplace.contract.DataContract"
	REVIEW_DOCTYPE                string                 = "com.lge.svl.datamarketplace.contract.Review"
	DATA_CONTRACT_TYPE_DOCTYPE    string                 = "com.lge.svl.datamarketplace.contract.DataContractType"
	SUBMIT_DATA_CONTRACT_PROPOSAL string                 = "com.lge.svl.datamarketplace.contract.SubmitDataContractProposal"
	PERSON_DOCTYPE                string                 = "com.lge.svl.datamarketplace.contract.Person"
	//TOKEN related
	TOKEN_DOCTYPE             string = "com.lge.svl.datamarketplace.contract.Token"
	ACCOUNT_DOCTYPE           string = "com.lge.svl.datamarketplace.contract.Account"
	ALLOWANCE_DOCTYPE         string = "com.lge.svl.datamarketplace.contract.Allowance"
	SET_ACCOUNT_BALANCE_EVENT string = "SetAccountBalance"
	TRANSFER_EVENT            string = "Transfer"
	TRANSFER_FROM_EVENT       string = "TransferFrom"
)

type Transfer struct {
	From  string  `json:""from`
	To    string  `json:"to"`
	Value float64 `json:"value"`
}

type OwnershipRevocation struct {
	RevocationTime string `json:"revocationTime"`
	RefundPolicy   string `json:"refundPolicy"`
}

//OwnershipVerificationMethod todo
type OwnershipVerificationMethod struct {
	AttributeName string `json:"attributeName"`
	Hash          Hash   `json:"hash"`
}

//Ownership describes the data ownership
type Ownership struct {
	OwnershipType               OwnershipType               `json:"ownershipType"`
	OwnershipRevocation         OwnershipRevocation         `json:"revocation"`
	OwnershipVerificationMethod OwnershipVerificationMethod `json:"ownershipVerificationMethod"`
}

type PriceType struct {
	DefinitionFormat string  `json:"definition_format"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
}

type ContractTypeExtras struct {
	DefinitionFormat     string     `json:"definition_format"`
	Frequency            int        `json:"frequency"`
	StreamType           StreamType `json:"streamType"`
	StreamSourceEndpoint string     `json:"streamSourceEndPoint"`
	StreamTargetEndpoint string     `json:"streamTargetEndPoint"`
	StreamTopic          string     `json:"streamTopic"`
	StartTime            time.Time  `json:"startTime"`
	EndTime              time.Time  `json:"endTime"`
	Hash                 Hash       `json:"hash"`
}

type ContractExtras struct {
	DefinitionFormat string             `json:"definition_format"`
	FileHash         Hash               `json:"fileHash"`
	FileStatus       DataContractStatus `json:"fileStatus"`
	EndDateTime      time.Time          `json:"endDateTime"`
}
