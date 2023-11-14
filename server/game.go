package server

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/* GameModule */
type GameModule struct{}

func (m *GameModule) RegisterRoutes(server *Server, router *gin.Engine) {
	users := router.Group("/game")
	{
		users.GET("/diff", server.gameDiffHandler)
		users.GET("/equip", server.gameEquipQueryHandler)
		users.POST("/equip", server.gameEquipUpdateHandler)
		users.POST("/end", server.gameEndHandler)
	}
}

/* Game-diff GET handle function */
func (server *Server) gameDiffHandler(ctx *gin.Context) {
	log.Println("================================gameDiffHandler: Start================================")

	// Initialize query parameters
	params := db.ListGameDiffSetsParams{
		Page:     1,    //start record page
		Pagesize: 1000, //return record quantity
	}

	// Read & Check input data
	userIDStr := strings.TrimSpace(ctx.Query("userId"))
	log.Println("userIDStr=", userIDStr)
	if len(userIDStr) == 0 {
		ctx.JSON(http.StatusBadRequest, errorCustomResponse("userIDStr is empty"))
		return
	}
	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check user exists
	user, httpStatus, err := db.CheckUserExistsByID(server.store, ctx, userIDInt)
	if err != nil {
		ctx.JSON(httpStatus, errorResponse(err))
		return
	}

	// Get data from database
	gameDifficultySettings, err := server.store.ListGameDiffSets(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	log.Println("gameDifficultySettings count:", len(gameDifficultySettings))
	if len(gameDifficultySettings) == 0 {
		ctx.JSON(http.StatusBadRequest, errorCustomResponse("No difficulty settings found"))
		return
	}

	// Extra adjust game difficulty settings by user
	extraDifficultyStrategy, err := getExtraGameDifficultySettingMode(server, ctx, user)
	if err != nil {
		return
	}
	extraDifficultyStrategy.AdjustGameSetting(gameDifficultySettings)
	// Return response
	ctx.JSON(http.StatusOK, gameDifficultySettings)

	log.Println("================================gameDiffHandler: End================================")
}

/* Game-equip GET received data */
type gameEquipQueryRequest struct {
	UserID int64 `json:"userId"`
}

/* Game-equip GET handle function */
func (server *Server) gameEquipQueryHandler(ctx *gin.Context) {
	log.Println("================================gameEquipQueryHandler: Start================================")

	// Read frontend data
	userID, err := strconv.ParseInt(ctx.Query("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("userID=", userID)

	// Initialize query parameters
	params := db.ListInventoriesByUserIDParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
		UserID: userID,
	}

	// Get data from database
	inventories, err := server.store.ListInventoriesByUserID(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	log.Println("inventories count:", len(inventories))
	ctx.JSON(http.StatusOK, inventories)

	log.Println("================================gameEquipQueryHandler: End================================")
}

/* Game-equip POST received data */
type gameEquipUpdateRequest struct {
	UserID  int64   `json:"userId"`
	ItemIDs []int64 `json:"itemId"`
}

/* Game-equip POST handle function */
func (server *Server) gameEquipUpdateHandler(ctx *gin.Context) {
	log.Println("================================gameEquipUpdateHandler: Start================================")

	var req gameEquipUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("UserID=", req.UserID)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user
		_, err := q.GetUserById(ctx, req.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				// If userID not exists, return err
				log.Println("User not exists")
				ctx.JSON(http.StatusNotFound, errorCustomResponse("UserID:["+strconv.FormatInt(req.UserID, 10)+"]not exists!"))
				return nil
			} else {
				return err
			}
		}

		// Iterate ItemIDs
		for i := 0; i < len(req.ItemIDs); i++ {
			// Read inventory
			inventory, err := q.GetInventoryByUserIDItemID(ctx, db.GetInventoryByUserIDItemIDParams{
				UserID: req.UserID,
				ItemID: req.ItemIDs[i],
			})
			if err != nil {
				if err == sql.ErrNoRows {
					// If inventory not exists, return err
					log.Println("Inventory not exists")
					ctx.JSON(http.StatusNotFound, errorCustomResponse("Inventory:["+strconv.FormatInt(req.ItemIDs[i], 10)+"]not exists!"))
					return nil
				} else {
					return err
				}
			}

			// Validate quantity
			if inventory.Quantity <= 0 {
				// If inventory quantity is not enough, return err
				log.Println("Inventory is not enough")
				ctx.JSON(http.StatusConflict, errorCustomResponse("["+strconv.FormatInt(req.UserID, 10)+"]-inventory:["+strconv.FormatInt(req.ItemIDs[i], 10)+"]is not enough!"))
				return nil
			}

			// Update user's inventories
			inventory, err = q.UpdateInventoryQuantity(ctx, db.UpdateInventoryQuantityParams{
				Amount: -1,
				UserID: inventory.UserID,
				ItemID: inventory.ItemID,
			})
			if err != nil {
				return err
			}
			log.Println("Updated inventory.Quantity=", inventory.Quantity)
		}

		ctx.JSON(http.StatusOK, nil)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("================================gameEquipUpdateHandler: End================================")
}

/* Game-end POST received data */
type gameEndRequest struct {
	UserID int64 `json:"userId"`
	Credit int32 `json:"credit"`
	Coupon int32 `json:"coupon"`
}

/* Game-end POST handle function */
func (server *Server) gameEndHandler(ctx *gin.Context) {
	log.Println("================================gameEndHandler: Start================================")

	var req gameEndRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("UserID=", req.UserID)
	log.Println("Credit=", req.Credit)
	log.Println("Coupon=", req.Coupon)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user
		user, err := q.GetUserById(ctx, req.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				// If userID not exists, return err
				log.Println("User not exists")
				ctx.JSON(http.StatusNotFound, errorCustomResponse("UserID:["+strconv.FormatInt(req.UserID, 10)+"]not exists!"))
				return nil
			} else {
				return err
			}
		}

		// Update user's Credit
		user, err = q.UpdateUserCredit(ctx, db.UpdateUserCreditParams{
			Amount: req.Credit,
			UserID: req.UserID,
		})
		if err != nil {
			return err
		}
		log.Println("Updated user.Credit=", user.Credit)

		// Update user's Coupon
		user, err = q.UpdateUserCoupon(ctx, db.UpdateUserCouponParams{
			Amount: req.Coupon,
			UserID: req.UserID,
		})
		if err != nil {
			return err
		}
		log.Println("Updated user.Coupon=", user.Coupon)

		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("================================gameEndHandler: End================================")
}
