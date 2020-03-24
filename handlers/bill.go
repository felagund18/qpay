package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/felagund18/boom"
	"github.com/labstack/echo"
	"net/http"

	"../qpay"
)

type RequestBill struct {
	TemplateID string `json:"template_id"`
	MerchantID string `json:"merchant_id" validate:"required"`
	BranchID   string `json:"branch_id"`
	PosID      string `json:"pos_id"`
	Receiver   struct {
		ID          string `json:"id"`
		RegisterNo  string `json:"register_no"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Note        string `json:"note"`
	} `json:"receiver"`
	BillNumber   string `json:"bill_no" validate:"required"`
	Date        string `json:"date" validate:"required"`
	Description string `json:"description" validate:"required"`
	Amount      int    `json:"amount" validate:"required"`
	BtukCode    string `json:"btuk_code"`
	VatFlag     string `json:"vat_flag"`
}

func HandleCreateBill(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	input := new(RequestBill)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest("Can't bind input"))
	}
	if err := c.Validate(input); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, boom.BadRequest("Input is not valid"))
	}

	inputJson, err := json.Marshal(input)
	if err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest("Can't marshal json"))
	}

	var bill qpay.Bill
	err = json.Unmarshal(inputJson, &bill)
	if err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest("Can't unmarshal json"))
	}

	result, err := qpay.CreateBill(bill)
	if err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest(err.Error()))
	}

	return c.JSON(http.StatusOK, result)
}
