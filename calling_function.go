package go_log

// IsCallingFunctionEnabled : Exported to be get to know the configuration
var IsCallingFunctionEnabled bool

// ConfigureCallingFunction : Exported to be able to call it
func ConfigureCallingFunction() {
	IsCallingFunctionEnabled = true
}
