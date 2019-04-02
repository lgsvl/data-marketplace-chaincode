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

var _ = Describe("Token", func() {
	var (
		fakeStub *fakes.ChaincodeStub
		logger   *shim.ChaincodeLogger
		err      error
	)
	BeforeEach(func() {
		logger = shim.NewLogger("token-test-logger")
		fakeStub = new(fakes.ChaincodeStub)

	})

	Context("initially", func() {
		var token resources.Token
		It("should fail when stub fails to get state", func() {

			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			_, err = resources.NewToken(logger, fakeStub, "fake-token")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-check-token-fake-error"))
		})

		It("should fail when token already exists", func() {
			fakeToken := resources.Token{ID: "fake-token"}
			fakeTokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())
			fakeStub.GetStateReturns(fakeTokenBytes, nil)

			_, err = resources.NewToken(logger, fakeStub, "fake-token")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("this-token-already-exists-fake-token"))
		})

		It("has 0 total supply", func() {
			fakeStub.GetStateReturnsOnCall(0, nil, nil)
			fakeToken := resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE}
			tokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())
			fakeStub.GetStateReturnsOnCall(1, tokenBytes, nil)
			token, err = resources.NewToken(logger, fakeStub, "fake-token")

			Expect(err).NotTo(HaveOccurred())
			Expect(token.TotalSupply(logger, fakeStub)).Should(BeZero())
		})

		It("should fail when the supply is negative", func() {
			_, err := resources.NewTokenWithSupply(logger, fakeStub, "fake-token", -1000.0)
			Expect(err).To(HaveOccurred())
		})
		It("should succeed", func() {
			_, err := resources.NewTokenWithSupply(logger, fakeStub, "fake-token", 1000.0)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("TotalSupply", func() {
		var (
			fakeToken resources.Token
		)
		BeforeEach(func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}
		})
		It("should fail when stub fails to get state", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			_, err := fakeToken.TotalSupply(logger, fakeStub)
			Expect(err).To(HaveOccurred())

			Expect(err.Error()).To(Equal("error-failed-to-get-token-state-for-fake-token"))
		})
		It("should fail when stub returns a bad json", func() {
			fakeStub.GetStateReturns([]byte("bad-json"), nil)
			_, err := fakeToken.TotalSupply(logger, fakeStub)
			Expect(err).To(HaveOccurred())

			Expect(err.Error()).To(Equal("error-unmarshalling-fake-token"))
		})

		It("should succeed and total supply should be equal to 1000.0", func() {
			tokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(tokenBytes, nil)
			supply, err := fakeToken.TotalSupply(logger, fakeStub)
			Expect(err).NotTo(HaveOccurred())

			Expect(supply).To(Equal(1000.0))
		})

	})

	Context("SetAccountBalance", func() {
		var (
			account   resources.Account
			fakeToken resources.Token
		)
		BeforeEach(func() {
			account = resources.Account{ID: "fake-account", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}

		})
		It("should fail when stub fails to get state", func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}

			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			_, err := fakeToken.SetAccountBalance(logger, fakeStub, account, 100.0)
			Expect(err).To(HaveOccurred())

			Expect(err.Error()).To(Equal("error-failed-to-get-token-state-for-fake-token"))
		})
		It("should fail when stub returns a bad json", func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}

			fakeStub.GetStateReturns([]byte("bad-json"), nil)
			_, err := fakeToken.SetAccountBalance(logger, fakeStub, account, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-unmarshalling-fake-token"))
		})

		It("should fail when transfer amount is greater then remaining supply", func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}

			tokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(tokenBytes, nil)
			_, err = fakeToken.SetAccountBalance(logger, fakeStub, account, 1000000.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-transfer-amount-greater-then-remaining-supply"))

		})

		It("should fail when transfer amount fail to putState of account to blockchain", func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}

			tokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())

			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, tokenBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, accountBytes, nil)

			fakeStub.PutStateReturnsOnCall(0, fmt.Errorf("fake-error"))

			_, err = fakeToken.SetAccountBalance(logger, fakeStub, account, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))

		})

		It("should fail when transfer amount fail to putState of token to blockchain", func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}

			tokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())

			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, tokenBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, accountBytes, nil)

			fakeStub.PutStateReturnsOnCall(0, nil)
			fakeStub.PutStateReturnsOnCall(1, fmt.Errorf("fake-error"))

			_, err = fakeToken.SetAccountBalance(logger, fakeStub, account, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-token-fake-error"))

		})

		It("should succeed when transfer amount succeeds to putState to blockchain", func() {
			fakeToken = resources.Token{ID: "fake-token", DocType: resources.TOKEN_DOCTYPE, TotalTokensSupply: 1000.0, RemainingSupply: 1000.0}

			tokenBytes, err := json.Marshal(fakeToken)
			Expect(err).NotTo(HaveOccurred())

			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, tokenBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, accountBytes, nil)

			fakeStub.PutStateReturnsOnCall(0, nil)
			fakeStub.PutStateReturnsOnCall(1, nil)

			success, err := fakeToken.SetAccountBalance(logger, fakeStub, account, 100.0)
			Expect(err).NotTo(HaveOccurred())
			Expect(success).To(BeTrue())

		})

	})

})
