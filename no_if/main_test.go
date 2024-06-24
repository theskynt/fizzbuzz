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

}