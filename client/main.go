package main

import (
	"context"
	"fmt"
	"messageRelayClient/infrastructure"
)

func main() {
	/**
	* register
	* subscribe to own topic
	* get active user ids
	* publish to whichever topics, include all the desired ids in the Receiver field
	 */
	userId := "b64eb839-ba73-4db4-a936-6ed56c66580a"
	ctx := context.Background()

	subscriber := infrastructure.NewSubscriber()
	pubsub := subscriber.Subscribe(ctx, userId)

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println("MESSAGE RECEIVED")
		fmt.Println(msg.Payload)
	}
}
