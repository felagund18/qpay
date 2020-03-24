package handlers

import (
	"github.com/felagund18/boom"
	"github.com/labstack/echo"
	"net/http"

	"../qpay"
)

type RequestAuth struct {
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	GrantType    string `json:"grant_type" validate:"required"`
	RefreshToken string `json:"refresh_token"`
}

func HandleAuth(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	input := new(RequestAuth)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest("Can't bind input"))
	}

	if err := c.Validate(input); err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest("Input is not valid"))
	}

	result, err := qpay.Authenticate(input.ClientID, input.ClientSecret, input.GrantType, input.RefreshToken)
	if err != nil {
		return c.String(http.StatusBadRequest, boom.BadRequest(err.Error()))
	}

	return c.JSON(http.StatusOK, result)
}
