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

var _ = Describe("Business", func() {
	var (
		fakeStub *fakes.ChaincodeStub
		logger   *shim.ChaincodeLogger
		business resources.Business
	)
	BeforeEach(func() {
		logger = shim.NewLogger("business-test-logger")
		fakeStub = new(fakes.ChaincodeStub)

	})

	Context(".CreateBusiness", func() {
		It("should fail when DocType does not correspond to business DocType", func() {
			business = resources.Business{DocType: "fake-docType"}
			response := resources.CreateBusiness(logger, fakeStub, business)
			errMsg := fmt.Sprintf("error-docType-does-not-match-fake-docType-vs-%s", resources.BUSINESS_DOCTYPE)
			Expect(response.Message).To(Equal(errMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
		})

		It("should fail when stub getState returns an error", func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			response := resources.CreateBusiness(logger, fakeStub, business)
			Expect(response.Message).To(Equal("failed-to-get-business-fake-error"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when business exists", func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}
			businessBytes, err := json.Marshal(business)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(businessBytes, nil)

			response := resources.CreateBusiness(logger, fakeStub, business)
			Expect(response.Message).To(Equal("this-business-already-exists-fake-business"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when stub fails to put business state", func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}

			fakeStub.GetStateReturns(nil, nil)
			fakeStub.PutStateReturns(fmt.Errorf("error-put-business"))

			response := resources.CreateBusiness(logger, fakeStub, business)
			Expect(response.Message).To(Equal("error-put-business"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to put business state", func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}

			response := resources.CreateBusiness(logger, fakeStub, business)
			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetBusiness", func() {
		It("should fail when GetBusinessState fails", func() {
			business = resources.Business{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-business",
			}
			businessBytes, err := json.Marshal(business)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(businessBytes, nil)
			response := resources.GetBusiness(logger, fakeStub, "fake-business")

			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.BUSINESS_DOCTYPE)
			Expect(response.Message).To(Equal(errorMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should succeed when GetBusinessState succeeds", func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}
			businessBytes, err := json.Marshal(business)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(businessBytes, nil)
			response := resources.GetBusiness(logger, fakeStub, "fake-business")

			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(response.Payload).To(Equal(businessBytes))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetBusinessState", func() {
		It("should fail when stub fails to get business state", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("error-getting-business"))
			businessBytes, err := resources.GetBusinessState(logger, fakeStub, "fake-business")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-business"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(businessBytes).To(BeNil())
		})

		It("should fail when stub returns nil to get business state", func() {
			fakeStub.GetStateReturns(nil, nil)
			businessBytes, err := resources.GetBusinessState(logger, fakeStub, "fake-business")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-business-does-not-exist-fake-business"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(businessBytes).To(BeNil())
		})

		It("should fail when json unmarshalling fails for business state", func() {
			fakeStub.GetStateReturns([]byte("fake-json"), nil)
			businessBytes, err := resources.GetBusinessState(logger, fakeStub, "fake-business")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-unmarshalling-fake-business"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(businessBytes).To(BeNil())
		})

		It("should fail when business DocType does not correspond to the returned bytes", func() {
			business = resources.Business{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-business",
			}
			businessBytes, err := json.Marshal(business)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(businessBytes, nil)
			businessBytes, err = resources.GetBusinessState(logger, fakeStub, "fake-business")

			Expect(err).To(HaveOccurred())
			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.BUSINESS_DOCTYPE)
			Expect(err.Error()).To(Equal(errorMsg))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(businessBytes).To(BeNil())
		})

		It("should succeed when business exists and returned DocType is correct", func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}
			businessBytes, err := json.Marshal(business)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(businessBytes, nil)
			businessBytes, err = resources.GetBusinessState(logger, fakeStub, "fake-business")

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(businessBytes).NotTo(BeNil())
		})
	})

	Context(".AddReview", func() {
		var (
			review, review2 resources.Review
		)

		BeforeEach(func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
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
		It("should fail when stub fails to put business state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("business-error"))
			err := business.AddReview(logger, fakeStub, review)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("business-error"))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to put business state", func() {
			fakeStub.PutStateReturns(nil)
			err := business.AddReview(logger, fakeStub, review)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
			Expect(len(business.Reviews)).To(Equal(1))
			Expect(business.Reviews[0].ID).To(Equal("fake-review"))
			Expect(business.Score).To(Equal(float32(4)))
			Expect(business.NumberOfReviews).To(Equal(1))
		})
		It("should succeed to add two reviews and change score", func() {
			fakeStub.PutStateReturns(nil)
			err := business.AddReview(logger, fakeStub, review)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
			Expect(len(business.Reviews)).To(Equal(1))
			Expect(business.Reviews[0].ID).To(Equal("fake-review"))
			Expect(business.Score).To(Equal(float32(4)))
			Expect(business.NumberOfReviews).To(Equal(1))

			err = business.AddReview(logger, fakeStub, review2)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(2))
			Expect(len(business.Reviews)).To(Equal(2))
			Expect(business.Score).To(Equal(float32(3)))
			Expect(business.NumberOfReviews).To(Equal(2))
		})

	})

	Context(".AddPerson", func() {
		var person resources.Person

		BeforeEach(func() {
			business = resources.Business{
				DocType: resources.BUSINESS_DOCTYPE,
				ID:      "fake-business",
			}
			person = resources.Person{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-person",
			}

		})
		It("should fail when stub fails to put business state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("business-error"))
			err := business.AddPerson(logger, fakeStub, person)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("business-error"))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to put business state", func() {
			fakeStub.PutStateReturns(nil)
			err := business.AddPerson(logger, fakeStub, person)

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
			Expect(len(business.Persons)).To(Equal(1))
			Expect(business.Persons[0].ID).To(Equal("fake-person"))
		})
	})

})
