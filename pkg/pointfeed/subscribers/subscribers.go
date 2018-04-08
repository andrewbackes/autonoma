// Package subscribers has interfaces for subscribing to a feed.
package subscribers

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

// Subscriber can subscribe.
type Subscriber interface {
	Subscribe(string, chan coordinates.Point)
}

// Unsubscriber can unsubscribe.
type Unsubscriber interface {
	Unsubscribe(string)
}

// SubscribeUnsubscriber can subscribe and unsubscribe.
type SubscribeUnsubscriber interface {
	Subscriber
	Unsubscriber
}
