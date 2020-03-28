package qpay

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSchema(t *testing.T) {
	os.Setenv("QPAY_URL", "https://sandbox.qpay.mn/v1")
	InitQPay("qpay_test", "1234")

	bill, err := CreateBill(Bill{
		TemplateID: "TEST_INVOICE",
		MerchantID: "TEST_MERCHANT",
		BranchID:   "1",
		PosID:      "1",
		Receiver: Receiver{
			ID:          "1",
			RegisterNo:  "1",
			Name:        "John Smith",
			Email:       "a@a.com",
			PhoneNumber: "1",
			Note:        "1",
		},
		BillNumber:  "2020-02-02-02",
		Date:        "2020-03-03",
		Description: "Hello World",
		Amount:      1000,
		BtukCode:    "",
		VatFlag:     "0",
	})

	if err != nil {
		log.Println(err)
	}
	fmt.Println(bill)
	//
	//checkStatus, err := qpay.Check("2020-02-02-02")
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(checkStatus.PaymentInfo.PaymentStatus)
}