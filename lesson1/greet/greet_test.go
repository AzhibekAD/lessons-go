package greet

import "testing"

func TestGreet(t *testing.T) {
	expectedOutput := "Hello, Azhibek! Your table is 12!"
	text := Greet("Azhibek", "12")

	if text != expectedOutput {
		t.Errorf("Greetings was incorrect, got: %s, expected: %s", text, expectedOutput)
	}
}
