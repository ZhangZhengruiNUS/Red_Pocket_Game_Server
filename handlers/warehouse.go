package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* Warehouse GET received data */
type warehouseQueryRequest struct {
	UserID int64 `json:"userId"`
}

// Define a struct that represents the combined warehouse data
type CombinedWarehouseData struct {
	Warehouse01 db.ListWarehouse01ByUserIDRow   `json:"warehouse01"`
	Warehouse02 []db.ListInventoriesByUserIDRow `json:"warehouse02"`
}

/* Warehouse GET handle function */
func (server *Server) warehouseInfoQueryHandler(ctx *gin.Context) {
	fmt.Println("================================warehouseQueryHandler: Start================================")

	// Read frontend data
	// userID, err := strconv.ParseInt(ctx.Query("userId"), 10, 64)

	// Read frontend data
	userIDStr := ctx.Query("userId")
	if userIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("userID=", userID)

	// Initialize query parameters
	// params := db.ListWarehouse02ByUserIDParams{
	// 	Limit:  1000, //return record quantity
	// 	Offset: 0,    //start record index
	// 	UserID: userID,
	// }
	params := db.ListInventoriesByUserIDParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
		UserID: userID,
	}

	// Get data from database
	warehouse01, err := server.store.ListWarehouse01ByUserID(ctx, userID)
	// ctx.JSON(http.StatusOK, warehouse01)
	// return

	// if userID == 2 {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Test 03"})
	// 	return
	// }

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	warehouse02, err := server.store.ListInventoriesByUserID(ctx, params)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Initialize the combined data structure
	combinedData := CombinedWarehouseData{
		Warehouse01: warehouse01,
		Warehouse02: warehouse02,
	}

	// Return response
	ctx.JSON(http.StatusOK, combinedData)

	fmt.Println("================================warehouseQueryHandler: End================================")
}
