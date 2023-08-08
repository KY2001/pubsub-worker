package pubsub

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/pubsub"

	"github.com/KY2001/pubsub-worker/config"
)

var (
	client *pubsub.Client
)

func InitClient() {
	conf := config.NewConfig()
	ctx := context.Background()
	var err error
	client, err = pubsub.NewClient(ctx, conf.Pubsub.ProjectID)
	if err != nil {
		log.Fatalf("InitClient: Failed to initialize pubsub client: %v\n", err)
	}
}

func CloseClient() {
	if client != nil {
		client.Close()
	}
}

func GetClient() *pubsub.Client {
	if client == nil {
		InitClient()
		log.Println("GetClient: pubsub client should be initialized when server starts.")
	}
	return client
}

func PullMessage(ctx context.Context, subscriptionID string, handler MessageHandler) error {
	client := GetClient()
	sub := client.Subscription(subscriptionID)

	err := sub.Receive(ctx, handler)
	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	return nil
}

type MessageHandler func(context.Context, *pubsub.Message)

func Handler(ctx context.Context, msg *pubsub.Message) {
	err := os.WriteFile("./judge/source.py", msg.Data, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	cmd := exec.Command("./judge/judge.sh")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}
	log.Println(string(output))

	msg.Ack()
	log.Printf("Got message: %q\n", string(msg.Data))
}
