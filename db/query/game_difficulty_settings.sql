-- name: ListGameDiffSets :many
SELECT * FROM game_difficulty_settings
ORDER BY diff_lv
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);

-- name: ListGameDiffSetsByDiffLv :one
SELECT * FROM game_difficulty_settings
WHERE diff_lv = $1;

-- name: CreateDiffLv :one
INSERT INTO game_difficulty_settings (
    diff_lv,
    award_density,
    enemy_density
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteDiffLv :exec
DELETE FROM game_difficulty_settings
WHERE diff_lv = $1
AND award_density = $2
AND enemy_density = $3;

-- name: UpdateDiffLv :one
UPDATE game_difficulty_settings
SET award_density = sqlc.arg(awardDensity), enemy_density = sqlc.arg(enemyDensity)
WHERE diff_lv = $1
RETURNING *;

-- name: DeleteAllGameDiffSets :exec
DELETE FROM game_difficulty_settings;