package test

import "testing"

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("err: %e", err)
	}
}

func Equal[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
