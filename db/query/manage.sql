-- name: CreatePrize :one
INSERT INTO prizes (
    prize_name,
    pic_path,
    weight
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPrizeByPrizeName :one
SELECT * FROM prizes
WHERE prize_name = $1 LIMIT 1;

-- name: DeletePrize :exec
DELETE FROM prizes
WHERE prize_name = $1;

-- name: UpdatePrize :one
UPDATE prizes
SET pic_path = sqlc.arg(picPath), weight = sqlc.arg(weight)
WHERE prize_name = sqlc.arg(prize_name)
RETURNING *;

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