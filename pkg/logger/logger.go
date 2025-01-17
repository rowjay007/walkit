package logger

import (
    "fmt"
    "io"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
    *zerolog.Logger
}

type Config struct {
    OutputPaths []string
    LogLevel    zerolog.Level
    Development bool
}

func New() *Logger {
    return NewWithConfig(Config{
        OutputPaths: []string{"stdout"},
        LogLevel:    zerolog.InfoLevel,
        Development: os.Getenv("APP_ENV") != "production",
    })
}

func NewWithConfig(config Config) *Logger {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

    var writers []io.Writer
    for _, path := range config.OutputPaths {
        if path == "stdout" {
            if config.Development {
                writers = append(writers, zerolog.ConsoleWriter{
                    Out:        os.Stdout,
                    TimeFormat: time.RFC3339,
                    NoColor:    false,
                })
            } else {
                writers = append(writers, os.Stdout)
            }
        } else {
            file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                fmt.Printf("Failed to open log file %s: %v\n", path, err)
                continue
            }
            writers = append(writers, file)
        }
    }

    var writer io.Writer
    if len(writers) > 1 {
        writer = io.MultiWriter(writers...)
    } else if len(writers) == 1 {
        writer = writers[0]
    } else {
        writer = os.Stdout
    }

    logger := zerolog.New(writer).
        Level(config.LogLevel).
        With().
        Timestamp().
        Caller().
        Logger()

    return &Logger{&logger}
}

func (l *Logger) GinLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        raw := c.Request.URL.RawQuery
        if raw != "" {
            path = path + "?" + raw
        }

        c.Next()

        end := time.Now()
        latency := end.Sub(start)

        msg := "Request"
        if len(c.Errors) > 0 {
            msg = c.Errors.String()
        }

        l.Info(msg,
            map[string]interface{}{
                "method":     c.Request.Method,
                "path":       path,
                "status":     c.Writer.Status(),
                "latency":    latency,
                "ip":         c.ClientIP(),
                "user-agent": c.Request.UserAgent(),
                "body-size":  c.Writer.Size(),
            })
    }
}

// Consolidated Info method to handle fields
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
    event := l.Logger.Info()

    // Include fields in the log
    for _, field := range fields {
        for key, value := range field {
            event = event.Interface(key, value)
        }
    }

    event.Msg(msg)
}

func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
    event := l.Logger.Debug()
    for _, field := range fields {
        for key, value := range field {
            event = event.Interface(key, value)
        }
    }
    event.Msg(msg)
}

func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
    event := l.Logger.Warn()
    for _, field := range fields {
        for key, value := range field {
            event = event.Interface(key, value)
        }
    }
    event.Msg(msg)
}

func (l *Logger) Error(err error, msg string, fields ...map[string]interface{}) {
    event := l.Logger.Error().Err(err)
    for _, field := range fields {
        for key, value := range field {
            event = event.Interface(key, value)
        }
    }
    event.Msg(msg)
}

func (l *Logger) Fatal(msg string, fields ...map[string]interface{}) {
    event := l.Logger.Fatal()
    for _, field := range fields {
        for key, value := range field {
            event = event.Interface(key, value)
        }
    }
    event.Msg(msg)
}
