package server

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* Catalog GET handle function */
func (server *Server) catalogHandler(ctx *gin.Context) {
	log.Println("================================catalogHandler: Start================================")

	// Initialize query parameters
	params := db.ListItemsParams{
		Page:     1,    //start record page
		Pagesize: 1000, //return record quantity
	}

	// Get data from database
	items, err := server.store.ListItems(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	log.Println("items count:", len(items))
	ctx.JSON(http.StatusOK, items)

	log.Println("================================catalogHandler: End================================")
}

/* Catalog-buy received data */
type catalogBuyRequest struct {
	UserID int64 `json:"userId"`
	ItemID int64 `json:"itemId"`
}

/* Catalog-buy Post handle function */
func (server *Server) catalogBuyHandler(ctx *gin.Context) {
	log.Println("================================catalogBuyHandler: Start================================")

	var req catalogBuyRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("UserID=", req.UserID)
	log.Println("ItemID=", req.ItemID)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user's credit
		user, err := server.store.GetUserById(ctx, req.UserID)
		if err != nil {
			return err
		}
		log.Println("user.Credit=", user.Credit)

		// Read item's price
		item, err := server.store.GetItem(ctx, req.ItemID)
		if err != nil {
			return err
		}
		log.Println("item.Price=", item.Price)

		// Vaildate credit >= price
		if user.Credit < item.Price {
			ctx.JSON(http.StatusPaymentRequired, errorCustomResponse("Credit is not enough!"))
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
		log.Println("Updated user.Credit=", user.Credit)

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
				log.Println("Inventory Created")
			} else {
				return err
			}
		} else {
			// If inventory exists, update inventory's quantity
			inventory, err = server.store.UpdateInventoryQuantity(ctx, db.UpdateInventoryQuantityParams{
				Amount: 1,
				UserID: inventory.UserID,
				ItemID: inventory.ItemID,
			})
			if err != nil {
				return err
			}
			log.Println("Updated inventory.Quantity=", inventory.Quantity)

		}
		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("================================catalogBuyHandler: End================================")
}

/* Catalog-user received data */
type catalogUserRequest struct {
	UserID int64 `json:"userId"`
}

/* Catalog-user GET handle function */
func (server *Server) catalogUserHandler(ctx *gin.Context) {
	log.Println("================================catalogUserHandler: Start================================")

	// Read frontend data
	userID, err := strconv.ParseInt(ctx.Query("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("userID=", userID)

	// Get data from database
	user, err := server.store.GetUserById(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, user)

	log.Println("================================catalogUserHandler: End================================")
}
