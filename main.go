package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-tweets-streamer/model"
	"github.com/go-tweets-streamer/stream_client"
	"github.com/go-tweets-streamer/utils"
)

func main() {

	creds := model.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}
	//Get http clinet to make twitter api requests :
	client := stream_client.Create_Streamer_Client(&creds)

	//Verify Credentails
	err := utils.Verify_User(client)
	if err != nil {
		log.Fatal(fmt.Printf("err: %v\n", err))
	}

	//demux := stream_client.Get_Demux()

	//Instatitae a channel to recieve app termination signal.
	interrupt_signal := make(chan os.Signal)
	signal.Notify(interrupt_signal, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Start Streaming Tweets")

	stream, err := stream_client.Start_Streamer(client)
	if err != nil {
		log.Fatal(fmt.Printf("err: %v\n", err))
		interrupt_signal <- syscall.SIGTERM
	}

	for message := range stream.Messages {
		fmt.Println(message)
	}

	//go demux.HandleChan(stream.Messages)

	//stream_client.Post_Tweet(client, "Test Tweet")
	//stream_client.Search_Tweet(client, "Test")

	log.Println(<-interrupt_signal)
	fmt.Println("Streaming Ends")
	stream.Stop()

}
