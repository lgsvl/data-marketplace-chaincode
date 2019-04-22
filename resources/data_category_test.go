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

var _ = Describe("DataCategory", func() {
	var (
		fakeStub     *fakes.ChaincodeStub
		logger       *shim.ChaincodeLogger
		dataCategory resources.DataCategory
	)
	BeforeEach(func() {
		logger = shim.NewLogger("dataCategory-test-logger")
		fakeStub = new(fakes.ChaincodeStub)

	})

	Context(".CreateDataCategory", func() {
		It("should fail when DocType does not correspond to DataCategory DocType", func() {
			dataCategory = resources.DataCategory{DocType: "fake-docType"}
			response := resources.CreateDataCategory(logger, fakeStub, dataCategory)
			Expect(response.Message).To(Equal("error-docType-does-not-match-fake-docType-vs-com.lge.svl.datamarketplace.contract.DataCategory"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
		})

		It("should fail when stub getState returns an error", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			response := resources.CreateDataCategory(logger, fakeStub, dataCategory)
			Expect(response.Message).To(Equal("fake-error"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when dataCategory exists", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataCategoryBytes, nil)

			response := resources.CreateDataCategory(logger, fakeStub, dataCategory)
			Expect(response.Message).To(Equal("this-category-already-exists-fake-category"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when stub fails to put business state", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}

			fakeStub.GetStateReturns(nil, nil)
			fakeStub.PutStateReturns(fmt.Errorf("error-put-dataCategory"))

			response := resources.CreateDataCategory(logger, fakeStub, dataCategory)
			Expect(response.Message).To(Equal("error-put-dataCategory"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to put business state", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}

			response := resources.CreateDataCategory(logger, fakeStub, dataCategory)
			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

	})

	Context(".GetDataCategory", func() {
		It("should fail when GetDataCategoryState fails", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataCategoryBytes, nil)
			response := resources.GetDataCategory(logger, fakeStub, "fake-category")

			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.DATA_CATEGORY_DOCTYPE)
			Expect(response.Message).To(Equal(errorMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should succeed when GetDataCategoryState succeeds", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataCategoryBytes, nil)
			response := resources.GetDataCategory(logger, fakeStub, "fake-category")

			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(response.Payload).To(Equal(dataCategoryBytes))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetDataCategoryState", func() {
		It("should fail when stub fails to get dataCategory state", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("category-not-found"))
			businessBytes, err := resources.GetDataCategoryState(logger, fakeStub, "fake-category")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-category"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(businessBytes).To(BeNil())
		})

		It("should fail when stub returns nil to get budataCategorysiness state", func() {
			fakeStub.GetStateReturns(nil, nil)
			dataCategoryBytes, err := resources.GetDataCategoryState(logger, fakeStub, "fake-category")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-category-does-not-exist-fake-category"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataCategoryBytes).To(BeNil())
		})

		It("should fail when json unmarshalling fails for dataCategory state", func() {
			fakeStub.GetStateReturns([]byte("fake-json"), nil)
			dataCategoryBytes, err := resources.GetDataCategoryState(logger, fakeStub, "fake-category")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-unmarshaling-category-fake-category"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataCategoryBytes).To(BeNil())
		})

		It("should fail when dataCategory DocType does not correspond to the returned bytes", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataCategoryBytes, nil)
			dataCategoryBytes, err = resources.GetDataCategoryState(logger, fakeStub, "fake-category")

			Expect(err).To(HaveOccurred())
			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.DATA_CATEGORY_DOCTYPE)
			Expect(err.Error()).To(Equal(errorMsg))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataCategoryBytes).To(BeNil())
		})

		It("should succeed when dataCategory exists and returned DocType is correct", func() {
			dataCategory = resources.DataCategory{
				DocType: resources.DATA_CATEGORY_DOCTYPE,
				ID:      "fake-category",
			}
			dataCategoryBytes, err := json.Marshal(dataCategory)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(dataCategoryBytes, nil)
			dataCategoryBytes, err = resources.GetDataCategoryState(logger, fakeStub, "fake-category")

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(dataCategoryBytes).NotTo(BeNil())
		})
	})

})
