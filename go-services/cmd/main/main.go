package main

import (
	"log"
	"net/http"

	"github.com/canergulay/goservices/server/chat"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// PG CONNECTION WAS DISCARDED !
	// THE MESSAGE PERSISTANCE WILL BE HANDLED IN THE MAIN SERVICE WHICH IS ACTUALLY API GATEWAY
	// INITIALLY, I THOUGHT IT MIGHT BE A GOOD IDEA TO HAVE POSTGRESQL FOR THE MESSAGES
	// THEN I FACE WITH A PROBLEM THAT MY MAIN DATABASE IS MONGODB AND USER PROFILES ARE LOCATED IN A MONGODB COLLECTION.
	// SO WHEN I FETCH THE MESSAGES FROM HERE,POSTGRESQL DE-NORMALIZATION WOULD COME INTO PLAY SINCE I ONLY HAVE USER IDS HERE NOT THE PROFILES.
	// SO AGAIN, FOR THE SAKE OF SIMPLICITY, MESSAGES WILL ALSO BE PERSISTED IN MAIN DATABASE,MONGODB.
	// THE COMMUNICATION BETWEEN TWO SERVERS WILL BE CONDUCTED VIA gRPC WHICH MAKES THIS REPO IMPORTANT FOR ME.
	// THIS IS THE FIRST EVER PROJECT FOR ME THAT I EVER COME UP WITH AN MICROSERVICES APPROACH WHERE I USED TWO DIFFERENT BACKEND TECHNOLOGIES
	// WITH A gRPC BRIDGE BETWEEN THEM.

	// pgConnection := global.InitPostgreSQL()

	socketPool := chat.InitializeSocketPool()
	socketServer := chat.InitializeSocketServer(socketPool)

	http.HandleFunc("/", socketServer.WebsocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
