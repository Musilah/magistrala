// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/absmach/magistrala"
	mgclients "github.com/absmach/magistrala/pkg/clients"
	"github.com/absmach/magistrala/users"
)

var _ users.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    users.Service
}

// LoggingMiddleware adds logging facilities to the clients service.
func LoggingMiddleware(svc users.Service, logger *slog.Logger) users.Service {
	return &loggingMiddleware{logger, svc}
}

// RegisterClient logs the register_client request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) RegisterClient(ctx context.Context, token string, client mgclients.Client) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("name", c.Name), slog.String("id", c.ID)),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Register client failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Register client completed successfully", args...)
	}(time.Now())
	return lm.svc.RegisterClient(ctx, token, client)
}

// IssueToken logs the issue_token request. It logs the client identity type and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) IssueToken(ctx context.Context, identity, secret, domainID string) (t *magistrala.Token, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", domainID),
		}
		if t.AccessType != "" {
			args = append(args, "access_type", t.AccessType)
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Issue token failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Issue token completed successfully", args...)
	}(time.Now())
	return lm.svc.IssueToken(ctx, identity, secret, domainID)
}

// RefreshToken logs the refresh_token request. It logs the refreshtoken, token type and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) RefreshToken(ctx context.Context, refreshToken, domainID string) (t *magistrala.Token, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", domainID),
		}
		if t.AccessType != "" {
			args = append(args, "access_type", t.AccessType)
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Refresh token failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Refresh token completed successfully", args...)
	}(time.Now())
	return lm.svc.RefreshToken(ctx, refreshToken, domainID)
}

// ViewClient logs the view_client request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ViewClient(ctx context.Context, token, id string) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("name", c.Name), slog.String("id", id)),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("View client failed to complete successfully", args...)
			return
		}
		lm.logger.Info("View client completed successfully", args...)
	}(time.Now())
	return lm.svc.ViewClient(ctx, token, id)
}

// ViewProfile logs the view_profile request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ViewProfile(ctx context.Context, token string) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("name", c.Name), slog.String("id", c.ID)),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("View profile failed to complete successfully", args...)
			return
		}
		lm.logger.Info("View profile completed successfully", args...)
	}(time.Now())
	return lm.svc.ViewProfile(ctx, token)
}

// ListClients logs the list_clients request. It logs the page metadata and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ListClients(ctx context.Context, token string, pm mgclients.Page) (cp mgclients.ClientsPage, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group(
				"page",
				slog.Any("limit", pm.Limit),
				slog.Any("offset", pm.Offset),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("List clients failed to complete successfully", args...)
			return
		}
		lm.logger.Info("List clients completed successfully", args...)
	}(time.Now())
	return lm.svc.ListClients(ctx, token, pm)
}

// UpdateClient logs the update_client request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateClient(ctx context.Context, token string, client mgclients.Client) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("name", c.Name), slog.String("id", c.ID)), slog.Any("metadata", c.Metadata),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update client failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update client completed successfully", args...)
	}(time.Now())
	return lm.svc.UpdateClient(ctx, token, client)
}

// UpdateClientTags logs the update_client_tags request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateClientTags(ctx context.Context, token string, client mgclients.Client) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("id", c.ID), slog.String("tags", fmt.Sprintf("%v", c.Tags))),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update client tags failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update client tags completed successfully", args...)
	}(time.Now())
	return lm.svc.UpdateClientTags(ctx, token, client)
}

// UpdateClientIdentity logs the update_identity request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateClientIdentity(ctx context.Context, token, id, identity string) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("id", c.ID), slog.String("identity", identity)),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update client identity failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update client identity completed successfully", args...)
	}(time.Now())
	return lm.svc.UpdateClientIdentity(ctx, token, id, identity)
}

// UpdateClientSecret logs the update_client_secret request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateClientSecret(ctx context.Context, token, oldSecret, newSecret string) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("id", c.ID), slog.String("name", c.Name)),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update client secret failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update client secret completed successfully", args...)
	}(time.Now())
	return lm.svc.UpdateClientSecret(ctx, token, oldSecret, newSecret)
}

// GenerateResetToken logs the generate_reset_token request. It logs the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) GenerateResetToken(ctx context.Context, email, host string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Generate reset token failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Generate reset token completed successfully", args...)
	}(time.Now())
	return lm.svc.GenerateResetToken(ctx, email, host)
}

// ResetSecret logs the reset_secret request. It logs the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ResetSecret(ctx context.Context, token, secret string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Reset secret failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Reset secret completed successfully", args...)
	}(time.Now())
	return lm.svc.ResetSecret(ctx, token, secret)
}

// SendPasswordReset logs the send_password_reset request. It logs the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) SendPasswordReset(ctx context.Context, host, email, user, token string) (err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Send password reset failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Send password reset completed successfully", args...)
	}(time.Now())
	return lm.svc.SendPasswordReset(ctx, host, email, user, token)
}

// UpdateClientRole logs the update_client_role request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) UpdateClientRole(ctx context.Context, token string, client mgclients.Client) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("user", slog.String("id", c.ID), slog.String("role", client.Role.String())),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update client role failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Update client role completed successfully", args...)
	}(time.Now())
	return lm.svc.UpdateClientRole(ctx, token, client)
}

// EnableClient logs the enable_client request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) EnableClient(ctx context.Context, token, id string) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("id", c.ID),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Enable client failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Enable client completed successfully", args...)
	}(time.Now())
	return lm.svc.EnableClient(ctx, token, id)
}

// DisableClient logs the disable_client request. It logs the client id and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) DisableClient(ctx context.Context, token, id string) (c mgclients.Client, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("id", c.ID),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Disable client failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Disable client completed successfully", args...)
	}(time.Now())
	return lm.svc.DisableClient(ctx, token, id)
}

// ListMembers logs the list_members request. It logs the group id, and the time it took to complete the request.
// If the request fails, it logs the error.
func (lm *loggingMiddleware) ListMembers(ctx context.Context, token, objectKind, objectID string, cp mgclients.Page) (mp mgclients.MembersPage, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
			slog.String("object_kind", objectKind),
			slog.String("object_id", objectID),
			slog.Group(
				"page",
				slog.Any("limit", cp.Limit),
				slog.Any("offset", cp.Offset),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("List members failed to complete successfully", args...)
			return
		}
		lm.logger.Info("List members completed successfully", args...)
	}(time.Now())
	return lm.svc.ListMembers(ctx, token, objectKind, objectID, cp)
}

// Identify logs the identify request. It logs the time it took to complete the request.
func (lm *loggingMiddleware) Identify(ctx context.Context, token string) (id string, err error) {
	defer func(begin time.Time) {
		args := []interface{}{
			slog.String("duration", time.Since(begin).String()),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Identify failed to complete successfully", args...)
			return
		}
		lm.logger.Info("Identify completed successfully", args...)
	}(time.Now())
	return lm.svc.Identify(ctx, token)
}
