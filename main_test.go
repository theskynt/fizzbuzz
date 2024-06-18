package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("should return 1 when input is 1", func(t *testing.T) {
		input := 1
		output := fizzbuzz(input)
		want := "1"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return 2 when input is 2", func(t *testing.T) {
		input := 2
		output := fizzbuzz(input)
		want := "2"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return Fizz when input is 3", func(t *testing.T) {
		input := 3
		output := fizzbuzz(input)
		want := "Fizz"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return 4 when input is 4", func(t *testing.T) {
		input := 4
		output := fizzbuzz(input)
		want := "4"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return Buzz when input is 5", func(t *testing.T) {
		input := 5
		output := fizzbuzz(input)
		want := "Buzz"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return Fizz when input is 6", func(t *testing.T) {
		input := 6
		output := fizzbuzz(input)
		want := "Fizz"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return 7 when input is 7", func(t *testing.T) {
		input := 7
		output := fizzbuzz(input)
		want := "7"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return 8 when input is 8", func(t *testing.T) {
		input := 8
		output := fizzbuzz(input)
		want := "8"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return Fizz when input is 9", func(t *testing.T) {
		input := 9
		output := fizzbuzz(input)
		want := "Fizz"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})

	t.Run("should return Buzz when input is 10", func(t *testing.T) {
		input := 10
		output := fizzbuzz(input)
		want := "Buzz"
		if output != want {
			t.Errorf("want %v , but output %v", want, output)
		}
	})
}