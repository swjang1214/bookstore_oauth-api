Skip to content
Search or jump to…

Pulls
Issues
Marketplace
Explore
 
@swjang1214 
Learn Git and GitHub without any code!
Using the Hello World guide, you’ll start a branch, write comments, and open a pull request.


federicoleon
/
bookstore_utils-go
2
717
Code
Issues
Pull requests
Actions
Projects
Wiki
Security
Insights
bookstore_utils-go/logger/logger.go /
@federicoleon
federicoleon using modules instead of dep
Latest commit a9b52b6 on 6 Apr 2020
 History
 1 contributor
96 lines (82 sloc)  1.69 KB
  
package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type bookstoreLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	output := strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}

func GetLogger() bookstoreLogger {
	return log
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}

// package logger

// import (
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// var (
// 	log *zap.Logger
// )

// func init() {
// 	colorFlag := false
// 	var encodeLevel zapcore.LevelEncoder = zapcore.LowercaseLevelEncoder
// 	if colorFlag {
// 		encodeLevel = zapcore.CapitalColorLevelEncoder
// 	}
// 	logConfig := zap.Config{
// 		OutputPaths: []string{"stdout"},
// 		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
// 		Encoding:    "json",
// 		EncoderConfig: zapcore.EncoderConfig{
// 			LevelKey:     "level",
// 			TimeKey:      "time",
// 			MessageKey:   "msg",
// 			EncodeTime:   zapcore.ISO8601TimeEncoder,
// 			EncodeLevel:  encodeLevel,
// 			EncodeCaller: zapcore.ShortCallerEncoder,
// 		},
// 	}
// 	var err error
// 	if log, err = logConfig.Build(); err != nil {
// 		panic(err)
// 	}
// }

// func GetLogger() *zap.Logger {
// 	return log
// }

// func Info(msg string, tags ...zap.Field) {
// 	log.Info(msg, tags...)
// 	log.Sync()
// }

// func Error(msg string, err error, tags ...zap.Field) {
// 	if err != nil {
// 		tags = append(tags, zap.NamedError("error", err))
// 	}
// 	log.Error(msg, tags...)
// 	log.Sync()
// }
