package testing

import (
	"testing"
	"time"

	"github.com/youpele52/lazysetup/pkg/models"
)

// MockState returns a fresh State instance for testing
func MockState() *models.State {
	return models.NewState()
}

// AssertEqual fails test if want != got
func AssertEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

// AssertNotEqual fails test if want == got
func AssertNotEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if want == got {
		t.Errorf("expected %v != %v", want, got)
	}
}

// AssertTrue fails test if !value
func AssertTrue(t *testing.T, value bool) {
	t.Helper()
	if !value {
		t.Errorf("expected true, got false")
	}
}

// AssertFalse fails test if value
func AssertFalse(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("expected false, got true")
	}
}

// AssertNil fails test if value != nil
func AssertNil(t *testing.T, value interface{}) {
	t.Helper()
	if value != nil {
		t.Errorf("expected nil, got %v", value)
	}
}

// AssertNotNil fails test if value == nil
func AssertNotNil(t *testing.T, value interface{}) {
	t.Helper()
	if value == nil {
		t.Errorf("expected non-nil value")
	}
}

// WaitForCondition waits up to duration for condition to become true
func WaitForCondition(t *testing.T, condition func() bool, timeout time.Duration, message string) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if condition() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Errorf("%s: timeout after %v", message, timeout)
}
