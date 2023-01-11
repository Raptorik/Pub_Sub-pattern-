package main

import (
	"context"
	"log"
	"sync"
)

// create a Hub to accept multiple Subscribers,
// and the structure of this Hub structure is as follows.
type hub struct {
	sync.Mutex
	subs map[*subscriber]struct{}
}

func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}

// Initialize subcribers with map via newHub
type message struct {
	data []byte
}

type subscriber struct {
	sync.Mutex

	name    string        // the name of the subscriber
	handler chan *message // passing the message to handler through channel
	quit    chan struct{}
}

func (s *subscriber) Run(ctx context.Context) { //receive messages after a successful subscription
	for {
		select {
		case msg := <-s.handler:
			log.Println(s.name, string(msg.data))
		case <-s.quit:
			return
		case <-ctx.Done():
			return

		}
	}
}

func NewSubscriber(name string) *subscriber { //initialization of a single Subscriber.
	return &subscriber{
		name:    name,
		handler: make(chan *message, 100),
		quit:    make(chan struct{}),
	}
}
func (h *hub) Subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{}
	h.Unlock()

	go s.Run(ctx)

	return nil
}
