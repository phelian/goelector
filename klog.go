package goelector

import (
	"github.com/go-logr/logr"
	"k8s.io/klog/v2"
)

// TurnOffKlog supresses leaderelector.go's usage of klog
func TurnOffKlog() {
	klog.SetLogger(&EmptyLogger{})
}

// EmptyLogger implements logr.Logging
type EmptyLogger struct{}

// Enabled always returns false
func (e *EmptyLogger) Enabled() bool {
	return false
}

// Info does nothing
func (e *EmptyLogger) Info(msg string, keysAndValues ...interface{}) {}

// Error does nothing
func (e *EmptyLogger) Error(err error, msg string, keysAndValues ...interface{}) {}

// V returns itself
func (e *EmptyLogger) V(level int) logr.Logger {
	return e
}

// WithValues returns itself
func (e *EmptyLogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	return e
}

// WithName returns itself
func (e *EmptyLogger) WithName(name string) logr.Logger {
	return e
}
