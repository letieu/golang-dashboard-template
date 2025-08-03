package db

import (
	"context"
)

type Agent struct {
	Id          uint
	Name        string
	Industry    string
	Personality string
}

func GetAgentsByUserId(userId int64, ctx context.Context) ([]Agent, error) {
	rows, err := DbPool.QueryContext(ctx, `
		SELECT id, name, industry, personality FROM agent ag
		LEFT JOIN agent_user au ON ag.id = au.agent_id
		WHERE au.user_id = ?
		`, userId)

	if err != nil {
		return nil, err
	}

	var agents []Agent
	for rows.Next() {
		var agent Agent
		err := rows.Scan(&agent.Id, &agent.Name, &agent.Industry, &agent.Personality)
		if err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}

	return agents, nil
}

type CreateAgentInput struct {
	Name        string
	Industry    string
	Personality string
	UserId      int64
}

func CreateAgent(input CreateAgentInput, ctx context.Context) (int64, error) {
	// INSERT INTO user (email, username)
	result, err := DbPool.ExecContext(ctx, `
		INSERT INTO agent (name, industry, personality)
        VALUES (?, ?, ?)
		`, input.Name, input.Industry, input.Personality)
	if err != nil {
		return 0, err
	}

	agentId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	result, err = DbPool.ExecContext(ctx, `
		INSERT INTO agent_user (agent_id, user_id)
        VALUES (?, ?)
		`, agentId, input.UserId)
	if err != nil {
		return 0, err
	}

	return agentId, nil
}
