package qpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
)

type AuthResponse struct {
	TokenType        string `json:"token_type"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	NotBeforePolicy  string `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
}

func Authenticate() (*AuthResponse, error) {
	clientVal, found := cacheInstance.Get("client")
	if !found {
		return nil, errors.New("No client information has found")
	}

	requestBody, err := json.Marshal(clientVal)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(os.Getenv("QPAY_URL") + "/auth/token", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body), os.Getenv("QPAY_URL") + "/auth/token")

	var data AuthResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	cacheInstance.Set("auth_response", data, cache.DefaultExpiration)
	cacheInstance.Set("token", data.AccessToken, cache.DefaultExpiration)

	return &data, nil
}
