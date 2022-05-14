package main

import (
	"microservices/order-aggregator/internal/config"
	"microservices/order-aggregator/internal/container"
	"os"

	"gitlab.com/pos-alfa-microservices-go/core/log"
)

func main() {
	log.Logger.Info("Starting order-aggregator ...")

	conf, err := config.Start()
	if err != nil {
		log.Logger.Fatal("", err)
	}

	container := container.NewContainer(conf)
	if err := container.Start(); err != nil {
		log.Logger.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		log.Logger.Info("signal received %v", sig)
		done <- true
	}()

	<-done

}
