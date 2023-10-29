// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: game_difficulty_settings.sql

package db

import (
	"context"
)

const listGameDiffSets = `-- name: ListGameDiffSets :many
SELECT diff_lv, award_density, enemy_density, reviser_id, revise_time, creator_id, create_time FROM game_difficulty_settings
ORDER BY diff_lv
LIMIT $1
OFFSET $2
`

type ListGameDiffSetsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListGameDiffSets(ctx context.Context, arg ListGameDiffSetsParams) ([]GameDifficultySetting, error) {
	rows, err := q.query(ctx, q.listGameDiffSetsStmt, listGameDiffSets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GameDifficultySetting{}
	for rows.Next() {
		var i GameDifficultySetting
		if err := rows.Scan(
			&i.DiffLv,
			&i.AwardDensity,
			&i.EnemyDensity,
			&i.ReviserID,
			&i.ReviseTime,
			&i.CreatorID,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}