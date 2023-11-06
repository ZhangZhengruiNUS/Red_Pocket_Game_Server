package server

import db "Red_Pocket_Game_Server/db/sqlc"

// Extra difficulty strategy interface
type ExtraDifficultyStrategy interface {
	AdjustGameSetting(gameDifficultySettings []db.GameDifficultySetting)
}

type NormalDiffStrategy struct{}

// Hard game difficulty game setting
func (h NormalDiffStrategy) AdjustGameSetting(gameDifficultySettings []db.GameDifficultySetting) {}

type EasyDiffStrategy struct{}

// Easy game difficulty setting
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

type HardDiffStrategy struct{}

// Hard game difficulty game setting
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

// Adjust Game Setting
func adjustGameSetting(strategy ExtraDifficultyStrategy, gameDifficultySettings []db.GameDifficultySetting) {
	strategy.AdjustGameSetting(gameDifficultySettings)
}
