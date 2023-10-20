-- name: ListGameDiffSets :many
SELECT * FROM game_difficulty_settings
ORDER BY diff_lv
LIMIT $1
OFFSET $2;
