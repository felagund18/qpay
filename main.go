package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"./qpay"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	qpay.InitQPay("qpay_test", "1234")
	bill, err := qpay.CreateBill(qpay.Bill{
		TemplateID: "TEST_INVOICE",
		MerchantID: "TEST_MERCHANT",
		BranchID:   "1",
		PosID:      "1",
		Receiver: qpay.Receiver{
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
	fmt.Println(bill.PaymentID)

	checkStatus, err := qpay.Check("2020-02-02-02")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(checkStatus.PaymentInfo.PaymentStatus)
}
