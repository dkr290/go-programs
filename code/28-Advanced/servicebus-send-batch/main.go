package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func main() {

	// Read env variables from .env file if it exists

	// Set background context

	//ctx, cancel := context.WithCancel(context.Background())

	//defer cancel()

	connectionString := os.Getenv("SERVICEBUS_CONNECTION_STRING")
	tTopic := os.Getenv("TOPIC")

	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		panic(err)
	}
	sender, err := client.NewSender(tTopic, nil)

	if err != nil {
		panic(err)
	}

	// close the sender when it's no longer needed
	defer sender.Close(context.TODO())

	// Create a message batch. It will automatically be sized for the Service Bus
	// Namespace's maximum message size.
	currentMessageBatch, err := sender.NewMessageBatch(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
	var messagesToSend []string
	for i := 0; i < 10000; i++ {
		messagesToSend = append(messagesToSend, "This_Is_test"+strconv.Itoa(i))

	}

	for i := 0; i < len(messagesToSend); i++ {
		// Add a message to our message batch. This can be called multiple times.
		err = currentMessageBatch.AddMessage(&azservicebus.Message{Body: []byte(messagesToSend[i])}, nil)

		if errors.Is(err, azservicebus.ErrMessageTooLarge) {
			if currentMessageBatch.NumMessages() == 0 {
				// This means the message itself is too large to be sent, even on its own.
				// This will require intervention from the user.
				panic("Single message is too large to be sent in a batch.")
			}

			fmt.Printf("Message batch is full. Sending it and creating a new one.\n")

			// send what we have since the batch is full
			err := sender.SendMessageBatch(context.TODO(), currentMessageBatch, nil)

			if err != nil {
				panic(err)
			}

			// Create a new batch and retry adding this message to our batch.
			newBatch, err := sender.NewMessageBatch(context.TODO(), nil)

			if err != nil {
				panic(err)
			}

			currentMessageBatch = newBatch

			// rewind the counter and attempt to add the message again (this batch
			// was full so it didn't go out with the previous SendMessageBatch call).
			i--
		} else if err != nil {
			panic(err)
		}
	}

	// check if any messages are remaining to be sent.
	if currentMessageBatch.NumMessages() > 0 {
		err := sender.SendMessageBatch(context.TODO(), currentMessageBatch, nil)

		if err != nil {
			panic(err)
		}

	}
}
