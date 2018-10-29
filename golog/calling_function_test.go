package golog

import "testing"

func TestConfigureCallingFunction(t *testing.T) {
	ConfigureCallingFunction()
	if IsCallingFunctionEnabled != true {
		t.Errorf("Expected valued of IsCallingFunctionEnabled was true but got false")
	}
}
