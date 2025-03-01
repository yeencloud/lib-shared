package shared

import (
	"fmt"
	"testing"
	"time"
)

// TestNewCustomContext checks if a new CustomContext is created correctly.
func TestNewCustomContext(t *testing.T) {
	ctx := NewCustomContext()
	if ctx == nil {
		t.Fatal("NewCustomContext() returned nil")
	}
}

// TestValue checks if values can be set and retrieved correctly.
func TestValue(t *testing.T) {
	ctx := NewCustomContext()
	ctx.WithValue("key1", "value1")
	ctx.WithValue("key2", 42)

	if ctx.Value("key1") != "value1" {
		t.Errorf("Expected 'value1', got %v", ctx.Value("key1"))
	}

	if ctx.Value("key2") != 42 {
		t.Errorf("Expected 42, got %v", ctx.Value("key2"))
	}

	if ctx.Value("nonexistent") != nil {
		t.Errorf("Expected nil for nonexistent key, got %v", ctx.Value("nonexistent"))
	}
}

// TestDeadline checks if the Deadline method behaves correctly.
func TestDeadline(t *testing.T) {
	ctx := NewCustomContext()
	deadline, ok := ctx.Deadline()

	if !deadline.IsZero() {
		t.Errorf("Expected zero time, got %v", deadline)
	}

	if ok {
		t.Errorf("Expected false, got true")
	}
}

// TestDone checks if Done() returns a closed channel after cancellation.
func TestDone(t *testing.T) {
	ctx := NewCustomContext()

	// Ensure channel is open before canceling
	select {
	case <-ctx.Done():
		t.Errorf("Done() channel should not be closed before canceling")
	default:
	}

	// Cancel context
	ctx.Cancel(fmt.Errorf("test error"))

	// Ensure channel is closed after canceling
	select {
	case <-ctx.Done():
		// Expected behavior, pass test
	default:
		t.Errorf("Done() channel should be closed after cancel")
	}
}

// TestErr checks if Err() returns the correct error after cancellation.
func TestErr(t *testing.T) {
	ctx := NewCustomContext()

	if ctx.Err() != nil {
		t.Errorf("Expected nil before cancellation, got %v", ctx.Err())
	}

	expectedErr := fmt.Errorf("custom error")
	ctx.Cancel(expectedErr)

	if ctx.Err() != expectedErr {
		t.Errorf("Expected '%v', got '%v'", expectedErr, ctx.Err())
	}
}

// TestCancel checks if Cancel() properly closes the context and sets the error.
func TestCancel(t *testing.T) {
	ctx := NewCustomContext()
	ctx.Cancel(fmt.Errorf("first error"))

	err1 := ctx.Err()
	if err1 == nil || err1.Error() != "first error" {
		t.Errorf("Expected 'first error', got %v", err1)
	}

	// Calling Cancel again should not overwrite the error
	ctx.Cancel(fmt.Errorf("second error"))
	err2 := ctx.Err()
	if err2 == nil || err2.Error() != "first error" {
		t.Errorf("Expected 'first error' to persist, but got %v", err2)
	}
}

// TestConcurrentAccess ensures that concurrent reads/writes do not cause race conditions.
func TestConcurrentAccess(t *testing.T) {
	ctx := NewCustomContext()
	numGoroutines := 100

	done := make(chan struct{})

	// Writing values concurrently
	go func() {
		for i := 0; i < numGoroutines; i++ {
			ctx.WithValue(fmt.Sprintf("key-%d", i), i)
		}
		close(done)
	}()

	// Simultaneously reading values
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			_ = ctx.Value(fmt.Sprintf("key-%d", i))
		}(i)
	}

	// Ensure writing is complete
	<-done
}

// TestRange checks if values can be iterated over correctly.
func TestRange(t *testing.T) {
	ctx := NewCustomContext()
	ctx.WithValue("user", "Alice")
	ctx.WithValue("role", "admin")
	ctx.WithValue("request_id", "12345")

	// Store values during iteration
	iteratedValues := make(map[interface{}]interface{})

	ctx.Range(func(key, value interface{}) bool {
		iteratedValues[key] = value
		return true
	})

	// Verify iteration captured all values
	expected := map[interface{}]interface{}{
		"user":       "Alice",
		"role":       "admin",
		"request_id": "12345",
	}

	for key, expectedValue := range expected {
		if iteratedValues[key] != expectedValue {
			t.Errorf("Expected %v, got %v", expectedValue, iteratedValues[key])
		}
	}
}

// TestRangeStopEarly checks if iteration can be stopped early.
func TestRangeStopEarly(t *testing.T) {
	ctx := NewCustomContext()
	ctx.WithValue("key1", "value1")
	ctx.WithValue("key2", "value2")
	ctx.WithValue("key3", "value3")

	count := 0

	ctx.Range(func(key, value interface{}) bool {
		count++
		return count < 2 // Stop after 2 iterations
	})

	if count != 2 {
		t.Errorf("Expected to iterate twice, but got %d iterations", count)
	}
}

func TestCancelAfterTimeout(t *testing.T) {
	ctx := NewCustomContext()

	// Simulate a timeout by canceling after 100ms
	time.AfterFunc(100*time.Millisecond, func() {
		ctx.Cancel(fmt.Errorf("timeout occurred"))
	})

	// Wait for cancellation
	<-ctx.Done()

	if ctx.Err().Error() != "timeout occurred" {
		t.Errorf("Expected 'timeout occurred', got %v", ctx.Err())
	}
}
