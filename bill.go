package qpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Bill struct {
	TemplateID  string   `json:"template_id"`
	MerchantID  string   `json:"merchant_id"`
	BranchID    string   `json:"branch_id"`
	PosID       string   `json:"pos_id"`
	Receiver    Receiver `json:"receiver"`
	BillNumber  string   `json:"bill_no"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	Amount      int      `json:"amount"`
	BtukCode    string   `json:"btuk_code"`
	VatFlag     string   `json:"vat_flag"`
}

type Receiver struct {
	ID          string `json:"id"`
	RegisterNo  string `json:"register_no"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Note        string `json:"note"`
}

type BillResponse struct {
	PaymentID    int64  `json:"payment_id"`
	QPayCode     string `json:"qPay_QRcode"`
	QPayImage    string `json:"qPay_QRimage"`
	QPayURL      string `json:"qPay_url"`
	QPayShortURL string `json:"qPay_shortUrl"`
	QPayDeepLink []struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"qPay_deeplink"`
}

func CreateBill(bill Bill) (*BillResponse, error) {
	requestBody, err := json.Marshal(bill)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token, err := GetToken()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", os.Getenv("QPAY_URL")+"/bill/create", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data *BillResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
