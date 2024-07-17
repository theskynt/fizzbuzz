package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("input 1 return 1", func(t *testing.T) {
		input := 1
		want := "1"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("want %v but got %v", input, want)
		}
	})
}
