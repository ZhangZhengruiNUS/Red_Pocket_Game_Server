package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItem(t *testing.T) {

	// Declare variables
	var item1 Item
	var item2 Item
	var item3 Item
	var item4 Item
	var item5 Item
	var err error

	// Delete test data when testing ends
	defer func() {
		err = testQueries.DeleteItem(context.Background(), item1.ItemID)
		require.NoError(t, err)
		err = testQueries.DeleteItem(context.Background(), item2.ItemID)
		require.NoError(t, err)
		err = testQueries.DeleteItem(context.Background(), item3.ItemID)
		require.NoError(t, err)
		err = testQueries.DeleteItem(context.Background(), item4.ItemID)
		require.NoError(t, err)
		err = testQueries.DeleteItem(context.Background(), item5.ItemID)
		require.NoError(t, err)
	}()

	// Get a userID for creating new items
	userID, err := getFirstUserIDFromDB(testQueries)
	require.NoError(t, err)

	// Create test data in the DB
	item1, err = insertRandomItemWithCreatorID(testQueries, userID)
	require.NoError(t, err)
	item2, err = insertRandomItemWithCreatorID(testQueries, userID)
	require.NoError(t, err)
	item3, err = insertRandomItemWithCreatorID(testQueries, userID)
	require.NoError(t, err)
	item4, err = insertRandomItemWithCreatorID(testQueries, userID)
	require.NoError(t, err)
	item5, err = insertRandomItemWithCreatorID(testQueries, userID)
	require.NoError(t, err)

	// Call ListItems for OK
	productList, err := testQueries.ListItems(context.Background(), ListItemsParams{
		Page:     2,
		Pagesize: 2,
	})
	// Check ListItems for OK
	require.NoError(t, err)
	require.NotEmpty(t, productList)
	require.Equal(t, productList, []Item{
		item3,
		item4,
	},
	)

	// Call GetItem for OK
	item1_getTest, err := testQueries.GetItem(context.Background(), item1.ItemID)
	// Check GetItem for OK
	require.NoError(t, err)
	require.NotEmpty(t, item1_getTest)
	require.Equal(t, item1_getTest, item1)

}
