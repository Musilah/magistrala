// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package logger

import (
	"context"
	"fmt"
	"io"
	"time"

	"golang.org/x/exp/slog"
)

// Logger specifies logging API.
type Logger interface {
	Debug(context.Context, string)
	Info(context.Context, string)
	Warn(context.Context, string)
	Error(context.Context, string)
}

type logger struct {
	sLogger *slog.Logger
	level   Level
}

// New returns a new slog logger.
func New(w io.Writer, levelText string) (Logger, error) {
	var level Level
	err := level.UnmarshalText(levelText)
	if err != nil {
		return nil, fmt.Errorf(`{"level":"error","message":"%s: %s","ts":"%s"}`, err, levelText, time.RFC3339Nano)
	}

	logHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	sLogger := slog.New(logHandler)

	return &logger{sLogger, level}, nil
}

func (l *logger) Debug(ctx context.Context, msg string) {
	if Debug.isAllowed(l.level) {
		l.sLogger.Log(ctx, slog.LevelDebug, "msg", msg)
	}
}

func (l *logger) Info(ctx context.Context, msg string) {
	if Info.isAllowed(l.level) {
		l.sLogger.Log(ctx, slog.LevelInfo, msg)
	}
}

func (l *logger) Warn(ctx context.Context, msg string) {
	if Warn.isAllowed(l.level) {
		l.sLogger.Log(ctx, slog.LevelWarn, msg)
	}
}

func (l *logger) Error(ctx context.Context, msg string) {
	if Error.isAllowed(l.level) {
		l.sLogger.Log(ctx, slog.LevelError, msg)
	}
}
