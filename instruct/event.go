package instruct

import "time"

//参考HTML标准

type Event interface {
	Type() string
	Target() EventTarget
	Prevent() error
	TimeStamp() time.Time
}

type EventListener func(event Event)

type EventTarget interface {
	DispatchEvent(event Event) error
	AddEventListener(listener EventListener) error
	RemoveEventListener(listener EventListener) error
}