package main

import (
	"net/http"
	"os"
	"os/signal"

	"brilla/internal/database"
	"brilla/internal/router"
)

func main() {
	dbClient := database.ConnectToDB()

	database.CreateBD(dbClient)

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
