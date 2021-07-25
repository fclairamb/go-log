# Tiny logging API

This library provides a basic API to be integrated with any logging framework. It
allows to use logging in your libraries without enforcing a particular logging
framework.

Here is it:
```golang
// Logger interface
type Logger interface {

	// Debug logging: Every details
	Debug(event string, keyvals ...interface{})

	// Info logging: Core events
	Info(event string, keyvals ...interface{})

	// Warning logging: Anything out of the ordinary but non-life threatening
	Warn(event string, keyvals ...interface{})

	// Error logging: Major issue
	Error(event string, keyvals ...interface{})

	// Context extending interface
	With(keyvals ...interface{}) Logger
}
```

It was designed for the [ftpserverlib](https://github.com/fclairamb/ftpserverlib).

It contains a [go-kit/log](https://github.com/go-kit/log) binding but can easily be 
adapted to any logging framework.
