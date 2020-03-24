package main

import (
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
}
