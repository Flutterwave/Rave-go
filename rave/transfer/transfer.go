package transfer

import (
	"Rave-go/rave"
	"Rave-go/rave/helper"
	"fmt"
	"go/types"
)

var noresponse = map[string]interface{}{
	"": "",
}

type SingleTransfer interface {
	InitiateSingleTransfer(data SinglePaymentData) (error error, response map[string]interface{})
}

type BulkTransfer interface {
	InitiateBulkTransfer(data BulkPaymentData) (error error, response map[string]interface{})
}

type SingleTransferGetter interface {
	FetchTransfer(reference string) (error error, response map[string]interface{})
}

type AllTransfersGetter interface {
	FetchAllTransfers() (error error, response map[string]interface{})
}

type TransferFeeGetter interface {
	GetTransferFee()
}

type RaveBalanceGetter interface {
	GetRaveBalance()
}

type BulkTransferStatus interface {
	GetBulkTransferStatus() (error error, response map[string]interface{})
}

type TransferHelpers interface {
	SingleTransferGetter
	AllTransfersGetter
	TransferFeeGetter
	RaveBalanceGetter
	BulkTransferStatus
}

type TransferInterface interface {
	SingleTransfer
	BulkTransfer
	TransferHelpers
}

type SinglePaymentData struct {
	SecKey          string      `json:"seckey"`
	AccountBank     string      `json:"account_bank"`
	AccountNumber   string      `json:"account_number"`
	Amount          int         `json:"amount"`
	Narration       string      `json:"narration"`
	Currency        string      `json:"currency"`
	Reference       string      `json:"reference"`
	Meta            types.Slice `json:"meta"`
	BeneficiaryName string      `json:"beneficiary_name"`
}

type BulkPaymentData struct {
	SecKey   string              `json:"seckey"`
	Title    string              `json:"title"`
	BulkData []map[string]string `json:"bulk_data"`
}

type AccountResolveData struct {
	PublicKey        string `json:"PBFPubKey"`
	RecipientAccount string `json:"recipientaccount"`
	DestBankCode     string `json:"destbankcode"`
	Currency         string `json:"currency"`
	Country          string `json:"country"`
}

type Transfer struct {
	rave.Rave
}

// Initiates a single transfer
func (t Transfer) InitiateSingleTransfer(paymentData SinglePaymentData) (error error, response map[string]interface{}) {
	paymentData.SecKey = t.GetSecretKey()
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "initiate")
	err, response := helper.MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

// initiates a bulk transfer
func (t Transfer) InitiateBulkTransfer(paymentData BulkPaymentData) (error error, response map[string]interface{}) {
	paymentData.SecKey = t.GetSecretKey()
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "bulk")
	err, response := helper.MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

// Gets a single transfer
func (t Transfer) FetchTransfer(reference string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"reference": reference,
		"seckey":    t.GetSecretKey(),
	}
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "fetch")
	err, response := helper.MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

// Gets all transfers
func (t Transfer) FetchAllTransfers(status string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": t.GetSecretKey(),
		"status": status,
	}
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "fetch")
	err, response := helper.MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (t Transfer) GetBulkTransferStatus(batch_id string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey":   t.GetSecretKey(),
		"batch_id": batch_id,
	}
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "fetch")
	err, response := helper.MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

// Gets the transfer fee
func (t Transfer) GetTransferFee(currency string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey":   t.GetSecretKey(),
		"currency": currency,
	}
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "fee")
	err, response := helper.MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

// Gets a customer's Rave available Balance
func (t Transfer) GetRaveBalance(currency string) (error error, response map[string]interface{}) {
	paymentData := struct {
		Currency string `json:"currency"`
		Seckey   string `json:"seckey"`
	}{
		Currency: currency,
		Seckey:   t.GetSecretKey(),
	}

	url := t.GetBaseURL() + t.GetEndpoint("transfer", "balance")
	err, response := helper.MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (t Transfer) ResolveAccount(account_data AccountResolveData) (error error, response map[string]interface{}) {
	url := t.GetBaseURL() + t.GetEndpoint("transfer", "accountVerification")
	err, response := helper.MakePostRequest(account_data, url)
	if err != nil {
		return err, noresponse
	}
	fmt.Printf("%v\n", response["status"])
	return nil, response
}
