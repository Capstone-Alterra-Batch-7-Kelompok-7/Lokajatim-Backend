package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type MidtransClient struct {
	ServerKey  string
	ClientKey  string
	APIBaseURL string
}

type NotificationRequest struct {
	TransactionID     string `json:"transaction_id"`
	PaymentMethod     string `json:"payment_method"`
	TransactionStatus string `json:"transaction_status"`
	PaymentURL        string `json:"payment_url"`
}

func NewMidtransClient() *MidtransClient {
	return &MidtransClient{
		ServerKey:  os.Getenv("MIDTRANS_SERVER_KEY"),
		ClientKey:  os.Getenv("MIDTRANS_CLIENT_KEY"),
		APIBaseURL: "https://api.sandbox.midtrans.com/v2/",
	}
}

type CreateTransactionRequest struct {
	PaymentType        string             `json:"payment_type"`
	Amount             float64            `json:"amount"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
}

type TransactionDetails struct {
	OrderID  string  `json:"order_id"`
	GrossAmt float64 `json:"gross_amount"`
}

type CreateTransactionResponse struct {
	PaymentURL string `json:"redirect_url"`
}

func (m *MidtransClient) CreatePayment(orderID string, paymentType string, amount float64) (string, error) {
	url := m.APIBaseURL + "charge"
	requestData := CreateTransactionRequest{
		PaymentType: paymentType,
		Amount:      amount,
		TransactionDetails: TransactionDetails{
			OrderID:  orderID,
			GrossAmt: amount,
		},
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Basic "+m.ServerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	var createResp CreateTransactionResponse
	if err := json.Unmarshal(respBody, &createResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Midtrans API error: ", string(respBody))
		return "", fmt.Errorf("failed to create payment: %v", string(respBody))
	}

	return createResp.PaymentURL, nil
}

func (m *MidtransClient) VerifyTransactionStatus(orderID string) (string, error) {
	url := m.APIBaseURL + orderID + "/status"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Basic "+m.ServerKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	var notification NotificationRequest
	if err := json.Unmarshal(respBody, &notification); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Midtrans API error: ", string(respBody))
		return "", fmt.Errorf("failed to verify payment status: %v", string(respBody))
	}

	return notification.TransactionStatus, nil
}
