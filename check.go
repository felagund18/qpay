package qpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CheckResponse struct {
	ID         int64  `json:"id"`
	TemplateID string `json:"template_id"`
	MerchantID string `json:"merchant_id"`
	BranchInfo struct {
		ID      string      `json:"id"`
		Email   string      `json:"email"`
		Name    string      `json:"name"`
		Phone   interface{} `json:"phone"`
		Address interface{} `json:"address"`
	} `json:"branch_info"`
	StaffInfo            interface{} `json:"staff_info"`
	MerchantCustomerInfo struct {
		ID          string `json:"id"`
		RegisterNo  string `json:"register_no"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Note        string `json:"note"`
	} `json:"merchant_customer_info"`
	BillNo              string `json:"bill_no"`
	AllowPartialPayment bool   `json:"allow_partial_payment"`
	AllowTip            bool   `json:"allow_tip"`
	CurrencyType        string `json:"currency_type"`
	GoodsDetail         []struct {
		ID        interface{}   `json:"id"`
		Name      string        `json:"name"`
		BarCode   interface{}   `json:"bar_code"`
		Quantity  int           `json:"quantity"`
		UnitPrice int           `json:"unit_price"`
		TaxInfo   []interface{} `json:"tax_info"`
	} `json:"goods_detail"`
	Surcharge struct {
	} `json:"surcharge"`
	Discount struct {
	} `json:"discount"`
	Note          interface{} `json:"note"`
	Attach        interface{} `json:"attach"`
	PaymentMethod string      `json:"payment_method"`
	CustomerQr    interface{} `json:"customer_qr"`
	VatInfo       struct {
		IsAutoGenerate bool          `json:"is_auto_generate"`
		Vats           []interface{} `json:"vats"`
	} `json:"vat_info"`
	PaymentInfo struct {
		PaymentStatus   string        `json:"payment_status"`
		TransactionID   interface{}   `json:"transaction_id"`
		LastPaymentDate interface{}   `json:"last_payment_date"`
		Transactions    []interface{} `json:"transactions"`
	} `json:"payment_info"`
}

func Check(billNumber string) (*CheckResponse, error) {
	requestBody, err := json.Marshal(map[string]string{
		"merchant_id": "TEST_MERCHANT",
		"bill_no":     billNumber,
	})
	if err != nil {
		return nil, err
	}

	token, err := GetToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", os.Getenv("QPAY_URL")+"/bill/check", bytes.NewBuffer(requestBody))
	if err != nil {
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

	if response.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	var data *CheckResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func CheckPaymentID(paymentID string) {
}
