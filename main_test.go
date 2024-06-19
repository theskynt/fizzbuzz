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
		// Arrange
		input := 2
		want := "2"

		// Act
		got := fizzbuzz(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Fizz when input is 3", func(t *testing.T) {
		// Arrange
		input := 3
		want := "Fizz"

		// Act
		got := fizzbuzz(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 4 when input is 4", func(t *testing.T) {
		// Arrange
		input := 4
		want := "4"

		// Act
		got := fizzbuzz(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return Buzz when input is 5", func(t *testing.T) {
		// Arrange
		input := 4
		want := "4"

		// Act
		got := fizzbuzz(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})
}