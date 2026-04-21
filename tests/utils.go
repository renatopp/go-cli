package tests

import (
	"fmt"
	"strings"
	"testing"
)

func make_args(v ...any) []string {
	var s []string
	for _, v := range v {
		s = append(s, fmt.Sprintf("%v", v))
	}
	return s
}

func printfContains(t *testing.T, expected string, message ...string) func(msg string, args ...any) {
	return func(msg string, args ...any) {
		msg = fmt.Sprintf(msg, args...)
		if !strings.Contains(msg, expected) {
			if len(message) > 0 {
				t.Fatalf("expected print to contain %q but got %q. %s", expected, msg, message[0])
			} else {
				t.Fatalf("expected print to contain %q but got %q", expected, msg)
			}
		}
	}
}

func expectPanicWith(t *testing.T, f func(), v any, message ...string) {
	defer func() {
		if r := recover(); r == nil {
			if len(message) > 0 {
				t.Fatalf("expected panic but did not panic. %s", message[0])
			} else {
				t.Fatal("expected panic but did not panic")
			}
		} else if r != v {
			t.Fatalf("expected panic with '%v' but got '%v'", v, r)
		}
	}()
	f()
}

func assert(t *testing.T, condition bool, message ...string) {
	if !condition {
		if len(message) > 0 {
			t.Fatalf("assertion failed: %s", message[0])
		} else {
			t.Fatal("assertion failed")
		}
	}
}

func assertEqual(t *testing.T, expected, actual any, message ...string) {
	if expected != actual {
		if len(message) > 0 {
			t.Fatalf("assertion failed: expected %v, got %v. %s", expected, actual, message[0])
		} else {
			t.Fatalf("assertion failed: expected %v, got %v", expected, actual)
		}
	}
}

func assertNotEqual(t *testing.T, expected, actual any, message ...string) {
	if expected == actual {
		if len(message) > 0 {
			t.Fatalf("assertion failed: expected not %v, but got %v. %s", expected, actual, message[0])
		} else {
			t.Fatalf("assertion failed: expected not %v, but got %v", expected, actual)
		}
	}
}

func assertTrue(t *testing.T, condition bool, message ...string) {
	if !condition {
		if len(message) > 0 {
			t.Fatalf("assertion failed: expected true but got false. %s", message[0])
		} else {
			t.Fatal("assertion failed: expected true but got false")
		}
	}
}

func assertFalse(t *testing.T, condition bool, message ...string) {
	if condition {
		if len(message) > 0 {
			t.Fatalf("assertion failed: expected false but got true. %s", message[0])
		} else {
			t.Fatal("assertion failed: expected false but got true")
		}
	}
}

func assertNil(t *testing.T, value any, message ...string) {
	if value != nil {
		if len(message) > 0 {
			t.Fatalf("assertion failed: expected nil but got %v. %s", value, message[0])
		} else {
			t.Fatalf("assertion failed: expected nil but got %v", value)
		}
	}
}

func assertNotNil(t *testing.T, value any, message ...string) {
	if value == nil {
		if len(message) > 0 {
			t.Fatalf("assertion failed: expected not nil but got nil. %s", message[0])
		} else {
			t.Fatal("assertion failed: expected not nil but got nil")
		}
	}
}

func assertError(t *testing.T, err error, message ...string) {
	if err == nil {
		if len(message) > 0 {
			t.Fatalf("expected an error but got nil. %s", message[0])
		} else {
			t.Fatal("expected an error but got nil")
		}
	}
}

func assertNoError(t *testing.T, err error, message ...string) {
	if err != nil {
		if len(message) > 0 {
			t.Fatalf("expected no error but got: %v. %s", err, message[0])
		} else {
			t.Fatalf("expected no error but got: %v", err)
		}
	}
}

func assertPanic(t *testing.T, f func(), message ...string) {
	defer func() {
		if r := recover(); r == nil {
			if len(message) > 0 {
				t.Fatalf("expected panic but did not panic. %s", message[0])
			} else {
				t.Fatal("expected panic but did not panic")
			}
		}
	}()
	f()
}

func assertNotPanic(t *testing.T, f func(), message ...string) {
	defer func() {
		if r := recover(); r != nil {
			if len(message) > 0 {
				t.Fatalf("expected no panic but got: %v. %s", r, message[0])
			} else {
				t.Fatalf("expected no panic but got: %v", r)
			}
		}
	}()
	f()
}
