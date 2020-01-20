package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"slack-notifier/kafka"
	"slack-notifier/services"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8000", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := services.NewNotifier()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := services.Endpoints{
		MessageEndpoint: services.MakeMessageEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("notifier is listening on port:", *httpAddr)
		handler := services.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	go func() {
		log.Println("Text Kafka consumer starting...")
		consumer := kafka.NewConsumer("localhost:9092", "text_topic", "text_topic_group", "text_topic_client_id", false, srv)
		consumer.Run()
	}()

	go func() {
		log.Println("Execution Kafka consumer starting...")
		consumer := kafka.NewConsumer("localhost:9092", "execution_topic", "execution_topic_group", "execution_topic_client_id", false, srv)
		consumer.Run()
	}()

	go func() {
		log.Println("Transaction Kafka consumer starting...")
		consumer := kafka.NewConsumer("localhost:9092", "transaction_topic", "transaction_topic_group", "transaction_topic_client_id", false, srv)
		consumer.Run()
	}()

	log.Fatalln(<-errChan)
}
