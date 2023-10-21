package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* Came-diff GET handle function */
func (server *Server) gameDiffHandler(ctx *gin.Context) {
	fmt.Println("================================gameDiffHandler: Start================================")

	// Initialize query parameters
	params := db.ListGameDiffSetsParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
	}

	// Get data from database
	game_difficulty_settings, err := server.store.ListGameDiffSets(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	fmt.Println("game_difficulty_settings count:", len(game_difficulty_settings))
	ctx.JSON(http.StatusOK, game_difficulty_settings)

	fmt.Println("================================gameDiffHandler: End================================")
}

/* Game-equip GET received data */
type gameEquipQueryRequest struct {
	UserID int64 `json:"userId"`
}

/* Game-equip GET handle function */
func (server *Server) gameEquipQueryHandler(ctx *gin.Context) {
	fmt.Println("================================gameEquipQueryHandler: Start================================")

	var req gameEquipQueryRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("userID=", req.UserID)

	// Initialize query parameters
	params := db.ListInventoriesByUserIDParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
		UserID: req.UserID,
	}

	// Get data from database
	inventories, err := server.store.ListInventoriesByUserID(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	fmt.Println("inventories count:", len(inventories))
	ctx.JSON(http.StatusOK, inventories)

	fmt.Println("================================gameEquipQueryHandler: End================================")
}

/* Game-equip POST received data */
type gameEquipUpdateRequest struct {
	UserID  int64   `json:"userId"`
	ItemIDs []int64 `json:"itemId"`
}

/* Game-equip POST handle function */
func (server *Server) gameEquipUpdateHandler(ctx *gin.Context) {
	fmt.Println("================================gameEquipUpdateHandler: Start================================")

	var req gameEquipUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("UserID=", req.UserID)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user
		_, err := server.store.GetUserById(ctx, req.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				// If userID not exists, return err
				fmt.Println("User not exists")
				ctx.JSON(http.StatusNotFound, commonResponse("UserID:["+strconv.FormatInt(req.UserID, 10)+"]not exists!"))
				return nil
			} else {
				return err
			}
		}

		// Iterate ItemIDs
		for i := 0; i < len(req.ItemIDs); i++ {
			// Read inventory
			inventory, err := server.store.GetInventoryByUserIDItemID(ctx, db.GetInventoryByUserIDItemIDParams{
				UserID: req.UserID,
				ItemID: req.ItemIDs[i],
			})
			if err != nil {
				if err == sql.ErrNoRows {
					// If inventory not exists, return err
					fmt.Println("Inventory not exists")
					ctx.JSON(http.StatusNotFound, commonResponse("Inventory:["+strconv.FormatInt(req.ItemIDs[i], 10)+"]not exists!"))
					return nil
				} else {
					return err
				}
			}

			// Validate quantity
			if inventory.Quantity <= 0 {
				// If inventory quantity is not enough, return err
				fmt.Println("Inventory is not enough")
				ctx.JSON(http.StatusConflict, commonResponse("["+strconv.FormatInt(req.UserID, 10)+"]-inventory:["+strconv.FormatInt(req.ItemIDs[i], 10)+"]is not enough!"))
				return nil
			}

			// Update user's inventories
			inventory, err = server.store.UpdateInventoryQuantity(ctx, db.UpdateInventoryQuantityParams{
				Amount: -1,
				UserID: inventory.UserID,
				ItemID: inventory.ItemID,
			})
			if err != nil {
				return err
			}
			fmt.Println("Updated inventory.Quantity=", inventory.Quantity)
		}

		ctx.JSON(http.StatusOK, nil)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================gameEquipUpdateHandler: End================================")
}

/* Game-end POST received data */
type gameEndRequest struct {
	UserID int64 `json:"userId"`
	Credit int32 `json:"credit"`
	Coupon int32 `json:"coupon"`
}

/* Game-end POST handle function */
func (server *Server) gameEndHandler(ctx *gin.Context) {
	fmt.Println("================================gameEndHandler: Start================================")

	var req gameEndRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("UserID=", req.UserID)
	fmt.Println("Credit=", req.Credit)
	fmt.Println("Coupon=", req.Coupon)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user
		user, err := server.store.GetUserById(ctx, req.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				// If userID not exists, return err
				fmt.Println("User not exists")
				ctx.JSON(http.StatusNotFound, commonResponse("UserID:["+strconv.FormatInt(req.UserID, 10)+"]not exists!"))
				return nil
			} else {
				return err
			}
		}

		// Update user's Credit
		user, err = server.store.UpdateUserCredit(ctx, db.UpdateUserCreditParams{
			Amount: req.Credit,
			UserID: req.UserID,
		})
		if err != nil {
			return err
		}
		fmt.Println("Updated user.Credit=", user.Credit)

		// Update user's Coupon
		user, err = server.store.UpdateUserCoupon(ctx, db.UpdateUserCouponParams{
			Amount: req.Coupon,
			UserID: req.UserID,
		})
		if err != nil {
			return err
		}
		fmt.Println("Updated user.Coupon=", user.Coupon)

		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================gameEndHandler: End================================")
}
