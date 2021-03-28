package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/arangodb/go-driver"
	arahttp "github.com/arangodb/go-driver/http"

	"brilla/internal/router"
)

func main() {
	dbClient := connectToDB()

	router := router.New(&dbClient)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, os.Kill)

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			panic(err)
		}
	}()

	println("Service listening on :8080")

	<-signalChannel

	println("Closing")
}

func connectToDB() driver.Client {
	conn, err := arahttp.NewConnection(arahttp.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
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
