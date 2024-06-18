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
	
}