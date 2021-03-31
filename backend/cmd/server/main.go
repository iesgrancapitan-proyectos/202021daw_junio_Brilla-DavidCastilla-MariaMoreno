package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"brilla/internal/database"
	"brilla/internal/router"
)

func main() {
	dbClient := database.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	database := database.CreateBD(ctx, dbClient)
	cancel()

	router := router.New(database)

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
