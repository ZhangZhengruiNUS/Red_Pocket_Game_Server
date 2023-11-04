package db

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new instance of CreateItemParams with validation and assignment
func CreateItemParamsFactory(store *Store, ctx *gin.Context, itemName string, describe string, picPath string, price int32, creatorID int64) (int, CreateItemParams, error) {
	// Initialize variable
	var newCreateItemParams CreateItemParams

	// Check the uniqueness of the itemName
	_, err := store.GetItemByItemName(ctx, itemName)
	if err == nil {
		return http.StatusConflict, newCreateItemParams, errors.New("itemName[" + itemName + "] has been used")
	}
	if err != nil && err != sql.ErrNoRows {
		return http.StatusInternalServerError, newCreateItemParams, err
	}

	// Check CreatorID exists
	_, httpStatus, err := CheckUserExistsByID(store, ctx, creatorID)
	if err != nil {
		return httpStatus, newCreateItemParams, err
	}

	// Create an instance
	newCreateItemParams = CreateItemParams{
		ItemName:   itemName,
		Describe:   describe,
		PicPath:    picPath,
		Price:      price,
		CreatorID:  creatorID,
		CreateTime: time.Now(),
	}

	return http.StatusOK, newCreateItemParams, nil
}

// Create a new instance of UpdateItemParams with validation and assignment
func CreateUpdateItemParamsFactory(store *Store, ctx *gin.Context, itemID int64, itemName string, describe string, picPath string, price int32, reviserid int64) (int, UpdateItemParams, error) {
	// Initialize variable
	var newUpdateItemParams UpdateItemParams

	// Check ItemID exists
	_, err := store.GetItem(ctx, itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, newUpdateItemParams, errors.New("ItemID[" + strconv.FormatInt(itemID, 10) + "] not exists")
		}
		return http.StatusInternalServerError, newUpdateItemParams, err
	}

	// Check the uniqueness of the itemName
	item, err := store.GetItemByItemName(ctx, itemName)
	if err == nil && item.ItemID != itemID {
		return http.StatusConflict, newUpdateItemParams, errors.New("itemName[" + itemName + "] has been used by itemID[" + strconv.FormatInt(itemID, 10) + "]")
	}
	if err != nil && err != sql.ErrNoRows {
		return http.StatusInternalServerError, newUpdateItemParams, err
	}

	// Check Reviserid exists
	_, httpStatus, err := CheckUserExistsByID(store, ctx, reviserid)
	if err != nil {
		return httpStatus, newUpdateItemParams, err
	}

	// Create an instance
	newUpdateItemParams = UpdateItemParams{
		ItemID:     itemID,
		ItemName:   itemName,
		Describe:   describe,
		PicPath:    picPath,
		Price:      price,
		Reviserid:  reviserid,
		Revisetime: time.Now(),
	}

	return http.StatusOK, newUpdateItemParams, nil
}
