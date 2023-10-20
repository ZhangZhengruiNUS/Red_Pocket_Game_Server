package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"fmt"
	"net/http"

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
