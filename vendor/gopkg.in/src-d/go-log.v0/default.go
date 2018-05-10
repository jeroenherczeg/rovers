package log

import (
	"github.com/src-d/envconfig"
)

// DefaultFactory logger used by the `log.New()` function. Pre-configured with
// the environment variables.
//
// The environment variables are autogenerated from the `LoggerFactory` fields,
// with a prefix `log`, it means that three different variables are available
// `LOG_LEVEL`, `LOG_FORMAT`, `LOG_FIELDS`, for more information about the
// values please read the `LoggerFactory` documentation.
var DefaultFactory *LoggerFactory

// DefaultLogger logger used by the `log.Infof()`, `log.Debugf()`,
// `log.Warningf()` and `log.Error()` functions. The DefaultLogger is lazily
// instantiated if nil once this functions are called.
var DefaultLogger Logger

// New returns a new logger using `DefaultFactory`.
func New() (Logger, error) {
	if DefaultFactory == nil {
		DefaultFactory = &LoggerFactory{}
		if err := envconfig.Process("log", DefaultFactory); err != nil {
			return nil, err
		}
	}

	return DefaultFactory.New()
}

// Debugf logs a message at level Debug.
func Debugf(format string, args ...interface{}) {
	getLogger().Debugf(format, args...)
}

// Infof logs a message at level Info.
func Infof(format string, args ...interface{}) {
	getLogger().Infof(format, args...)
}

// Warningf logs a message at level Warning.
func Warningf(format string, args ...interface{}) {
	getLogger().Warningf(format, args...)

}

// Errorf logs an error with a message at level Error.
func Errorf(err error, format string, args ...interface{}) {
	getLogger().Errorf(err, format, args...)
}

func getLogger() Logger {
	if DefaultLogger == nil {
		var err error
		DefaultLogger, err = New()
		if err != nil {
			panic(err)
		}
	}

	return DefaultLogger
}
