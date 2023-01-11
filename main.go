package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	h := newHub()
	sub1 := NewSubscriber("Subscriber #1")
	sub2 := NewSubscriber("Subscriber #2")
	sub3 := NewSubscriber("Subscriber #3")

	h.Subscribe(ctx, sub1)
	h.Subscribe(ctx, sub2)
	h.Subscribe(ctx, sub3)

	_ = h.Publish(ctx, &message{data: []byte("Studying golang course during combat duty")})
	_ = h.Publish(ctx, &message{data: []byte("Gopher emotional stories to be told at the smoking time")})
	_ = h.Publish(ctx, &message{data: []byte("How military intelligence officer can become an golang-coder")})
	time.Sleep(1 * time.Second)

	h.UnSubscribe(ctx, sub3)
	_ = h.Publish(ctx, &message{data: []byte("Is it possible to come along with all that complicated stuff?")})
	_ = h.Publish(ctx, &message{data: []byte("I guess i'll be the first and the luckiest to escape this matrix of military service")})
	time.Sleep(1 * time.Second)
}
