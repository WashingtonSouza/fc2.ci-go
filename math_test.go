package main

import "testing"

func testSum(t *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		t.Errorf("The result is invalid: The result %d and the expected is %d", total, 30)
	}
}
