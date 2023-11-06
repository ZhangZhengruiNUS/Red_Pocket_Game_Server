package server

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
	fmt.Println("userID =", userID)

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

/* Warehouse Rolltable Query GET handle function */
func (server *Server) warehouseRolltableQueryHandler(ctx *gin.Context) {
	fmt.Println("================================warehouseRolltableHandler: Start================================")

	// Initialize query parameters
	params := db.ListRolltableParams{
		Page:     1,    //start record page
		Pagesize: 1000, //return record quantity
	}

	// Get data from database
	prizes, err := server.store.ListRolltable(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Define a struct that represents the key-value pair of prizes
	type PrizesKeyValuePair struct {
		Prizes []db.ListRolltableRow `json:"data"`
	}

	// Initialize the key-value pair of prizes
	prizesKeyValuePair := PrizesKeyValuePair{
		Prizes: prizes,
	}

	// Return response
	fmt.Println("prizes count:", len(prizes))
	ctx.JSON(http.StatusOK, prizesKeyValuePair)

	fmt.Println("================================warehouseRolltableHandler: End================================")
}

/* Warehouse Rolltable Update POST received data */
type warehouseRolltableUpdateRequest struct {
	UserID int64 `json:"userId"`
}

/*  Warehouse Rolltable Update POST handle function */
func (server *Server) warehouseRolltableUpdateHandler(ctx *gin.Context) {
	fmt.Println("================================warehouseRolltableUpdateHandler: Start================================")

	var req warehouseRolltableUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("UserID =", req.UserID)

	// Get data of Coupon from database
	coupon, err := server.store.ListCouponByUserID(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("Current prizes count:", coupon)

	// Validate coupon
	if coupon <= 0 {
		// If coupon is not enough, return err
		fmt.Println("Coupon is not enough")
		ctx.JSON(http.StatusPaymentRequired, errorCustomResponse("User ["+strconv.FormatInt(req.UserID, 10)+"]'s coupon is not enough!"))
		return
	}

	// Update user's coupon
	coupon, err = server.store.UpdateCoupon(ctx, db.UpdateCouponParams{
		Amount: -1,
		UserID: req.UserID,
	})
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, nil)
	fmt.Println("Updated prizes count:", coupon)

	fmt.Println("================================warehouseRolltableUpdateHandler: End================================")
}

// For test:
// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Test 01"})
// return
