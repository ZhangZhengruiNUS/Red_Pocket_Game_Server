package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Catalog GET handle function */
func (server *Server) catalogHandler(ctx *gin.Context) {
	fmt.Println("================================catalogHandler: Start================================")

	// Initialize query parameters
	params := db.ListItemsParams{
		Limit:  1000, //return record quantity
		Offset: 0,    //start record index
	}

	// Get data from database
	items, err := server.store.ListItems(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	fmt.Println("items count:", len(items))
	ctx.JSON(http.StatusOK, items)

	fmt.Println("================================catalogHandler: End================================")
}

/* Catalog-buy received data */
type catalogBuyRequest struct {
	UserID int64 `json:"userId"`
	ItemID int64 `json:"itemId"`
}

/* Catalog-buy Post handle function */
func (server *Server) catalogBuyHandler(ctx *gin.Context) {
	fmt.Println("================================catalogBuyHandler: Start================================")

	var req catalogBuyRequest

	// Read frontend data
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(string(bodyBytes))
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("UserID=", req.UserID)
	fmt.Println("ItemID=", req.ItemID)

	// Start database transaction
	err = server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user's credit
		user, err := server.store.GetUserById(ctx, req.UserID)
		if err != nil {
			return err
		}
		fmt.Println("user.Credit=", user.Credit)

		// Read item's price
		item, err := server.store.GetItem(ctx, req.ItemID)
		if err != nil {
			return err
		}
		fmt.Println("item.Price=", item.Price)

		// Vaildate credit >= price
		if user.Credit < item.Price {
			ctx.JSON(http.StatusPaymentRequired, commonResponse("Credit is not enough!"))
			return nil
		}

		// Update user's credit
		user, err = server.store.UpdateUserCredit(ctx, db.UpdateUserCreditParams{
			Amount: -item.Price,
			UserID: user.UserID,
		})
		if err != nil {
			return err
		}
		fmt.Println("Updated user.Credit=", user.Credit)

		// Read inventory
		inventory, err := server.store.GetInventoryByUserIDItemID(ctx, db.GetInventoryByUserIDItemIDParams{
			UserID: user.UserID,
			ItemID: item.ItemID,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				// If inventory not exists, create it
				inventory, err = server.store.CreateInventory(ctx, db.CreateInventoryParams{
					UserID:   user.UserID,
					ItemID:   item.ItemID,
					Quantity: 1,
				})
				if err != nil {
					return err
				}
				fmt.Println("Inventory Created")
			} else {
				return err
			}

		} else {
			// If inventory not exists, update inventory's quantity
			inventory, err = server.store.UpdateInventoryQuantity(ctx, db.UpdateInventoryQuantityParams{
				Amount: 1,
				UserID: inventory.UserID,
				ItemID: inventory.ItemID,
			})
			if err != nil {
				return err
			}
			fmt.Println("Updated inventory.Quantity=", inventory.Quantity)

		}
		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================catalogBuyHandler: End================================")
}

/* Catalog-user received data */
type catalogUserRequest struct {
	UserID int64 `json:"userId"`
}

/* Catalog-user GET handle function */
func (server *Server) catalogUserHandler(ctx *gin.Context) {
	fmt.Println("================================loginHandler: Start================================")

	var req catalogUserRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("userID=", req.UserID)

	// Get data from database
	user, err := server.store.GetUserById(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, user)

	fmt.Println("================================loginHandler: End================================")
}
