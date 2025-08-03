package db

import (
	"context"
	"time"
)

func GetUserIdByEmail(email string, ctx context.Context) (int64, error) {
	var userId int64
	row := DbPool.QueryRowContext(ctx, "SELECT id FROM user WHERE email=?", email)
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

func SetSession(userId int64, sessionToken string, expireAt time.Time) error {
	_, err := DbPool.Exec(`
		INSERT INTO session (token, user_id, expire_at) VALUES (?, ?, ?)
	`, sessionToken, userId, expireAt)
	return err
}
