package logger

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/Aboagye-Dacosta/shopBackend/internal/constants"
)

var (
	infoFile  *os.File
	errorFile *os.File
	InfoFile  = "app.log"
	ErrorFile = "error.log"
)

type AppLogger struct {
	InfoLogger *slog.Logger
	ErrLogger  *slog.Logger
}

type CustomHandler struct {
	slog.Handler
}

func (c CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	if val, ok := ctx.Value(constants.REQUEST_ID_KEY).(string); ok {
		r.AddAttrs(slog.String(string(constants.REQUEST_ID_KEY), val))
	}

	if val, ok := ctx.Value(constants.USER_ID_KEY).(string); ok {
		r.AddAttrs(slog.String(string(constants.USER_ID_KEY), val))
	}

	if val, ok := ctx.Value(constants.TRACE_ID_KEY).(string); ok {
		r.AddAttrs(slog.String(string(constants.TRACE_ID_KEY), val))
	}

	return c.Handler.Handle(ctx, r)
}

func Init() *AppLogger {
	var err error

	// Open info log file
	infoFile, err = os.OpenFile(InfoFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("failed to open info log file", "error", err)
		return nil
	}

	// Open error log file
	errorFile, err = os.OpenFile(ErrorFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("failed to open error log file", "error", err)
		infoFile.Close() // Clean up info file if error file fails
		return nil
	}

	// Create handlers - InfoLog writes to info file, ErrLog writes to error file
	// But ErrLog can also write errors to both files if needed (see below)
	multiInfoLog := io.MultiWriter(os.Stdout, infoFile)
	infoLogger := slog.NewJSONHandler(multiInfoLog, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// Error logger only logs errors, but writes to separate error file
	multiErrLog := io.MultiWriter(os.Stdout, errorFile)
	errLogger := slog.NewJSONHandler(multiErrLog, &slog.HandlerOptions{
		Level: slog.LevelError,
	})

	InfoLog := slog.New(&CustomHandler{infoLogger})
	ErrLog := slog.New(&CustomHandler{errLogger})

	return &AppLogger{
		InfoLog, ErrLog,
	}
}

// Close closes all log files. Returns the first error encountered, if any.
// It is safe to call Close() multiple times.
func Close() error {
	var firstErr error

	if infoFile != nil {
		if err := infoFile.Close(); err != nil {
			firstErr = err
		}
		infoFile = nil
	}

	if errorFile != nil {
		if err := errorFile.Close(); err != nil {
			if firstErr == nil {
				firstErr = err
			}
		}
		errorFile = nil
	}

	return firstErr
}

func FromContext(ctx context.Context) *AppLogger {

	if logger, ok := ctx.Value(constants.LOGGER_KEY).(*AppLogger); ok {
		return logger
	}

	return &AppLogger{
		InfoLogger: slog.Default(),
		ErrLogger:  slog.Default(),
	}
}
