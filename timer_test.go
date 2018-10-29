package go_log

import "testing"

func TestTimer(t *testing.T) {
	ConfigureTimer()
	if IsTimeEnabled != true {
		t.Errorf("Expected valued of isTimerEnabled was true but got false")
	}
}
