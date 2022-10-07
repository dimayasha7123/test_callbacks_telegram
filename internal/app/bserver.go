package app

import (
	"net/http"

	"github.com/dimayasha7123/test_callbacks_telegram/internal/models/users"
)

type bserver struct {
	apiKey     string
	users      users.SyncMap
	httpClient http.Client
}

func New(apiKey string) *bserver {
	return &bserver{
		apiKey: "",
		users:  *users.New(),
	}
}
