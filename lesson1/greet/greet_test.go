package greet

import "testing"

func TestGreet(t *testing.T) {
	expected_output := "Hello, Azhibek! \n"
	text := Greet("Azhibek")

	if text != expected_output {
		t.Errorf("Greetings was incorrect, got: %s, expected: %s", text, expected_output)
	}
}
