package types

var EventMap = map[string]func(){}

// AddEvent registers a new event in the event map
func AddEvent(eventName string, handler func()) {
	EventMap[eventName] = handler
}

// ClearEvents clears the event map after each loop
func ClearEvents() {
	EventMap = map[string]func(){}
}
