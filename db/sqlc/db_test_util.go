package db

import (
	"Red_Pocket_Game_Server/util"
	"context"
)

// Insert a random item, but specific CreatorID in the DB
func insertRandomItemWithCreatorID(testQueries *Queries, creatorId int64) (Item, error) {
	return testQueries.CreateItem(context.Background(), CreateItemParams{
		ItemName:  util.RandomString(20),
		Describe:  util.RandomString(20),
		PicPath:   util.RandomString(20),
		Price:     util.RandomInt32(10, 1000),
		CreatorID: creatorId,
	})
}

// Get First user_id in the DB
func getFirstUserIDFromDB(testQueries *Queries) (int64, error) {
	// Get first user in the DB
	users, err := testQueries.ListUsers(context.Background(), ListUsersParams{
		Page:     1,
		Pagesize: 1,
	})
	if err != nil {
		return 0, err
	}
	// If no user in the DB, then create one
	if len(users) == 0 {
		user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
			UserName: "admin",
			Password: "admin",
			RoleType: 1,
		})
		if err != nil {
			return 0, err
		}
		return user.UserID, nil
	}
	return users[0].UserID, nil
}
