package db

import (
	"context"
	"database/sql"
)

type Agent struct {
	Id          int64
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

type AgentUserMapping struct {
	userId  int
	agentId int
	role    sql.NullString
}

func GetAgentUserMapping(userId int, agentId int, ctx context.Context) (AgentUserMapping, error) {
	var mapping AgentUserMapping
	row := DbPool.QueryRowContext(ctx,
		`SELECT role FROM agent_user WHERE user_id=? AND agent_id=?`,
		userId, agentId,
	)

	err := row.Scan(&mapping.role)
	if err != nil {
		return AgentUserMapping{}, err
	}
	mapping.agentId = agentId
	mapping.userId = userId

	return mapping, nil
}

func GetAgentById(id int, userId int, ctx context.Context) (Agent, error) {
	row := DbPool.QueryRowContext(ctx,
		`SELECT id, name, industry, personality FROM agent ag
		LEFT JOIN agent_user au ON au.agent_id = ag.id
		WHERE ag.id = ? and au.user_id = ?`,
		id, userId)

	var agent Agent
	err := row.Scan(&agent.Id, &agent.Name, &agent.Industry, &agent.Personality)
	if err != nil {
		return Agent{}, err
	}

	return agent, err
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

type UpdateAgentInput struct {
	Id          int
	Name        string
	Industry    string
	Personality string
}

func UpdateAgent(updateInput UpdateAgentInput, ctx context.Context) error {
	_, err := DbPool.ExecContext(ctx, `
		UPDATE agent
        SET name = ?, industry = ?, personality = ?
		WHERE id = ?
		`, updateInput.Name, updateInput.Industry, updateInput.Personality, updateInput.Id)
	if err != nil {
		return nil
	}

	return nil
}
