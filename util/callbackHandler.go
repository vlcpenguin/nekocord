package util

// Function that gets callback with message object with type interface{}
var Callbacks = map[string]func(map[string]interface{}){}

// Set a callback for certain events like MESSAGE_CREATE etc
func SetCallback(callbackType string, callback func(map[string]interface{})) {
	if Callbacks == nil {
		Callbacks = make(map[string]func(map[string]interface{}))
	}
	Callbacks[callbackType] = callback
}
