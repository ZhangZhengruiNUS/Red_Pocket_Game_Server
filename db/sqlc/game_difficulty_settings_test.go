package db

// import (
// 	"Red_Pocket_Game_Server/util"
// 	"context"
// 	"testing"

// 	"sort"

// 	"github.com/stretchr/testify/require"
// )

// func TestGameDiffSettings(t *testing.T) {

// 	// Declare variables
// 	var gameDiff1 GameDifficultySetting
// 	var gameDiff2 GameDifficultySetting
// 	var gameDiff3 GameDifficultySetting
// 	var gameDiff4 GameDifficultySetting
// 	var gameDiff5 GameDifficultySetting
// 	var err error

// 	// Delete test data when testing ends
// 	defer func() {
// 		err = testQueries.DeleteDiffLv(context.Background(), DeleteDiffLvParams{
// 			DiffLv:       gameDiff1.DiffLv,
// 			AwardDensity: gameDiff1.AwardDensity,
// 			EnemyDensity: gameDiff1.EnemyDensity,
// 		})
// 		require.NoError(t, err)
// 		err = testQueries.DeleteDiffLv(context.Background(), DeleteDiffLvParams{
// 			DiffLv:       gameDiff2.DiffLv,
// 			AwardDensity: gameDiff2.AwardDensity,
// 			EnemyDensity: gameDiff2.EnemyDensity,
// 		})
// 		require.NoError(t, err)
// 		err = testQueries.DeleteDiffLv(context.Background(), DeleteDiffLvParams{
// 			DiffLv:       gameDiff3.DiffLv,
// 			AwardDensity: gameDiff3.AwardDensity,
// 			EnemyDensity: gameDiff3.EnemyDensity,
// 		})
// 		require.NoError(t, err)
// 		err = testQueries.DeleteDiffLv(context.Background(), DeleteDiffLvParams{
// 			DiffLv:       gameDiff4.DiffLv,
// 			AwardDensity: gameDiff4.AwardDensity,
// 			EnemyDensity: gameDiff4.EnemyDensity,
// 		})
// 		require.NoError(t, err)
// 		err = testQueries.DeleteDiffLv(context.Background(), DeleteDiffLvParams{
// 			DiffLv:       gameDiff5.DiffLv,
// 			AwardDensity: gameDiff5.AwardDensity,
// 			EnemyDensity: gameDiff5.EnemyDensity,
// 		})
// 		require.NoError(t, err)
// 	}()

// 	// Create test data in the DB
// 	gameDiff1, err = insertRandomGameDiffWithCreatorID(testQueries)
// 	require.NoError(t, err)
// 	gameDiff2, err = insertRandomGameDiffWithCreatorID(testQueries)
// 	require.NoError(t, err)
// 	gameDiff3, err = insertRandomGameDiffWithCreatorID(testQueries)
// 	require.NoError(t, err)
// 	gameDiff4, err = insertRandomGameDiffWithCreatorID(testQueries)
// 	require.NoError(t, err)
// 	gameDiff5, err = insertRandomGameDiffWithCreatorID(testQueries)
// 	require.NoError(t, err)

// 	gameDiff := []GameDifficultySetting{
// 		gameDiff1,
// 		gameDiff2,
// 		gameDiff3,
// 		gameDiff4,
// 		gameDiff5,
// 	}

// 	// Order the gameDiffs by DiffLv as ListGameDiffSets dose
// 	sort.Slice(gameDiff, func(i, j int) bool {
// 		return gameDiff[i].DiffLv < gameDiff[j].DiffLv
// 	})

// 	// Call ListGameDiffSets for OK
// 	gameDiffList, err := testQueries.ListGameDiffSets(context.Background(), ListGameDiffSetsParams{
// 		Page:     2,
// 		Pagesize: 2,
// 	})

// 	// Check ListGameDiffSets for OK
// 	require.NoError(t, err)
// 	require.NotEmpty(t, gameDiffList)
// 	require.Equal(t, gameDiffList, []GameDifficultySetting{
// 		gameDiff[2],
// 		gameDiff[3],
// 	},
// 	)

// 	// Updata gameDiff5
// 	gameDiff5.AwardDensity = util.RandomInt32(10, 1000)
// 	gameDiff5.EnemyDensity = util.RandomInt32(10, 1000)

// 	// Updata test data in the DB
// 	gameDiff5Updated, err := testQueries.UpdateDiffLv(context.Background(), UpdateDiffLvParams{
// 		DiffLv:       gameDiff5.DiffLv,
// 		Awarddensity: gameDiff5.AwardDensity,
// 		Enemydensity: gameDiff5.EnemyDensity,
// 	})

// 	// Check UpdateDiffLv for OK
// 	require.NoError(t, err)
// 	require.NotEmpty(t, gameDiff5Updated)
// 	require.Equal(t, gameDiff5, gameDiff5Updated)
// }
