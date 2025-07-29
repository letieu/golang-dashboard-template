package db

import (
	"context"
	"fmt"
	"time"
)

func GetUserIdByEmail(email string, ctx context.Context) (uint, error) {
	var userId uint
	row := DbPool.QueryRowContext(ctx, "SELECT rowid FROM user WHERE email=?", email)
	err := row.Scan(&userId)
	return userId, err
}

func UpsertUser(email, username string, ctx context.Context) error {
	_, err := DbPool.ExecContext(ctx, `
        INSERT INTO user (email, username)
        VALUES (?, ?)
        ON CONFLICT(email) DO UPDATE SET
            username = excluded.username
    `, email, username)

	return err
}

func SetSession(email string, sessionToken string, expireAt time.Time) error {
	_, err := DbPool.Exec(`
		INSERT INTO session (token, user_email, expire_at) VALUES (?, ?, ?)
	`, sessionToken, email, expireAt)
	return err
}

func GetUserTenantIds(userId uint, ctx context.Context) ([]uint, error) {
	rows, err := DbPool.QueryContext(ctx, "SELECT tenant_id FROM tenant_user tu WHERE tu.user_id = ?", userId)
	if err != nil {
		return nil, err
	}

	var tenantIDs []uint
	for rows.Next() {
		var tenantID uint
		if err := rows.Scan(&tenantID); err != nil {
			return nil, err
		}
		tenantIDs = append(tenantIDs, tenantID)
	}

	return tenantIDs, nil
}

func CreateDefaultTenantForUser(userId uint, ctx context.Context) (uint, error) {
	res, err := DbPool.ExecContext(ctx, "INSERT INTO tenant (name) VALUES (?)", "Default Tenant")
	if err != nil {
		return 0, fmt.Errorf("failed to create tenant: %w", err)
	}

	tenantID64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get new tenant ID: %w", err)
	}
	tenantId := uint(tenantID64)

	_, err = DbPool.ExecContext(ctx, "INSERT INTO tenant_user (tenant_id, user_id) VALUES (?, ?)", tenantId, userId)
	if err != nil {
		return 0, fmt.Errorf("failed to link user to tenant: %w", err)
	}

	return tenantId, nil
}

func CreateDefaultAgentForTenant(tenantId uint, ctx context.Context) error {
	_, err := DbPool.ExecContext(ctx, "INSERT INTO agent (name, tenant_id) VALUES (?, ?)", "Default Agent", tenantId)
	if err != nil {
		return fmt.Errorf("failed to create agent: %w", err)
	}
	return nil
}

func GetAgentIdsByTenantId(tenantId uint, ctx context.Context) ([]uint, error) {
	rows, err := DbPool.QueryContext(ctx, "SELECT rowid FROM agent ag WHERE ag.tenant_id = ?", tenantId)
	if err != nil {
		return nil, err
	}

	var agentIDs []uint
	for rows.Next() {
		var agentId uint
		if err := rows.Scan(&agentId); err != nil {
			return nil, err
		}
		agentIDs = append(agentIDs, agentId)
	}

	return agentIDs, nil
}
