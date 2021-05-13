package database

import (
	"context"
	"fmt"
	"log"
	"time"

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

func CreateBD(ctx context.Context, client driver.Client) (db driver.Database) {

	dbChan := make(chan driver.Database, 1)
	go func() {
		for {
			db, err := client.CreateDatabase(context.Background(), "bd_brilla", nil)
			if driver.IsConflict(err) {
				db, err = client.Database(context.Background(), "bd_brilla")
				if err != nil {
					fmt.Printf("\r%vRetrying to connect to database", time.Now().Format("2006/01/02 15:04:05"))
					continue
				}
				dbChan <- db
			} else if err != nil {
				fmt.Printf("\r%v Retrying to connect to database", time.Now().Format("2006/01/02 15:04:05"))
				continue
			}
			fmt.Println()
			log.Println("Success connecting to database")
			dbChan <- db
		}
	}()

	select {
	case db = <-dbChan:
	case <-ctx.Done():
		panic("Timeout trying to connect to DB")
	}

	coll, _ := db.CreateCollection(context.Background(), "User", nil)
	if coll != nil {
		coll.EnsurePersistentIndex(context.Background(), []string{"email"}, &driver.EnsurePersistentIndexOptions{
			Unique: true,
		})
	}
	db.CreateCollection(context.Background(), "Brillo", nil)
	db.CreateCollection(context.Background(), "DeactivatedUser", nil)

	db.CreateCollection(context.Background(), "Follows", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateCollection(context.Background(), "Interactions", &driver.CreateCollectionOptions{
		Type: driver.CollectionTypeEdge,
	})

	db.CreateCollection(context.Background(), "Rebrillo", &driver.CreateCollectionOptions{
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

	// db.CreateArangoSearchView(context.Background(), "BrillosView", &driver.ArangoSearchViewProperties{
	//     Links: driver.ArangoSearchLinks{
	//         "Brillo": driver.ArangoSearchElementProperties{
	//             }
	//     },
	// })

	return
}
