package main

import "testing"

func TestXxx(t *testing.T) {
	actual := FibOptimal(8)
	if actual != 21 {
		t.Log(FibOptimal(8))
		t.Error("expected value was 21, value received was")
	}
}
