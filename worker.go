package main

import (
	"context"

	"github.com/KY2001/pubsub-worker/infrastructure/pubsub"
)

const subscriptionID = "pubsub-worker-poc"

func main() {
	ctx := context.Background()

	pubsub.InitClient()
	defer pubsub.CloseClient()

	pubsubWorker(ctx)
}

func pubsubWorker(ctx context.Context) {
	handler := pubsub.Handler
	for {
		pubsub.PullMessage(ctx, subscriptionID, handler)
	}
}
