package main

import "context"

func (h *hub) Publish(ctx context.Context, msg *message) error {
	h.Lock()
	for s := range h.subs {
		s.Publish(ctx, msg)
	}
	h.Unlock()

	return nil
}
func (s *subscriber) Publish(ctx context.Context, msg *message) {
	select {
	case <-ctx.Done():
		return
	case s.handler <- msg:
	default:
	}
}
func (h *hub) UnSubscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	delete(h.subs, s)
	h.Unlock()
	close(s.quit)
	return nil
}
