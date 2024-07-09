package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("input 1 return 1", func(t *testing.T) {
		input := 1
		want := "1"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want : %v but got %v", want, got)
		}
	})

	t.Run("input 2 return 2", func(t *testing.T) {
		input := 2
		want := "2"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want : %v but got %v", want, got)
		}
	})

	t.Run("input 3 return Fizz", func(t *testing.T) {
		input := 3
		want := "Fizz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want : %v but got %v", want, got)
		}
	})
}
