package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("input 1 return 1", func(t *testing.T) {
		input := 1
		want := "1"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 2 return 2", func(t *testing.T) {
		input := 2
		want := "2"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 3 return Fizz", func(t *testing.T) {
		input := 3
		want := "Fizz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 4 return 4", func(t *testing.T) {
		input := 4
		want := "4"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 5 return Buzz", func(t *testing.T) {
		input := 5
		want := "Buzz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 6 return Fizz", func(t *testing.T) {
		input := 6
		want := "Fizz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 7 return 7", func(t *testing.T) {
		input := 7
		want := "7"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 8 return 8", func(t *testing.T) {
		input := 8
		want := "8"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 9 return Fizz", func(t *testing.T) {
		input := 9
		want := "Fizz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 10 return Buzz", func(t *testing.T) {
		input := 10
		want := "Buzz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 11 return 11", func(t *testing.T) {
		input := 11
		want := "11"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 12 return Fizz", func(t *testing.T) {
		input := 12
		want := "Fizz"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})

	t.Run("input 13 return 13", func(t *testing.T) {
		input := 13
		want := "13"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want: %v but got %v", want, got)
		}
	})
}
