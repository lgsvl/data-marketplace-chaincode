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

var _ = Describe("Account", func() {
	var (
		fakeStub *fakes.ChaincodeStub
		logger   *shim.ChaincodeLogger
		account  resources.Account
	)
	BeforeEach(func() {
		logger = shim.NewLogger("account-test-logger")
		fakeStub = new(fakes.ChaincodeStub)

	})

	Context(".CreateAccount", func() {
		It("should fail when DocType does not correspond to account DocType", func() {
			account = resources.Account{DocType: "fake-docType"}
			response := resources.CreateAccount(logger, fakeStub, account)
			errMsg := fmt.Sprintf("error-docType-does-not-match-fake-docType-vs-%s", resources.ACCOUNT_DOCTYPE)
			Expect(response.Message).To(Equal(errMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
		})

		It("should fail when stub getState returns an error", func() {
			account = resources.Account{
				DocType: resources.ACCOUNT_DOCTYPE,
				ID:      "fake-account",
			}
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			response := resources.CreateAccount(logger, fakeStub, account)
			Expect(response.Message).To(Equal("failed-to-get-account-fake-error"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when Account exists", func() {
			account = resources.Account{
				DocType: resources.ACCOUNT_DOCTYPE,
				ID:      "fake-account",
			}
			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)

			response := resources.CreateAccount(logger, fakeStub, account)
			Expect(response.Message).To(Equal("this-account-already-exists-fake-account"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(0))
		})

		It("should fail when stub fails to put Account state", func() {
			account = resources.Account{
				DocType: resources.ACCOUNT_DOCTYPE,
				ID:      "fake-account",
			}

			fakeStub.GetStateReturns(nil, nil)
			fakeStub.PutStateReturns(fmt.Errorf("error-put-Account"))

			response := resources.CreateAccount(logger, fakeStub, account)
			Expect(response.Message).To(Equal("failed-to-update-account-error-put-Account"))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to put Account state and account balance should equal 0", func() {
			account = resources.Account{
				DocType:    resources.ACCOUNT_DOCTYPE,
				ID:         "fake-account",
				Allowances: map[string]float64{},
			}

			fakeStub.GetStateReturns(nil, nil)
			fakeStub.PutStateReturns(nil)

			response := resources.CreateAccount(logger, fakeStub, account)
			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))

			createdAccount := resources.Account{}

			err := json.Unmarshal(response.Payload, &createdAccount)
			Expect(err).NotTo(HaveOccurred())

			Expect(createdAccount.Balance).To(Equal(0.0))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetAccount", func() {
		It("should fail when GetAccountState fails", func() {
			account = resources.Account{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-account",
			}
			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, fmt.Errorf("fake-error"))
			response := resources.GetAccount(logger, fakeStub, "fake-account")

			errorMsg := fmt.Sprintf("error-failed-to-get-state-for-fake-account")
			Expect(response.Message).To(Equal(errorMsg))
			Expect(response.Status).To(Equal(int32(shim.ERROR)))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should succeed when GetAccountState succeeds", func() {
			account = resources.Account{
				DocType: resources.ACCOUNT_DOCTYPE,
				ID:      "fake-account",
			}
			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)
			response := resources.GetAccount(logger, fakeStub, "fake-account")

			Expect(response.Message).To(Equal(""))
			Expect(response.Status).To(Equal(int32(shim.OK)))
			Expect(response.Payload).To(Equal(accountBytes))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})
	})

	Context(".GetAccountState", func() {
		It("should fail when stub fails to get account state", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("error-getting-account"))
			_, err := resources.GetAccountState(logger, fakeStub, "fake-account")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-account"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should fail when stub returns nil to get Account state", func() {
			fakeStub.GetStateReturns(nil, nil)
			_, err := resources.GetAccountState(logger, fakeStub, "fake-account")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-account-does-not-exist-fake-account"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should fail when json unmarshalling fails for Account state", func() {
			fakeStub.GetStateReturns([]byte("fake-json"), nil)
			_, err := resources.GetAccountState(logger, fakeStub, "fake-account")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-unmarshalling-fake-account"))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should fail when Account DocType does not correspond to the returned bytes", func() {
			account = resources.Account{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-account",
			}
			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)
			account, err = resources.GetAccountState(logger, fakeStub, "fake-account")

			Expect(err).To(HaveOccurred())
			errorMsg := fmt.Sprintf("error-docType-does-not-match-%s-vs-%s", resources.PERSON_DOCTYPE, resources.ACCOUNT_DOCTYPE)
			Expect(err.Error()).To(Equal(errorMsg))
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
		})

		It("should succeed when Account exists and returned DocType is correct", func() {
			account = resources.Account{
				DocType: resources.ACCOUNT_DOCTYPE,
				ID:      "fake-account",
			}
			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)
			account, err = resources.GetAccountState(logger, fakeStub, "fake-account")

			Expect(err).NotTo(HaveOccurred())
			Expect(fakeStub.GetStateCallCount()).To(Equal(1))
			Expect(account.ID).To(Equal("fake-account"))
		})
	})

	Context(".SetAccountState", func() {

		It("should fail when stub fails to set account state", func() {
			account = resources.Account{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-account",
			}

			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))
			err := resources.SetAccountState(logger, fakeStub, account)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})

		It("should succeed when stub succeeds to set account state", func() {
			account = resources.Account{
				DocType: resources.PERSON_DOCTYPE,
				ID:      "fake-account",
			}

			fakeStub.PutStateReturns(nil)
			err := resources.SetAccountState(logger, fakeStub, account)

			Expect(err).NotTo(HaveOccurred())

			Expect(fakeStub.PutStateCallCount()).To(Equal(1))
		})
	})

	Context(".BalanceOf", func() {
		var (
			account resources.Account
		)
		BeforeEach(func() {

			account = resources.Account{ID: "fake-account", DocType: resources.ACCOUNT_DOCTYPE, Balance: 100}

		})

		It("should fail when stub fails to get  account", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))

			_, err := resources.BalanceOf(logger, fakeStub, account)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-account"))
		})

		It("should return balance when stub succeeds to get account", func() {
			accountBytes, err := json.Marshal(account)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)

			balance, err := resources.BalanceOf(logger, fakeStub, account)
			Expect(err).NotTo(HaveOccurred())
			Expect(balance).To(Equal(100.0))
		})

	})

	Context(".Approve", func() {
		var (
			owner   resources.Account
			spender resources.Account
		)
		BeforeEach(func() {
			owner = resources.Account{ID: "fake-owner", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}
			spender = resources.Account{ID: "fake-spender", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}
		})

		It("should fail when stub fails to get owner account", func() {
			fakeStub.GetStateReturnsOnCall(0, nil, fmt.Errorf("fake-error"))

			_, err := resources.Approve(logger, fakeStub, owner.ID, spender.ID, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-owner"))
		})
		It("should fail when owner account does not have enough funds to fulfill all his allowances", func() {
			owner.Balance = 100
			owner.TotalAllowances = 90
			ownerBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(ownerBytes, nil)

			_, err = resources.Approve(logger, fakeStub, owner.ID, spender.ID, 15)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no-enough-funds-to-fulfill-allowances"))
		})

		It("should fail when stub fails to get spender account", func() {
			owner.Balance = 100
			owner.TotalAllowances = 90
			ownerBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, ownerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, nil, fmt.Errorf("fake-error"))

			_, err = resources.Approve(logger, fakeStub, owner.ID, spender.ID, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-spender"))
		})
		It("should fail when putState fails to update blockchain", func() {
			owner.Balance = 100
			owner.TotalAllowances = 90
			ownerBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			spenderBytes, err := json.Marshal(spender)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, ownerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, spenderBytes, nil)
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))

			_, err = resources.Approve(logger, fakeStub, owner.ID, spender.ID, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
		})
		It("should succeed when putState succeeds to update blockchain", func() {
			owner.Balance = 100
			owner.TotalAllowances = 90
			ownerBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			spenderBytes, err := json.Marshal(spender)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, ownerBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, spenderBytes, nil)
			fakeStub.PutStateReturns(nil)

			success, err := resources.Approve(logger, fakeStub, owner.ID, spender.ID, 10)
			Expect(err).NotTo(HaveOccurred())
			Expect(success).To(BeTrue())
		})

	})

	Context(".Allowance", func() {
		var (
			owner   resources.Account
			spender resources.Account
		)
		BeforeEach(func() {

			owner = resources.Account{ID: "fake-owner", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}
			spender = resources.Account{ID: "fake-spender", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}

		})

		It("should fail when stub fails to get owner account", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))

			_, err := resources.Allowance(logger, fakeStub, owner, spender)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-owner"))
		})

		It("should return 0 when the owner has no allowances", func() {
			accountBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)

			allowance, err := resources.Allowance(logger, fakeStub, owner, spender)
			Expect(err).NotTo(HaveOccurred())
			Expect(allowance).To(Equal(0.0))
		})
		It("should return 0 when the owner has no allowance for the spender", func() {
			owner.Allowances["fake"] = 10
			accountBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)

			allowance, err := resources.Allowance(logger, fakeStub, owner, spender)
			Expect(err).NotTo(HaveOccurred())
			Expect(allowance).To(Equal(0.0))
		})
		It("should return allowance when there is an allowance", func() {
			owner.Allowances[spender.ID] = 10
			accountBytes, err := json.Marshal(owner)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturns(accountBytes, nil)

			allowance, err := resources.Allowance(logger, fakeStub, owner, spender)
			Expect(err).NotTo(HaveOccurred())
			Expect(allowance).To(Equal(10.0))
		})
	})

	Context(".TransferFrom", func() {
		var (
			source      resources.Account
			destination resources.Account
		)
		BeforeEach(func() {
			source = resources.Account{ID: "fake-source", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}
			destination = resources.Account{ID: "fake-destination", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}

		})

		It("should fail when stub fails to get source account", func() {
			fakeStub.GetStateReturns(nil, fmt.Errorf("fake-error"))
			_, err := resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-source"))

		})

		It("should fail when stub fails to get destination account", func() {
			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, nil, fmt.Errorf("fake-error"))

			_, err = resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("error-failed-to-get-state-for-fake-destination"))

		})

		It("should fail when source account does not have allowance to destination account", func() {
			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())
			destinationBytes, err := json.Marshal(destination)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, destinationBytes, nil)

			_, err = resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no-enough-allowance-to-transfer"))

		})

		It("should fail when source account does not have enough allowance", func() {
			source.Allowances["fake-destination"] = 50
			source.Balance = 1000
			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())
			destinationBytes, err := json.Marshal(destination)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, destinationBytes, nil)

			_, err = resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no-enough-allowance-to-transfer"))
		})

		It("should fail when stub fails to put source account state", func() {
			source.Allowances["fake-destination"] = 500
			source.Balance = 1000
			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())
			destinationBytes, err := json.Marshal(destination)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, destinationBytes, nil)

			fakeStub.PutStateReturnsOnCall(0, fmt.Errorf("fake-error"))
			_, err = resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
		})

		It("should fail when stub fails to reduce allowance from source account state", func() {
			source.Allowances["fake-destination"] = 500
			source.Balance = 1000
			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())
			destinationBytes, err := json.Marshal(destination)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, destinationBytes, nil)

			fakeStub.PutStateReturnsOnCall(0, nil)
			fakeStub.PutStateReturnsOnCall(1, fmt.Errorf("fake-error"))

			_, err = resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
		})

		It("should fail when stub fails to put destination account state", func() {
			source.Allowances["fake-destination"] = 500
			source.Balance = 1000

			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())
			destinationBytes, err := json.Marshal(destination)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, destinationBytes, nil)
			fakeStub.PutStateReturnsOnCall(0, nil)
			fakeStub.PutStateReturnsOnCall(1, nil)
			fakeStub.PutStateReturnsOnCall(2, fmt.Errorf("fake-error"))
			fakeStub.PutStateReturnsOnCall(3, nil)

			_, err = resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 100.0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
		})

		It("should succeed when stub succeeds to put destination account state", func() {
			source.Allowances["fake-destination"] = 500
			source.Balance = 1000

			sourceBytes, err := json.Marshal(source)
			Expect(err).NotTo(HaveOccurred())
			destinationBytes, err := json.Marshal(destination)
			Expect(err).NotTo(HaveOccurred())

			fakeStub.GetStateReturnsOnCall(0, sourceBytes, nil)
			fakeStub.GetStateReturnsOnCall(1, destinationBytes, nil)
			fakeStub.PutStateReturns(nil)

			success, err := resources.TransferFrom(logger, fakeStub, source.ID, destination.ID, 100.0)
			Expect(err).NotTo(HaveOccurred())
			Expect(success).To(BeTrue())
		})
	})

	Context(".ReduceAllowance", func() {
		var (
			source      resources.Account
			destination resources.Account
		)
		BeforeEach(func() {
			source = resources.Account{ID: "fake-source", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{"fake-destination": 10}}
			destination = resources.Account{ID: "fake-destination", DocType: resources.ACCOUNT_DOCTYPE, Allowances: map[string]float64{}}

		})

		It("should fail when stub fails to update account state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))
			newAllowance, err := source.ReduceAllowance(logger, fakeStub, destination.ID, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
			Expect(newAllowance).To(Equal(10.0))
		})

		It("should return 0 if there are no allowances", func() {
			newAllowance, err := source.ReduceAllowance(logger, fakeStub, destination.ID, 6)
			Expect(err).NotTo(HaveOccurred())
			Expect(newAllowance).To(Equal(4.0))
		})

	})
	Context(".SetBalance", func() {
		var (
			account resources.Account
		)
		BeforeEach(func() {
			account = resources.Account{ID: "fake-source", DocType: resources.ACCOUNT_DOCTYPE, Balance: 5}
		})

		It("should fail when stub fails to update account state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))
			newBalance, err := account.SetBalance(logger, fakeStub, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
			Expect(newBalance).To(Equal(5.0))
		})

		It("should succeed and return new balance", func() {
			newBalance, err := account.SetBalance(logger, fakeStub, 6)
			Expect(err).NotTo(HaveOccurred())
			Expect(newBalance).To(Equal(6.0))
		})
	})

	Context(".AddFunds", func() {
		var (
			account resources.Account
		)
		BeforeEach(func() {
			account = resources.Account{ID: "fake-source", DocType: resources.ACCOUNT_DOCTYPE, Balance: 5}
		})

		It("should fail when stub fails to update account state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))
			newBalance, err := account.AddFunds(logger, fakeStub, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
			Expect(newBalance).To(Equal(5.0))
		})

		It("should return 0 if there are no allowances", func() {
			newBalance, err := account.AddFunds(logger, fakeStub, 6)
			Expect(err).NotTo(HaveOccurred())
			Expect(newBalance).To(Equal(11.0))
		})
	})

	Context(".RetrieveFunds", func() {
		var (
			account resources.Account
		)
		BeforeEach(func() {
			account = resources.Account{ID: "fake-source", DocType: resources.ACCOUNT_DOCTYPE, Balance: 5}
		})

		It("should fail when there are no enough funds", func() {
			newBalance, err := account.RetrieveFunds(logger, fakeStub, 10)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no-enough-funds-to-retrieve"))
			Expect(newBalance).To(Equal(5.0))
		})

		It("should fail when stub fails to update account state", func() {
			fakeStub.PutStateReturns(fmt.Errorf("fake-error"))
			newBalance, err := account.RetrieveFunds(logger, fakeStub, 2)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed-to-update-account-fake-error"))
			Expect(newBalance).To(Equal(5.0))
		})

		It("should succeed and returns new balance", func() {
			newBalance, err := account.RetrieveFunds(logger, fakeStub, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(newBalance).To(Equal(3.0))
		})
	})
})
