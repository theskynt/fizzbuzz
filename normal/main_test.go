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
		input := 5
		want := "Buzz"

		// Act
		got := fizzbuzz(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	// t.Run("should return Fizz when input is 6", func(t *testing.T) {
	// 	// Arrange
	// 	input := 6
	// 	want := "Fizz"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return 7 when input is 7", func(t *testing.T) {
	// 	// Arrange
	// 	input := 7
	// 	want := "7"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return 8 when input is 8", func(t *testing.T) {
	// 	// Arrange
	// 	input := 8
	// 	want := "8"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return Fizz when input is 9", func(t *testing.T) {
	// 	// Arrange
	// 	input := 9
	// 	want := "Fizz"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return Buzz when input is 10", func(t *testing.T) {
	// 	// Arrange
	// 	input := 10
	// 	want := "Buzz"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return 11 when input is 11", func(t *testing.T) {
	// 	// Arrange
	// 	input := 11
	// 	want := "11"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return Fizz when input is 12", func(t *testing.T) {
	// 	// Arrange
	// 	input := 12
	// 	want := "Fizz"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return 13 when input is 13", func(t *testing.T) {
	// 	// Arrange
	// 	input := 13
	// 	want := "13"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return 14 when input is 14", func(t *testing.T) {
	// 	// Arrange
	// 	input := 14
	// 	want := "14"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })

	// t.Run("should return FizzBuzz when input is 15", func(t *testing.T) {
	// 	// Arrange
	// 	input := 15
	// 	want := "FizzBuzz"

	// 	// Act
	// 	got := fizzbuzz(input)

	// 	// Assert
	// 	if got != want {
	// 		t.Errorf("want %v , but got %v", want, got)
	// 	}
	// })
}