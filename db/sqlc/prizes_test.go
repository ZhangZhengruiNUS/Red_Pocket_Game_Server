package db

import (
	"Red_Pocket_Game_Server/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrize(t *testing.T) {

	// Declare variables
	var prize1 Prize
	var prize2 Prize
	var prize3 Prize
	var prize4 Prize
	var prize5 Prize
	var err error

	// Delete test data when testing ends
	defer func() {
		err = testQueries.DeletePrize(context.Background(), prize1.PrizeName)
		require.NoError(t, err)
		err = testQueries.DeletePrize(context.Background(), prize2.PrizeName)
		require.NoError(t, err)
		err = testQueries.DeletePrize(context.Background(), prize3.PrizeName)
		require.NoError(t, err)
		err = testQueries.DeletePrize(context.Background(), prize4.PrizeName)
		require.NoError(t, err)
		err = testQueries.DeletePrize(context.Background(), prize5.PrizeName)
		require.NoError(t, err)
	}()

	// // Check or create a user with user_id == -1 to ensure foreign key constraints
	// err = checkOrCreateUserIdMinusOne(testQueries)

	// Create test data in the DB
	prize1, err = insertRandomPrizeWithCreatorID(testQueries)
	require.NoError(t, err)
	prize2, err = insertRandomPrizeWithCreatorID(testQueries)
	require.NoError(t, err)
	prize3, err = insertRandomPrizeWithCreatorID(testQueries)
	require.NoError(t, err)
	prize4, err = insertRandomPrizeWithCreatorID(testQueries)
	require.NoError(t, err)
	prize5, err = insertRandomPrizeWithCreatorID(testQueries)
	require.NoError(t, err)

	// Call ListRolltable for OK
	prizeList, err := testQueries.ListRolltable(context.Background(), ListRolltableParams{
		Page:     2,
		Pagesize: 2,
	})

	// Get a sclice of Prize as ListRolltable actually returns
	var prize3Slice = ListRolltableRow{
		PrizeName: prize3.PrizeName,
		PicPath:   prize3.PicPath,
		Weight:    prize3.Weight,
	}
	var prize4Slice = ListRolltableRow{
		PrizeName: prize4.PrizeName,
		PicPath:   prize4.PicPath,
		Weight:    prize4.Weight,
	}

	// Check ListRolltable for OK
	require.NoError(t, err)
	require.NotEmpty(t, prizeList)
	require.Equal(t, prizeList, []ListRolltableRow{
		prize3Slice,
		prize4Slice,
	},
	)

	// Call GetPrizeByPrizeName for OK
	prize1_getTest, err := testQueries.GetPrizeByPrizeName(context.Background(), prize1.PrizeName)
	// Check GetPrizeByPrizeName for OK
	require.NoError(t, err)
	require.NotEmpty(t, prize1_getTest)
	require.Equal(t, prize1_getTest, prize1)

	// Updata prize5
	prize5.PicPath = util.RandomString(20)
	prize5.Weight = util.RandomInt32(10, 1000)

	// Updata test data in the DB
	prize5Updated, err := testQueries.UpdatePrize(context.Background(), UpdatePrizeParams{
		Picpath:   prize5.PicPath,
		Weight:    prize5.Weight,
		PrizeName: prize5.PrizeName,
	})

	// Check UpdateDiffLv for OK
	require.NoError(t, err)
	require.NotEmpty(t, prize5Updated)
	require.Equal(t, prize5, prize5Updated)
}
