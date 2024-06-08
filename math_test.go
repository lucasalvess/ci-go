package main

import "testing"

func TestSum(t *testing.T) {

	total := sumNumber(15, 15)

	if total != 30 {
		t.Errorf("Result of sum is invalid: Result %d. Expected: %d", total, 30)
	}
}
