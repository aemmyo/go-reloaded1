package main

import "testing"

func TestHex(t *testing.T) {
	input := "1E (hex)"
	expected := "30"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestUpSingle(t *testing.T) {
	input := "hello (up)"
	expected := "HELLO"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestLowSingle(t *testing.T) {
	input := "HELLO (low)"
	expected := "hello"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestCapSingle(t *testing.T) {
	input := "hello (cap)"
	expected := "Hello"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestUpMultiple(t *testing.T) {
	input := "this is fun (up, 2)"
	expected := "this IS FUN"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestLowMultiple(t *testing.T) {
	input := "THIS IS FUN (low, 2)"
	expected := "THIS is fun"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestCapMultiple(t *testing.T) {
	input := "welcome to the brooklyn bridge (cap, 2)"
	expected := "welcome to the Brooklyn Bridge"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFixAtoAn(t *testing.T) {

	input := "This is a apple"
	expected := "This is an apple"

	result := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFixPunctuation(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"Hello , world !", "Hello, world!"},
		{"This is good .", "This is good."},
		{"What is this ?", "What is this?"},
		{"Wait ; stop !", "Wait; stop!"},
		{"Go : the language", "Go: the language"},
	}

	for _, test := range tests {

		result := ProcessText(test.input)

		if result != test.expected {
			t.Errorf("Input: %q Expected: %q Got: %q", test.input, test.expected, result)
		}
	}
}

func TestFixQuotes(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"' hello '", "'hello'"},
		{"This is ' a test '", "This is 'a test'"},
		{"He said ' hello world '", "He said 'hello world'"},
		{"' Go ' is nice", "'Go' is nice"},
		{"No quotes here", "No quotes here"},
	}

	for _, test := range tests {

		result := ProcessText(test.input)

		if result != test.expected {
			t.Errorf("Input: %q Expected: %q Got: %q", test.input, test.expected, result)
		}
	}
}
