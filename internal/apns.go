package internal

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"log"
)

type Provider struct {
	KeyId   string
	TeamId  string
	KeyFile string
	Topic   string
	Client  *apns2.Client
}

func (provider *Provider) SetupProvider() *Provider {
	authKey, err := token.AuthKeyFromFile(provider.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Key loaded")

	newToken := &token.Token{
		AuthKey: authKey,
		KeyID:   provider.KeyId,
		TeamID:  provider.TeamId,
	}

	provider.Client = apns2.NewTokenClient(newToken).Development()

	log.Println("Provider client setup done")

	return provider
}

func (provider *Provider) SendNotification(deviceToken string, message string) (*apns2.Response, error) {
	notification := &apns2.Notification{
		DeviceToken: deviceToken,
		Topic:       provider.Topic,
		Payload:     []byte(fmt.Sprintf(`{"aps":{"alert":"%s"}}`, message)),
	}

	res, err := provider.Client.Push(notification)
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}
