//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package resources_test

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/lgsvl/data-marketplace-chaincode/fakes"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DataContractType", func() {
	var (
		fakeStub         *fakes.ChaincodeStub
		logger           *shim.ChaincodeLogger
		dataContractType resources.DataContractType
		extras           resources.ContractTypeExtras
	)
	BeforeEach(func() {
		logger = shim.NewLogger("data-contract-type-test-logger")
		fakeStub = new(fakes.ChaincodeStub)

	})

	Context(".CreateDataContractType", func() {
		It("should fail when DocType does not correspond to dataContractType DocType", func() {
			dataContractType = resources.DataContractType{DocType: "fake-docType"}
			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)
			errMsg := fmt.Sprintf("error-docType-does-not-match-fake-docType-vs-%s", resources.DATA_CONTRACT_TYPE_DOCTYPE)
			Expect(response.Message).To(Equal(errMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
		})

		It("should fail when data type is stream and source is not specified", func() {
			extras = resources.ContractTypeExtras{
				StreamType: resources.PULL,
			}
			dataContractType = resources.DataContractType{
				DocType:  resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:       "fake-data-contract-type",
				DataType: resources.STREAM,
				Extras:   extras,
			}

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)
			Expect(response.Message).To(Equal("error-stream-source-endpoint-is-required"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
		})

		It("should fail when stub getState returns an error", func() {
			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
			}
			fakeStub.GetStateReturnsOnCall(0, nil, fmt.Errorf("fake-error"))

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)
			Expect(response.Message).To(Equal("error-failed-to-get-state-for-fake-category"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when getState for Business provider returns an error", func() {

			dataCategory := resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
			}
			fakeStub.GetStateReturnsOnCall(0, dataCategoryBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, nil, fmt.Errorf("fake-get-business-error"))

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)

			Expect(response.Message).To(Equal("error-failed-to-get-state-for-fake-provider"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(2))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when provider does not exist", func() {

			dataCategory := resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
			}
			fakeStub.GetStateReturnsOnCall(0, dataCategoryBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, nil, nil)

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)

			Expect(response.Message).To(Equal("error-business-does-not-exist-fake-provider"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(2))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when getStated on dataContractType fails", func() {
			provider := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-provider",
			}
			dataCategory := resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			providerBytes, err := json.Marshal(provider)
			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
			}
			fakeStub.GetStateReturnsOnCall(0, dataCategoryBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, providerBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, nil, fmt.Errorf("fake-err"))

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)

			Expect(response.Message).To(Equal("error-failed-to-get-state-for-fake-data-contract-type"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(3))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when dataContractType exists", func() {
			provider := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-provider",
			}
			dataCategory := resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			providerBytes, err := json.Marshal(provider)
			Expect(err).NotTo(HaveOccurred())

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
			}
			fakeStub.GetStateReturnsOnCall(0, dataCategoryBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, providerBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, []byte("fake"), nil)

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)

			Expect(response.Message).To(Equal("error-dataContractType-already-exists-fake-data-contract-type"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(3))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when putState fails", func() {
			provider := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-provider",
			}
			dataCategory := resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			providerBytes, err := json.Marshal(provider)
			Expect(err).NotTo(HaveOccurred())

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
			}
			fakeStub.GetStateReturnsOnCall(0, dataCategoryBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, providerBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, nil, nil)
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)

			Expect(response.Message).To(Equal("fake-error"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(3))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when putState succeeds", func() {
			provider := resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-provider",
			}
			dataCategory := resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			providerBytes, err := json.Marshal(provider)
			Expect(err).NotTo(HaveOccurred())

			dataContractType = resources.DataContractType{
				DocType:    resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:         "fake-data-contract-type",
				CategoryID: "fake-category",
				ProviderID: "fake-provider",
			}
			fakeStub.GetStateReturnsOnCall(0, dataCategoryBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, providerBytes, nil)
			fakeStub.GetStateReturnsOnCall(2, nil, nil)
			fakeStub.PutStateReturns(nil)

			response := resources.CreateDataContractType(logger, fakeStub, dataContractType)

			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(3))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetDataContractType", func() {
		It("should fail when GetDataContractTypeState fails", func() {
			dataContractType = resources.DataContractType{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-data-contract-type",
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractTypeBytes, nil)
			response := resources.GetDataContractType(logger, fakeStub, "fake-data-contract-type")

			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.DATA_CONTRACT_TYPE_DOCTYPE)
			Expect(response.Message).To(Equal(errorMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should succeed when GetDataContractState succeeds", func() {
			dataContractType = resources.DataContractType{
				DocType: resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:      "fake-data-contract-type",
			}
			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractTypeBytes, nil)
			response := resources.GetDataContractType(logger, fakeStub, "fake-data-contract-type")

			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(response.Payload).To(Equal(dataContractTypeBytes))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetDataContractTypeState", func() {
		It("should fail when stub fails to get GetDataContractTypeState state", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("error-getting-data-contract-type"))
			dataContractTypeBytes, err := resources.GetDataContractTypeState(logger, fakeStub, "fake-data-contract-type")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-data-contract-type"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractTypeBytes).To(BeNil())
		})

		It("should fail when stub fails to get GetDataContractTypeState state", func() {
			fakeStub.GetStateReturns(nil, nil)
			dataContractTypeBytes, err := resources.GetDataContractTypeState(logger, fakeStub, "fake-data-contract-type")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-dataContractType-does-not-exist-fake-data-contract-type"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractTypeBytes).To(BeNil())
		})

		It("should fail when json unmarshalling fails for GetDataContractTypeState state", func() {
			fakeStub.GetStateReturns([]byte("fake-type"), nil)
			dataContractTypeBytes, err := resources.GetDataContractTypeState(logger, fakeStub, "fake-data-contract-type")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-unmarshaling-fake-data-contract-type"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractTypeBytes).To(BeNil())
		})

		It("should fail when dataContractType class does not correspond to the returned bytes", func() {
			dataContractType = resources.DataContractType{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-data-contract-type",
			}

			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractTypeBytes, nil)
			dataContractTypeBytes, err = resources.GetDataContractTypeState(logger, fakeStub, "fake-data-contract-type")

			Expect(err).To(HaveOccurred())
			errMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.DATA_CONTRACT_TYPE_DOCTYPE)
			Expect(err.Error()).To(Equal(errMsg))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractTypeBytes).To(BeNil())
		})

		It("should succeed when dataContractType exists with correct class", func() {
			dataContractType = resources.DataContractType{
				DocType: resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:      "fake-data-contract-type",
			}

			dataContractTypeBytes, err := json.Marshal(dataContractType)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataContractTypeBytes, nil)
			dataContractTypeBytes, err = resources.GetDataContractTypeState(logger, fakeStub, "fake-data-contract-type")

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataContractTypeBytes).NotTo(BeNil())
		})
	})

	Context(".AddReview", func() {
		var (
			review, review2 resources.Review
		)

		BeforeEach(func() {
			dataContractType = resources.DataContractType{
				DocType: resources.DATA_CONTRACT_TYPE_DOCTYPE,
				ID:      "fake-data-contract-type",
			}
			review = resources.Review{
				DocType: resources.REVIEW_DOCTYPE,
				ID:      "fake-review",
				Score:   4,
			}
			review2 = resources.Review{
				DocType: resources.REVIEW_DOCTYPE,
				ID:      "fake-review2",
				Score:   2,
			}

		})
		It("should fail when stub fails to put dataContractType state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("data-contract-type-error"))
			err := dataContractType.AddReview(logger, fakeStub, review)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("data-contract-type-error"))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to put dataContractType state", func() {
			fakeStub.PutStateReturns(nil)
			err := dataContractType.AddReview(logger, fakeStub, review)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
			Expect(len(dataContractType.Reviews)).To(Equal(1))
			Expect(dataContractType.Reviews[0].ID).To(Equal("fake-review"))
			Expect(dataContractType.Score).To(Equal(float32(4)))
			Expect(dataContractType.NumberOfReviews).To(Equal(1))
		})
		It("should succeed to add two reviews and change score", func() {
			fakeStub.PutStateReturns(nil)
			err := dataContractType.AddReview(logger, fakeStub, review)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
			Expect(len(dataContractType.Reviews)).To(Equal(1))
			Expect(dataContractType.Reviews[0].ID).To(Equal("fake-review"))
			Expect(dataContractType.Score).To(Equal(float32(4)))
			Expect(dataContractType.NumberOfReviews).To(Equal(1))

			err = dataContractType.AddReview(logger, fakeStub, review2)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(2))
			Expect(len(dataContractType.Reviews)).To(Equal(2))
			Expect(dataContractType.Score).To(Equal(float32(3)))
			Expect(dataContractType.NumberOfReviews).To(Equal(2))
		})

	})

})
