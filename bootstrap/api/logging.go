// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

//go:build !test

package api

import (
	"context"
	"log/slog"
	"time"

	"github.com/absmach/magistrala/bootstrap"
)

var _ bootstrap.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    bootstrap.Service
}

// LoggingMiddleware adds logging facilities to the bootstrap service.
func LoggingMiddleware(svc bootstrap.Service, logger *slog.Logger) bootstrap.Service {
	return &loggingMiddleware{logger, svc}
}

// Add logs the add request. It logs the thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) Add(ctx context.Context, token string, cfg bootstrap.Config) (saved bootstrap.Config, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("thing_id", saved.ThingID),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Add failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Add completed successfully", args...)
	}(time.Now())

	return lm.svc.Add(ctx, token, cfg)
}

// View logs the view request. It logs the thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) View(ctx context.Context, token, id string) (saved bootstrap.Config, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("thing_id", id),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("View failed to complete successfully", args...)
			return
		}
		lm.logger.Info("View completed successfully", args...)
	}(time.Now())

	return lm.svc.View(ctx, token, id)
}

// Update logs the update request. It logs bootstrap thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) Update(ctx context.Context, token string, cfg bootstrap.Config) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("thing_id", cfg.ThingID),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update completed successfully", args...)
	}(time.Now())

	return lm.svc.Update(ctx, token, cfg)
}

// UpdateCert logs the update_cert request. It logs thing ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateCert(ctx context.Context, token, thingID, clientCert, clientKey, caCert string) (cfg bootstrap.Config, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("thing_id", cfg.ThingID),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update cert failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update cert completed successfully", args...)
	}(time.Now())

	return lm.svc.UpdateCert(ctx, token, thingID, clientCert, clientKey, caCert)
}

// UpdateConnections logs the update_connections request. It logs bootstrap ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateConnections(ctx context.Context, token, id string, connections []string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("id", id),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update connections failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update connections completed successfully", args...)
	}(time.Now())

	return lm.svc.UpdateConnections(ctx, token, id, connections)
}

// List logs the list request. It logs offset, limit and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) List(ctx context.Context, token string, filter bootstrap.Filter, offset, limit uint64) (res bootstrap.ConfigsPage, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group(
				"filter",
				slog.Uint64("offset", offset),
				slog.Uint64("limit", limit),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("List failed to complete successfully", args...)
			return
		}
		lm.logger.Info("List completed successfully", args...)
	}(time.Now())

	return lm.svc.List(ctx, token, filter, offset, limit)
}

// Remove logs the remove request. It logs bootstrap ID and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) Remove(ctx context.Context, token, id string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("id", id),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Remove failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Remove completed successfully", args...)
	}(time.Now())

	return lm.svc.Remove(ctx, token, id)
}

func (lm *loggingMiddleware) Bootstrap(ctx context.Context, externalKey, externalID string, secure bool) (cfg bootstrap.Config, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group(
				"config",
				slog.String("external_id", externalID),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Bootstrap failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Bootstrap completed successfully", args...)
	}(time.Now())

	return lm.svc.Bootstrap(ctx, externalKey, externalID, secure)
}

func (lm *loggingMiddleware) ChangeState(ctx context.Context, token, id string, state bootstrap.State) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("id", id),
			slog.Any("state", state),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Change state failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Change state completed successfully", args...)
	}(time.Now())

	return lm.svc.ChangeState(ctx, token, id, state)
}

func (lm *loggingMiddleware) UpdateChannelHandler(ctx context.Context, channel bootstrap.Channel) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group(
				"channel",
				slog.String("channel_id", channel.ID),
				slog.String("channel_name", channel.Name),
				slog.Any("channel_metadata", channel.Metadata),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update channel handler failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update channel handler completed successfully", args...)
	}(time.Now())

	return lm.svc.UpdateChannelHandler(ctx, channel)
}

func (lm *loggingMiddleware) RemoveConfigHandler(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("config_id", id),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Remove config handler failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Remove config handler completed successfully", args...)
	}(time.Now())

	return lm.svc.RemoveConfigHandler(ctx, id)
}

func (lm *loggingMiddleware) RemoveChannelHandler(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("id", id),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Remove channel handler failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Remove channel handler completed successfully", args...)
	}(time.Now())

	return lm.svc.RemoveChannelHandler(ctx, id)
}

func (lm *loggingMiddleware) DisconnectThingHandler(ctx context.Context, channelID, thingID string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("channel_id", channelID),
			slog.String("thing_id", thingID),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Disconnect thing handler failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Disconnect thing handler completed successfully", args...)
	}(time.Now())

	return lm.svc.DisconnectThingHandler(ctx, channelID, thingID)
}
