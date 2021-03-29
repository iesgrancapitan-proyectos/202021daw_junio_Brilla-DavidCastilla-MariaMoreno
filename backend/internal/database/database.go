package database

import (
	"context"

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

func CreateBD(client driver.Client) {
	bd, err := client.CreateDatabase(context.Background(), "bd_brilla", nil)

	if err != nil {
		bd, err = client.Database(context.Background(), "bd_brilla")

		if err != nil {
			panic(err)
		}
	}

	bd.CreateCollection(context.Background(), "User", nil)

	bd.CreateCollection(context.Background(), "Brillo", nil)

	bd.CreateCollection(context.Background(), "Follows", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	bd.CreateCollection(context.Background(), "Author", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	bd.CreateCollection(context.Background(), "Comments", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	bd.CreateCollection(context.Background(), "Interactions", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	bd.CreateCollection(context.Background(), "Rebrilla", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})
}
