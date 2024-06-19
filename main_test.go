package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("should return 1 when input is 1", func(t *testing.T) {
		// Arrange
		input := 1
		want := "1"

		// Act
		got := fizzbuzz(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 2 when input is 2", func(t *testing.T) {
		input := 2
		got := fizzbuzz(input)
		want := "2"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Fizz when input is 3", func(t *testing.T) {
		input := 3
		got := fizzbuzz(input)
		want := "Fizz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 4 when input is 4", func(t *testing.T) {
		input := 4
		got := fizzbuzz(input)
		want := "4"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Buzz when input is 5", func(t *testing.T) {
		input := 5
		got := fizzbuzz(input)
		want := "Buzz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Fizz when input is 6", func(t *testing.T) {
		input := 6
		got := fizzbuzz(input)
		want := "Fizz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 7 when input is 7", func(t *testing.T) {
		input := 7
		got := fizzbuzz(input)
		want := "7"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 8 when input is 8", func(t *testing.T) {
		input := 8
		got := fizzbuzz(input)
		want := "8"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Fizz when input is 9", func(t *testing.T) {
		input := 9
		got := fizzbuzz(input)
		want := "Fizz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Buzz when input is 10", func(t *testing.T) {
		input := 10
		got := fizzbuzz(input)
		want := "Buzz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 11 when input is 11", func(t *testing.T) {
		input := 11
		got := fizzbuzz(input)
		want := "11"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Fizz when input is 12", func(t *testing.T) {
		input := 12
		got := fizzbuzz(input)
		want := "Fizz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 13 when input is 13", func(t *testing.T) {
		input := 13
		got := fizzbuzz(input)
		want := "13"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 14 when input is 14", func(t *testing.T) {
		input := 14
		got := fizzbuzz(input)
		want := "14"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Fizzbuzz when input is 15", func(t *testing.T) {
		input := 15
		got := fizzbuzz(input)
		want := "Fizzbuzz"
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})
}