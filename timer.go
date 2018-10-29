package main

// IsTimeEnabled : Exported to be get to know the configuration
var IsTimeEnabled bool

// ConfigureTimer : Exported to be able to call it
func ConfigureTimer() {
	IsTimeEnabled = true
}
