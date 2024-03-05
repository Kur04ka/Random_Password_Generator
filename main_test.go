package main

import "testing"

func TestIsNeeded(t *testing.T) {
	t.Run("Group isNeeded", func(t *testing.T) {
		t.Run("isNeeded false test", func(t *testing.T) {
			expectedResult := false
			actualResult := isNeeded("n", "")
			if actualResult != expectedResult {
				t.Errorf("actual result: %t doesn't match expected result: %t", actualResult, expectedResult)
			}
		})

		t.Run("isNeeded true test", func(t *testing.T) {
			expectedResult := true
			actualResult := isNeeded("y", "")
			if actualResult != expectedResult {
				t.Errorf("actual result: %t doesn't match expected result: %t", actualResult, expectedResult)
			}
		})

		t.Run("isNeeded unknown letter test", func(t *testing.T) {
			expectedResult := false
			actualResult := isNeeded("q", "")
			if actualResult != expectedResult {
				t.Errorf("actual result: %t doesn't match expected result: %t", actualResult, expectedResult)
			}
		})
	})
}
