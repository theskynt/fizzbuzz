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

	t.Run("input 2 return 2", func(t *testing.T) {
		input := 2
		want := "2"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("want %v but got %v", input, want)
		}
	})

	t.Run("input 3 return Fizz", func(t *testing.T) {
		input := 3
		want := "Fizz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("want %v but got %v", input, want)
		}
	})

	t.Run("input 4 return 4", func(t *testing.T) {
		input := 4
		want := "4"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("want %v but got %v", input, want)
		}
	})

	t.Run("input 5 return Buzz", func(t *testing.T) {
		input := 5
		want := "Buzz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("want %v but got %v", input, want)
		}
	})
}
