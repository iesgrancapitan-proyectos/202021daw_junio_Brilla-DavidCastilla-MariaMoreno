package database

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"

	"brilla/internal/database/config"
)

func ConnectToDB() (client driver.Client) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.DB_URL},
	})
	if err != nil {
		panic(err)
	}
	client, err = driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		panic(err)
	}
	return
}

func CreateBD(client driver.Client) (db driver.Database) {

	// FIX: Better handling
	db, err := client.CreateDatabase(context.Background(), "bd_brilla", nil)
	if err != nil {
		db, err = client.Database(context.Background(), "bd_brilla")

		if err != nil {
			panic(err)
		}
	}

	db.CreateCollection(context.Background(), "User", nil)

	db.CreateCollection(context.Background(), "Brillo", nil)

	db.CreateCollection(context.Background(), "Follows", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateCollection(context.Background(), "Author", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateCollection(context.Background(), "Comments", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateCollection(context.Background(), "Interactions", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateCollection(context.Background(), "Rebrilla", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateGraph(context.Background(), "Follows", &driver.CreateGraphOptions{
		EdgeDefinitions: []driver.EdgeDefinition{
			{
				Collection: "Follows",
				To:         []string{"User"},
				From:       []string{"User"},
			},
		},
	})

	db.CreateGraph(context.Background(), "Comments", &driver.CreateGraphOptions{
		EdgeDefinitions: []driver.EdgeDefinition{
			{
				Collection: "Comments",
				To:         []string{"Brillo"},
				From:       []string{"Brillo"},
			},
		},
	})

	db.CreateGraph(context.Background(), "Interactions", &driver.CreateGraphOptions{
		EdgeDefinitions: []driver.EdgeDefinition{
			{
				Collection: "Interactions",
				To:         []string{"User"},
				From:       []string{"Brillo"},
			},
		},
	})

	db.CreateGraph(context.Background(), "Rebrilla", &driver.CreateGraphOptions{
		EdgeDefinitions: []driver.EdgeDefinition{
			{
				Collection: "Rebrilla",
				To:         []string{"User"},
				From:       []string{"Brillo"},
			},
		},
	})

	return
}
