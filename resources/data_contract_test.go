package resources_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/lgsvl/data-marketplace-chaincode/fakes"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DataContract", func() {
	var (
		fakeStub                   *fakes.ChaincodeStub
		logger                     *shim.ChaincodeLogger
		dataContractType           resources.DataContractType
		dataContractProposal       resources.DataContractProposal
		dataContractTypeExtras     resources.ContractTypeExtras
		dataContractProposalExtras resources.ContractExtras
		consumerAccount            resources.Account
		providerAccount            resources.Account
		price                      resources.PriceType
		consumerAccountBytes       []byte
		providerAccountBytes       []byte
		err                        error
	)
	BeforeEach(func() {
		logger = shim.NewLogger("data-contract-test-logger")
		fakeStub = new(fakes.ChaincodeStub)
		consumerAccount = resources.Account{ID: "fake-consumer", Balance: 10, DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}
		providerAccount = resources.Account{ID: "fake-provider", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}
		price = resources.PriceType{
			Amount: 2,
		}
		consumerAccountBytes, err = json.Marshal(consumerAccount)
		Expect(err).NotTo(HaveOccurred())

		providerAccountBytes, err = json.Marshal(providerAccount)
		Expect(err).NotTo(HaveOccurred())
	})

	Context(".SubmitDataContractProposal", func() {

		It("should fail when stub getState for consumer returns an error", func() {
			dataContractProposal = resources.DataContractProposal{
				DataContractID: "fake-data-contract",
				ConsumerID:     "fake-consumer",
			}
			fakeStub.GetStateReturnsOnCall(0, nil, fmt.Errorf("fake-error"))

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)
			Expect(response.Message).To(Equal("error-failed-to-get-state-for-fake-consumer"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when consumer does not exist", func() {
			dataContractProposal = resources.DataContractProposal{
				DataContractID: "fake-data-contract",
				ConsumerID:     "fake-consumer",
			}
			fakeStub.GetStateReturnsOnCall(0, nil, nil)

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)
			Expect(response.Message).To(Equal("error-business-does-not-exist-fake-consumer"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when getState for dataContractType returns an error", func() {

			dataContractProposal = resources.DataContractProposal{
				DataContractID:     "fake-data-contract",
				ConsumerID:         "fake-consumer",
				DataContractTypeID: "fake-data-contract-type",
			}
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, nil, fmt.Errorf("fake-get-data-contract-type-error"))

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal("error-failed-to-get-state-for-fake-data-contract-type"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(2))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when dataContractType does not exist", func() {

			dataContractProposal = resources.DataContractProposal{
				DataContractID:     "fake-data-contract",
				ConsumerID:         "fake-consumer",
				DataContractTypeID: "fake-data-contract-type",
			}
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, nil, nil)

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal("error-dataContractType-does-not-exist-fake-data-contract-type"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(2))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when unmarshaling fails for dataContractType", func() {

			dataContractProposal = resources.DataContractProposal{
				DataContractID:     "fake-data-contract",
				ConsumerID:         "fake-consumer",
				DataContractTypeID: "fake-data-contract-type",
			}
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, []byte("fake-data-contact-type"), nil)

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal("error-unmarshaling-fake-data-contract-type"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(2))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when payment is not approved", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())
			price.Amount = 200
			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)

			dataContractProposal = resources.DataContractProposal{
				DataContractID:     "fake-data-contract",
				ConsumerID:         "fake-consumer",
				DataContractTypeID: "fake-data-contract-type",
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal("no-enough-funds-to-fulfill-allowances"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(3))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when getState on dataContract fails", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(3, providerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(4, nil, fmt.Errorf("fake-err"))

			dataContractProposal = resources.DataContractProposal{
				DataContractID:     "fake-data-contract",
				ConsumerID:         "fake-consumer",
				DataContractTypeID: "fake-data-contract-type",
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal("error-failed-to-get-state-for-fake-data-contract"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(5))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should fail when dataContract exists", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())
			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(3, providerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(4, []byte("fake-bytes"), nil)

			dataContractProposal = resources.DataContractProposal{
				DataContractID:     "fake-data-contract",
				ConsumerID:         "fake-consumer",
				DataContractTypeID: "fake-data-contract-type",
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal("error-resource-already-exists-fake-data-contract"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(5))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should fail when dataContract proposal timestamp is before dataContract type startTime", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			startTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2018-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			endTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2020-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractTypeExtras = resources.ContractTypeExtras{
				StartTime: startTime,
				EndTime:   endTime,
			}

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				DataType:   "FILE",
				Extras:     dataContractTypeExtras,
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(3, providerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(4, nil, nil)

			contractStartTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2017-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractProposal = resources.DataContractProposal{
				DataContractID:        "fake-data-contract",
				ConsumerID:            "fake-consumer",
				DataContractTypeID:    "fake-data-contract-type",
				DataContractTimestamp: contractStartTime,
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal(fmt.Sprintf("error-contract-creation-time-should-be-between-%s-and-%s", startTime, endTime)))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(5))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should fail when dataContract proposal timestamp is after dataContract type  endTime", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			startTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2018-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			endTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2020-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractTypeExtras = resources.ContractTypeExtras{
				StartTime: startTime,
				EndTime:   endTime,
			}

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				DataType:   "FILE",
				Extras:     dataContractTypeExtras,
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(3, providerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(4, nil, nil)

			contractStartTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2022-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractProposal = resources.DataContractProposal{
				DataContractID:        "fake-data-contract",
				ConsumerID:            "fake-consumer",
				DataContractTypeID:    "fake-data-contract-type",
				DataContractTimestamp: contractStartTime,
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal(fmt.Sprintf("error-contract-creation-time-should-be-between-%s-and-%s", startTime, endTime)))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(5))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should fail when dataContract proposal endDateTime is before dataContract type  startTime", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			startTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2018-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			endTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2020-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractTypeExtras = resources.ContractTypeExtras{
				StartTime: startTime,
				EndTime:   endTime,
			}

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				DataType:   "STREAM",
				Extras:     dataContractTypeExtras,
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(3, providerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(4, nil, nil)

			contractStartTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2019-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			contractEndTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2015-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractProposalExtras = resources.ContractExtras{
				EndDateTime: contractEndTime,
			}

			dataContractProposal = resources.DataContractProposal{
				DataContractID:        "fake-data-contract",
				ConsumerID:            "fake-consumer",
				DataContractTypeID:    "fake-data-contract-type",
				DataContractTimestamp: contractStartTime,
				Extras:                dataContractProposalExtras,
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal(fmt.Sprintf("error-contract-EndTime-should-be-between-%s-and-%s", startTime, endTime)))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(5))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should fail when dataContract proposal endDateTime is after dataContract type  endTime", func() {
			consumer := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-consumer",
			}

			consumerBytes, err := json.Marshal(consumer)
			Expect(err).NotTo(HaveOccurred())

			startTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2018-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			endTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2020-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractTypeExtras = resources.ContractTypeExtras{
				StartTime: startTime,
				EndTime:   endTime,
			}

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
				DataType:   "STREAM",
				Extras:     dataContractTypeExtras,
				PriceType:  price,
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, consumerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, dataContractTypeBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, consumerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(3, providerAccountBytes, nil)
			fakeStub.GetStateReturnsOnCall(4, nil, nil)

			contractStartTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2019-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			contractEndTime, err := time.Parse("2006-01-02T15:04:05.000Z", "2040-10-26T21:57:53.397Z")
			Expect(err).NotTo(HaveOccurred())

			dataContractProposalExtras = resources.ContractExtras{
				EndDateTime: contractEndTime,
			}

			dataContractProposal = resources.DataContractProposal{
				DataContractID:        "fake-data-contract",
				ConsumerID:            "ake-consumer",
				DataContractTypeID:    "fake-data-contract-type",
				DataContractTimestamp: contractStartTime,
				Extras:                dataContractProposalExtras,
			}

			response := resources.SubmitDataContractProposal(logger, fakeStub, dataContractProposal)

			Expect(response.Message).To(Equal(fmt.Sprintf("error-contract-EndTime-should-be-between-%s-and-%s", startTime, endTime)))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(5))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetDataContract", func() {

		It("should fail when GetDataContractState fails", func() {
			dataContract := resources.DataContract{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-data-contract",
			}
			dataContractBytes, err := json.Marshal(dataContract)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractBytes, nil)
			response := resources.GetDataContract(logger, fakeStub, "fake-data-contract")

			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.DATA_CONTRACT_DOCTYPE)
			Expect(response.Message).To(Equal(errorMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should succeed when GetDataContractState succeeds", func() {
			dataContract := resources.DataContract{
				DocType: resources.DATA_CONTRACT_DOCTYPE,
				ID:      "fake-data-contract",
			}
			dataContractBytes, err := json.Marshal(dataContract)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractBytes, nil)
			response := resources.GetDataContract(logger, fakeStub, "fake-data-contract")

			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(response.Payload).To(Equal(dataContractBytes))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetDataContractState", func() {
		It("should fail when stub fails to get GetDataContractState state", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("error-getting-data-contract"))
			dataContractBytes, err := resources.GetDataContractState(logger, fakeStub, "fake-data-contract")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-data-contract"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractBytes).To(BeNil())
		})

		It("should fail when stub fails to get GetDataContractState state", func() {
			fakeStub.GetStateReturns(nil, nil)
			dataContractBytes, err := resources.GetDataContractState(logger, fakeStub, "fake-data-contract")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-dataContract-does-not-exist-fake-data-contract"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractBytes).To(BeNil())
		})

		It("should fail when json unmarshalling fails for GetDataContractState state", func() {
			fakeStub.GetStateReturns([]byte("fake-type"), nil)
			dataContractBytes, err := resources.GetDataContractState(logger, fakeStub, "fake-data-contract")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-unmarshaling-fake-data-contract"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractBytes).To(BeNil())
		})

		It("should fail when dataContract docType does not correspond to the returned bytes", func() {
			dataContract := resources.DataContract{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-data-contract",
			}

			dataContractBytes, err := json.Marshal(dataContract)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractBytes, nil)
			dataContractBytes, err = resources.GetDataContractState(logger, fakeStub, "fake-data-contract")

			Expect(err).To(HaveOccurred())
			errMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.DATA_CONTRACT_DOCTYPE)
			Expect(err.Error()).To(Equal(errMsg))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractBytes).To(BeNil())
		})

		It("should succeed when dataContractType exists with correen docType", func() {
			dataContract := resources.DataContract{
				DocType: resources.DATA_CONTRACT_DOCTYPE,
				ID:      "fake-data-contract",
			}

			dataContractBytes, err := json.Marshal(dataContract)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractBytes, nil)
			dataContractBytes, err = resources.GetDataContractState(logger, fakeStub, "fake-data-contract")

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractBytes).NotTo(BeNil())
		})
	})
	Context(".SetFileStatus", func() {
		It("should fail when stub fails update dataContract", func() {
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))
			dataContract := resources.DataContract{
				DocType: resources.DATA_CONTRACT_DOCTYPE,
				ID:      "fake-data-contract",
			}

			err = dataContract.SetFileStatus(logger, fakeStub, resources.PROPOSAL, resources.Hash{}, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-dataContract-fake-error"))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))

		})

		It("should succeed when stub succeeds to update dataContract", func() {
			dataContract := resources.DataContract{
				DocType: resources.DATA_CONTRACT_DOCTYPE,
				ID:      "fake-data-contract",
			}

			err = dataContract.SetFileStatus(logger, fakeStub, resources.PROPOSAL, resources.Hash{}, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))

		})
	})
})
