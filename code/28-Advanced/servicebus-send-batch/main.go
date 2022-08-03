package main

import (
	"fmt"

	"os"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func main() {

	// Read env variables from .env file if it exists

	// Set background context

	//ctx, cancel := context.WithCancel(context.Background())

	//defer cancel()

	connStr := os.Getenv("SERVICEBUS_CONNECTION_STRING")

	tName := os.Getenv("TOPIC_NAME")

	sSubs := os.Getenv("SUBSCRIPTION_NAME")

	if connStr == "" || tName == "" || sSubs == "" {

		fmt.Println("FATAL: expected environment variables SERVICEBUS_CONNECTION_STRING or TOPIC_NAME oe SUBSCRIPTION_NAME not set")

		return

	}

	// Create a client to communicate with a Service Bus Namespace.

	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))

	if err != nil {

		fmt.Println(err)

		return

	}

	topic, err := ns.NewTopic(tName)

	if err != nil {

		fmt.Println(err)

		return

	}

	// Create topic receiver

	subs, err := topic.NewSubscription(sSubs)

	if err != nil {

		fmt.Println(err)

		return

	}

	messages := []string{"foo", "bar", "bazz", "buzz"}
	for _, msg := range messages {
		//topic.Send(ctx, servicebus.NewMessageFromString(msg))
		//subs.Send(ctx, servicebus.NewMessageFromString(msg))
		fmt.Println(msg)
	}

}
