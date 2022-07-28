package main

import (
	"context"

	"fmt"

	"os"

	servicebus "github.com/Azure/azure-service-bus-go"

	"github.com/Azure/go-amqp"
)

func main() {

	// Read env variables from .env file if it exists

	// Set background context

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

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

	for {

		if err = subs.ReceiveOne(ctx, servicebus.HandlerFunc(func(ctx context.Context, m *servicebus.Message) error {

			fmt.Println(string(m.Data))

			return m.Complete(ctx)

		})); err != nil {

			if innerErr, ok := err.(*amqp.Error); ok && innerErr.Condition == "com.microsoft:timeout" {

				fmt.Println("Timeout waiting for messages. Entering next loop.")

				continue

			}

			fmt.Println(err)

			return

		}

	}

}
