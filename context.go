package shared

import (
	"sync"
	"time"
)

// Context is our implementation of context.Context
type Context struct {
	mu     sync.RWMutex
	values map[interface{}]interface{}
	logs   []map[string]interface{}
	done   chan struct{}
	err    error
}

// NewCustomContext initializes a new custom context
func NewCustomContext() *Context {
	return &Context{
		values: make(map[interface{}]interface{}),
		done:   make(chan struct{}),
	}
}

// Deadline returns no deadline (false)
func (c *Context) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

// Done returns a channel that is closed when the context is canceled
func (c *Context) Done() <-chan struct{} {
	return c.done
}

// Err returns the error if the context is canceled
func (c *Context) Err() error {
	return c.err
}

// Value retrieves a value by key
func (c *Context) Value(key interface{}) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.values[key]
}

// WithValue sets a key-value pair
func (c *Context) WithValue(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

func (c *Context) WithLog(value map[string]interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.logs = append(c.logs, value)
}

func (c *Context) Logs() []map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.logs

}

// Range iterates over all key-value pairs
func (c *Context) Range(f func(key, value interface{}) bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k, v := range c.values {
		if !f(k, v) {
			break // Stop iteration if function returns false
		}
	}
}

// Cancel closes the context and sets an error
func (c *Context) Cancel(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.err = err
		close(c.done)
	}
}
