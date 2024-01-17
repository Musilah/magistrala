// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

//go:build !test

package api

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/absmach/magistrala/certs"
)

var _ certs.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    certs.Service
}

// LoggingMiddleware adds logging facilities to the bootstrap service.
func LoggingMiddleware(svc certs.Service, logger *slog.Logger) certs.Service {
	return &loggingMiddleware{logger, svc}
}

// IssueCert logs the issue_cert request. It logs the token, thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) IssueCert(ctx context.Context, token, thingID, ttl string) (c certs.Cert, err error) {
	defer func(begin time.Time) {
		message:= "Method issue_cert completed"
		if err != nil {
			lm.logger.Warn(
				fmt.Sprintf("%s with error.", message),
				slog.String("error", err.Error()),
				slog.String("duration", time.Since(begin).String()),
			)
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message),
			slog.String("thing_id", thingID),
			slog.String("token", token),
			slog.String("ttl", ttl),
			slog.String("duration", time.Since(begin).String()),
		)
	}(time.Now())

	return lm.svc.IssueCert(ctx, token, thingID, ttl)
}

// ListCerts logs the list_certs request. It logs the token, thing ID and the time it took to complete the request.
func (lm *loggingMiddleware) ListCerts(ctx context.Context, token, thingID string, offset, limit uint64) (cp certs.Page, err error) {
	defer func(begin time.Time) {
		message:= "Method list_certs completed"
		if err != nil {
			lm.logger.Warn(
				fmt.Sprintf("%s with error.", message),
				slog.String("error", err.Error()),
				slog.String("duration", time.Since(begin).String()),
			)
			return
		}
		lm.logger.Info(
			fmt.Sprintf("%s without errors.", message),
			slog.String("thing_id", thingID),
			slog.String("token", token),
			slog.Uint64("offset", offset),
			slog.Uint64("limit", limit),
			slog.String("duration", time.Since(begin).String()),
		)
	}(time.Now())

	return lm.svc.ListCerts(ctx, token, thingID, offset, limit)
}

// ListSerials logs the list_serials request. It logs the token, thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ListSerials(ctx context.Context, token, thingID string, offset, limit uint64) (cp certs.Page, err error) {
	defer func(begin time.Time) {
		message:= "Method list_serials completed"
		if err != nil {
			lm.logger.Warn(
				fmt.Sprintf("%s with error.", message),
				slog.String("error", err.Error()),
				slog.String("duration", time.Since(begin).String()),
			)
			return
		}
		lm.logger.Info(
			fmt.Sprintf("%s without errors.", message),
			slog.String("token", token),
			slog.String("thing_id", thingID),
			slog.Uint64("offset", offset),
			slog.Uint64("limit", limit),
			slog.String("duration", time.Since(begin).String()),
		)
	}(time.Now())

	return lm.svc.ListSerials(ctx, token, thingID, offset, limit)
}

// ViewCert logs the view_cert request. It logs the token, serial ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ViewCert(ctx context.Context, token, serialID string) (c certs.Cert, err error) {
	defer func(begin time.Time) {
		message:= "Method view_cert completed"
		if err != nil {
			lm.logger.Warn(
				fmt.Sprintf("%s with error.", message),
				slog.String("error", err.Error()),
				slog.String("duration", time.Since(begin).String()),
			)
			return
		}
		lm.logger.Info(
			fmt.Sprintf("%s without errors.", message),
			slog.String("token", token),
			slog.String("serial_id", serialID),
			slog.String("duration", time.Since(begin).String()),
		)
	}(time.Now())

	return lm.svc.ViewCert(ctx, token, serialID)
}

// RevokeCert logs the revoke_cert request. It logs the token, thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) RevokeCert(ctx context.Context, token, thingID string) (c certs.Revoke, err error) {
	defer func(begin time.Time) {
		message:= "Method revoke_cert completed"
		if err != nil {
			lm.logger.Warn(
				fmt.Sprintf("%s with error.", message),
				slog.String("error", err.Error()),
				slog.String("duration", time.Since(begin).String()),
			)
			return
		}
		lm.logger.Info(
			fmt.Sprintf("%s without errors.", message),
			slog.String("token", token),
			slog.String("thing_id", thingID),
			slog.String("duration", time.Since(begin).String()),
		)
	}(time.Now())

	return lm.svc.RevokeCert(ctx, token, thingID)
}
