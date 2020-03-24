package qpay

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var cacheInstance *cache.Cache

type client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

func InitQPay(clientID string, clientSecret string) {
	cacheInstance = cache.New(864000*time.Second, 1*time.Hour)
	cacheInstance.Set("client", client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client",
		RefreshToken: "",
	}, cache.NoExpiration)
}

func GetToken() (string, error) {
	token, err := cacheInstance.Get("token")
	if err != false {
		log.Println("Token has not found")
		if _, err := Authenticate(); err != nil {
			log.Println(err)
		}
	}

	return fmt.Sprintf("%v", token), nil
}
