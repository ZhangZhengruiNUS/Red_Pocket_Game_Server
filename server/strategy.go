package server

import db "Red_Pocket_Game_Server/db/sqlc"

// Extra difficulty strategy interface
type ExtraDifficultyStrategy interface {
	AdjustGameSetting(gameDifficultySettings []db.GameDifficultySetting)
}

/* Normal game difficulty setting */
type NormalDiffStrategy struct{}

func (h NormalDiffStrategy) AdjustGameSetting(gameDifficultySettings []db.GameDifficultySetting) {}

/* Easy game difficulty setting  */
type EasyDiffStrategy struct{}

func (e EasyDiffStrategy) AdjustGameSetting(gameDifficultySettings []db.GameDifficultySetting) {
	for index := range gameDifficultySettings {
		awardDensity := gameDifficultySettings[index].AwardDensity
		enemyDensity := gameDifficultySettings[index].EnemyDensity

		gameDifficultySettings[index].AwardDensity = awardDensity + 1
		if enemyDensity > 1 {
			gameDifficultySettings[index].EnemyDensity = enemyDensity - 1
		}
	}
}

/* Hard game difficulty setting */
type HardDiffStrategy struct{}

func (h HardDiffStrategy) AdjustGameSetting(gameDifficultySettings []db.GameDifficultySetting) {
	for index := range gameDifficultySettings {
		awardDensity := gameDifficultySettings[index].AwardDensity
		enemyDensity := gameDifficultySettings[index].EnemyDensity
		if awardDensity > 1 {
			gameDifficultySettings[index].AwardDensity = awardDensity - 1
		}
		gameDifficultySettings[index].EnemyDensity = enemyDensity + 1
	}
}
