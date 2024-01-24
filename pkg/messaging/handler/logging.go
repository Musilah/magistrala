// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

//go:build !test

package handler

import (
	"context"
	"log/slog"
	"time"

	"github.com/absmach/mproxy/pkg/session"
)

var _ session.Handler = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    session.Handler
}

// AuthConnect implements session.Handler.
func (lm *loggingMiddleware) AuthConnect(ctx context.Context) (err error) {
	return lm.logAction(ctx, "AuthConnect", nil)
}

// AuthPublish implements session.Handler.
func (lm *loggingMiddleware) AuthPublish(ctx context.Context, topic *string, payload *[]byte) (err error) {
	return lm.logAction(ctx, "AuthPublish", &[]string{*topic})
}

// AuthSubscribe implements session.Handler.
func (lm *loggingMiddleware) AuthSubscribe(ctx context.Context, topics *[]string) (err error) {
	return lm.logAction(ctx, "AuthSubscribe", topics)
}

// Connect implements session.Handler.
func (lm *loggingMiddleware) Connect(ctx context.Context) (err error) {
	return lm.logAction(ctx, "Connect", nil)
}

// Disconnect implements session.Handler.
func (lm *loggingMiddleware) Disconnect(ctx context.Context) (err error) {
	return lm.logAction(ctx, "Disconnect", nil)
}

// Publish logs the publish request. It logs the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) Publish(ctx context.Context, topic *string, payload *[]byte) (err error) {
	return lm.logAction(ctx, "Publish", &[]string{*topic})
}

// Subscribe implements session.Handler.
func (lm *loggingMiddleware) Subscribe(ctx context.Context, topics *[]string) (err error) {
	return lm.logAction(ctx, "Subscribe", topics)
}

// Unsubscribe implements session.Handler.
func (lm *loggingMiddleware) Unsubscribe(ctx context.Context, topics *[]string) (err error) {
	return lm.logAction(ctx, "Unsubscribe", topics)
}

// LoggingMiddleware adds logging facilities to the adapter.
func LoggingMiddleware(svc session.Handler, logger *slog.Logger) session.Handler {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) logAction(ctx context.Context, action string, topics *[]string) (err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
		}
		if topics != nil {
			args = append(args, slog.Any("topics", *topics))
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn(action+"() failed to complete successfully", args...)
			return
		}
		lm.logger.Info(action+"() completed successfully", args...)
	}(time.Now())

	return nil
}
