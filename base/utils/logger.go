package utils

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/rs/zerolog"
)

// Get Log Level From .env If Present, Else Use INFO As The Log Level
func getLogLevel() zerolog.Level {

	var err error
	var logLevel zerolog.Level

	logLevelString := GetEnv("LOG_LEVEL")

	if logLevelString == "" {
		return zerolog.InfoLevel
	}
	if logLevel, err = zerolog.ParseLevel(logLevelString); err != nil {
		return zerolog.InfoLevel
	}
	return logLevel
}

var Logger NewLogger
var singleton sync.Once
var logLevel = getLogLevel()

// Logger Struct To Have Some Extra Functionality On-Top Of `zerolog.Logger`
type NewLogger struct {
	Logger zerolog.Logger
}

// Returns NewLogger Instance
func GetLogger() NewLogger {

	logger := zerolog.New(os.Stdout).Level(logLevel).With().Timestamp().Int("PID", os.Getpid()).Logger()

	return NewLogger{Logger: logger}

}

// Overridden zerolog.Logger().Debug() Method
func (logger *NewLogger) Debug(message string) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Debug().Str("caller", callerInfo).Msg(message)

}

// Overridden zerolog.Logger().Debug() Method With Formatted Message
func (logger *NewLogger) DebugF(message string, values ...any) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Debug().Str("caller", callerInfo).Msg(fmt.Sprintf(message, values...))

}

// Overridden zerolog.Logger().Info() Method
func (logger *NewLogger) Info(message string) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Info().Str("caller", callerInfo).Msg(message)

}

// Overridden zerolog.Logger().Info() Method With Formatted Message
func (logger *NewLogger) InfoF(message string, values ...any) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Info().Str("caller", callerInfo).Msg(fmt.Sprintf(message, values...))

}

// Overridden zerolog.Logger().Warn() Method
func (logger *NewLogger) Warn(message string) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Warn().Str("caller", callerInfo).Msg(message)

}

// Overridden zerolog.Logger().Warn() Method With Formatted Message
func (logger *NewLogger) WarnF(message string, values ...any) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Warn().Str("caller", callerInfo).Msg(fmt.Sprintf(message, values...))

}

// Overridden zerolog.Logger().Warn().AnErr() Method
func (logger *NewLogger) WarnWithError(key string, exception error, message string) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Warn().Str("caller", callerInfo).AnErr(key, exception).Msg(message)

}

// Overridden zerolog.Logger().Warn().AnErr() Method With Formatted Message
func (logger *NewLogger) WarnWithErrorF(key string, exception error, message string, values ...any) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Warn().Str("caller", callerInfo).AnErr(key, exception).Msg(fmt.Sprintf(message, values...))

}

// Overridden zerolog.Logger().Error() Method To Capture Exception On Sentry
func (logger *NewLogger) Error(exception error, message string) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Error().Str("caller", callerInfo).Err(exception).Msgf(message)

}

// Overridden zerolog.Logger().Error() Method To Capture Exception On Sentry With Formatted Message
func (logger *NewLogger) ErrorF(exception error, message string, values ...any) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Error().Str("caller", callerInfo).Err(exception).Msgf(fmt.Sprintf(message, values...))

}

// Overridden zerolog.Logger().Fatal() Method To Capture Exception On Sentry And Then Create Panic
func (logger *NewLogger) Fatal(exception error, message string) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Fatal().Str("caller", callerInfo).Err(exception).Msgf(message)

	panic(exception)

}

// Overridden zerolog.Logger().Fatal() Method To Capture Exception On Sentry With Formatted Message And Then Create Panic
func (logger *NewLogger) FatalF(exception error, message string, values ...any) {

	_, file, line, _ := runtime.Caller(1)
	callerInfo := fmt.Sprintf("%s:%d", file, line)
	logger.Logger.Fatal().Str("caller", callerInfo).Err(exception).Msgf(fmt.Sprintf(message, values...))

	panic(exception)

}

// Init A Global Logger
func InitLogger() {
	singleton.Do(func() {
		Logger = GetLogger()
	})
}
