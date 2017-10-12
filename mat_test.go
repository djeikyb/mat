package main

import "testing"

func Test_name(t *testing.T) {
	input := "bar/foo.jpg"
	expected := "bar/foo_mat_4x6.jpg"
	actual := name(input)
	if actual != expected {
		t.Fatal(describe("wrong name", expected, actual))
	}
}
