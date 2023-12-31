package db

import (
	"Red_Pocket_Game_Server/util"
	"context"
	"database/sql"
	"log"
	"strconv"
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

// Insert a random gameDiff, but specific CreatorID in the DB
func insertRandomGameDiffWithCreatorID(testQueries *Queries) (GameDifficultySetting, error) {
	return testQueries.CreateDiffLv(context.Background(), CreateDiffLvParams{
		DiffLv:       util.RandomInt32(10, 1000),
		AwardDensity: util.RandomInt32(10, 1000),
		EnemyDensity: util.RandomInt32(10, 1000),
	})
}

// Insert a random prize in the DB
func insertRandomPrizeWithCreatorID(testQueries *Queries) (Prize, error) {
	return testQueries.CreatePrize(context.Background(), CreatePrizeParams{
		PrizeName: util.RandomString(20),
		PicPath:   util.RandomString(20),
		Weight:    util.RandomInt32(10, 1000),
	})
}

// Check or create a user with user_id == -1 to ensure foreign key constraints
func checkOrCreateUserIdMinusOne(testQueries *Queries) error {
	// var exist bool = false
	// Check any user with user_id == -1
	_, err := testQueries.GetUserById(context.Background(), -1)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no user with user_id == -1, create one
			user, err := testQueries.CreateUserByUserId(context.Background(), CreateUserByUserIdParams{
				UserID:   -1,
				UserName: util.RandomString(20),
				Password: util.RandomString(20),
				RoleType: 1,
			})
			if err != nil {
				return err
			}
			log.Println("No user with user_id ==" + strconv.FormatInt(user.UserID, 10) + ", create one")
		} else {
			// exist = true
			return err
		}
	}
	return nil
}
