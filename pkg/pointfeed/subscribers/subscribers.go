// Package subscribers has interfaces for subscribing to a feed.
package subscribers

import (
	"github.com/andrewbackes/autonoma/pkg/point"
)

// Subscriber can subscribe.
type Subscriber interface {
	Subscribe(string, chan point.Point)
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
