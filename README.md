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

	// Panic logging: We want to crash
	Panic(event string, keyvals ...interface{})

	// Context extending interface
	With(keyvals ...interface{}) Logger
}
```

It was designed for the [ftpserverlib](https://github.com/fclairamb/ftpserverlib).

## Compatible logging frameworks

### [go-kit/log](https://github.com/go-kit/log)
```golang

import (
	log "github.com/fclairamb/go-log"
	gokit "github.com/fclairamb/go-log/logkit"
)

func main() {
	var logger log.Logger
	logger = logkit.New()

	logger.Info("Hello world")
}
```

### [log15](https://github.com/inconshreveable/log15)
```golang

import (
	log "github.com/fclairamb/go-log"
	log15 "github.com/fclairamb/go-log/log15"
)

func main() {
	var logger log.Logger
	logger = log15.New()

	logger.Info("Hello world")
}
```

### [zap](https://github.com/uber-go/zap)
```golang
import (
	log "github.com/fclairamb/go-log"
	zap "github.com/fclairamb/go-log/zap"
)

func main() {
	var logger log.Logger
	logger = zap.New()

	logger.Info("Hello world")
}
```

### [zerolog](https://github.com/rs/zerolog/)
```golang
import (
	log "github.com/fclairamb/go-log"
	zap "github.com/fclairamb/go-log/zerolog"
)

func main() {
	var logger log.Logger
	logger = zerolog.New()

	logger.Info("Hello world")
}
```

### [logrus](https://github.com/sirupsen/logrus) (not recommended)
```golang
import (
	log "github.com/fclairamb/go-log"
	zap "github.com/fclairamb/go-log/logrus"
)

func main() {
	var logger log.Logger
	logger = logrus.New()

	logger.Info("Hello world")
}
```
