package main

// IsCallingFunctionEnabled : Exported to be get to know the conifguration
var IsCallingFunctionEnabled bool

// ConfigureCallingFunction : Exported to be able to call it
func ConfigureCallingFunction() {
	IsCallingFunctionEnabled = true
}
