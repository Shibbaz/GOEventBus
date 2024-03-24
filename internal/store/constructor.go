package store

import (
	. "events"
)

func NewEventsStore(dispatcher EventDispatcher) *EventsStore {
	return &EventsStore{
		Stream:     []Event{},
		Dispatcher: dispatcher,
	}
}
