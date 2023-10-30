package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* Warehouse GET handle function */
func (server *Server) warehouseInfoQueryHandler(ctx *gin.Context) {
	fmt.Println("================================warehouseQueryHandler: Start================================")

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

	// Get data from database
	creditAndCoupon, err := server.store.ListWarehouse01ByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	params := db.ListWarehouse02ByUserIDParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
		UserID: userID,
	}

	equipts, err := server.store.ListWarehouse02ByUserID(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Define a struct that represents the combined warehouse data
	type CombinedWarehouseData struct {
		Credit  int32                           `json:"credit"`
		Coupon  int32                           `json:"coupon"`
		Equipts []db.ListWarehouse02ByUserIDRow `json:"equipts"`
	}

	// Initialize the combined data structure
	combinedData := CombinedWarehouseData{
		Credit:  creditAndCoupon.Credit,
		Coupon:  creditAndCoupon.Coupon,
		Equipts: equipts,
	}

	// Return response
	ctx.JSON(http.StatusOK, combinedData)

	fmt.Println("================================warehouseQueryHandler: End================================")
}

/* Warehouse Rolltable GET handle function */
func (server *Server) warehouseRolltableHandler(ctx *gin.Context) {
	fmt.Println("================================warehouseRolltableHandler: Start================================")

	// Initialize query parameters
	params := db.ListRolltableParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
	}

	// Get data from database
	prizes, err := server.store.ListRolltable(ctx, params)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	fmt.Println("prizes count:", len(prizes))
	ctx.JSON(http.StatusOK, prizes)

	fmt.Println("================================warehouseRolltableHandler: End================================")
}

// For test
// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Test 01"})
// return
