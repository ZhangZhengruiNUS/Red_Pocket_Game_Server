-- name: ListGameDiffSets :many
SELECT * FROM game_difficulty_settings
ORDER BY diff_lv
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);
