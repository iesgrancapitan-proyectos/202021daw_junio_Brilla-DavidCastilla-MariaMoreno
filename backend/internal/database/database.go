package database

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"

	"brilla/internal/database/config"
)

func ConnectToDB() driver.Client {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.DB_URL},
	})
	if err != nil {
		panic(err)
	}
	client, _ := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		panic(err)
	}
	return client
}
