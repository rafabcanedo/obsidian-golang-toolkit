package testutil

import "testing"

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func EqualError(t *testing.T, err error, want string) {
	t.Helper()
	Error(t, err)

	if err.Error() != want {
		t.Errorf("error message = %q; want %q", err.Error(), want)
	}
}
