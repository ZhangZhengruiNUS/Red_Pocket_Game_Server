package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* Manage Rolltable Update POST received data */
type manageRolltableUpdateRequest struct {
	Status    string `json:"status"`
	PrizeName string `json:"prizeName"`
	PicPath   string `json:"picPath"`
	Weight    int32  `json:"weight"`
}

/*  Manage Rolltable Update POST handle function */
func (server *Server) manageRolltableUpdateHandler(ctx *gin.Context) {
	fmt.Println("================================manageRolltableUpdateHandler: Start================================")

	var req manageRolltableUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("prizeName =", req.PrizeName)

	switch req.Status {
	case "create":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			prize, err := server.store.GetPrizeByPrizeName(ctx, req.PrizeName)
			if err != nil {
				if err == sql.ErrNoRows {
					// If prize_name not exists, create it
					prize, err = server.store.CreatePrize(ctx, db.CreatePrizeParams{
						PrizeName: req.PrizeName,
						PicPath:   req.PicPath,
						Weight:    req.Weight,
					})
					if err != nil {
						return err
					}
					fmt.Println("Prize Created")
				} else {
					return err
				}
			} else {
				// If prize exists, return err
				ctx.JSON(http.StatusConflict, errorCustomResponse("PrizeName:["+prize.PrizeName+"] has already been created!"))
				return nil
			}

			ctx.JSON(http.StatusOK, nil)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	case "delete":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			prize, err := server.store.GetPrizeByPrizeName(ctx, req.PrizeName)
			if err != nil {
				if err == sql.ErrNoRows {
					// If prize_name not exists, return an error
					ctx.JSON(http.StatusConflict, errorCustomResponse("PrizeName:["+req.PrizeName+"] has not been created!"))
					return nil
				} else {
					return err
				}
			} else {
				// If prize exists, delete it
				err = server.store.DeletePrize(ctx, req.PrizeName)
				if err != nil {
					return err
				}
				fmt.Println("Prize Deleted")
			}

			ctx.JSON(http.StatusOK, prize)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	case "update":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			prize, err := server.store.GetPrizeByPrizeName(ctx, req.PrizeName)
			if err != nil {
				if err == sql.ErrNoRows {
					// If prize_name not exists, return an error
					ctx.JSON(http.StatusConflict, errorCustomResponse("PrizeName:["+req.PrizeName+"] has not been created!"))
					return nil
				} else {
					return err
				}
			} else {
				// If prize exists, update it
				prize, err = server.store.UpdatePrize(ctx, db.UpdatePrizeParams{
					Picpath:   req.PicPath,
					Weight:    req.Weight,
					PrizeName: req.PrizeName,
				})
				if err != nil {
					return err
				}
				fmt.Println("Prize Updated")
			}

			ctx.JSON(http.StatusOK, prize)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	fmt.Println("================================manageRolltableUpdateHandler: End================================")
}

/* Manage Catalog Update PUT received data */
type manageCatalogUpdateRequest struct {
	Status   string `json:"status"`
	ItemId   int64  `json:"itemId"`
	ItemName string `json:"itemName"`
	Describe string `json:"describe"`
	Price    int32  `json:"price"`
	PicPath  string `json:"picPath"`
}

/*  Manage Rolltable Update PUT handle function */
func (server *Server) manageCatalogUpdateHandler(ctx *gin.Context) {
	fmt.Println("================================manageCatalogUpdateHandler: Start================================")

	var req manageCatalogUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("itemName =", req.ItemName)

	switch req.Status {
	case "create":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			item, err := server.store.GetItemByItemName(ctx, req.ItemName)
			if err != nil {
				if err == sql.ErrNoRows {
					// If item_name not exists, create it
					item, err = server.store.CreateItem(ctx, db.CreateItemParams{
						ItemName:  req.ItemName,
						Describe:  req.Describe,
						PicPath:   req.PicPath,
						Price:     req.Price,
						CreatorID: -1,
					})
					if err != nil {
						return err
					}
					fmt.Println("Item Created")
				} else {
					return err
				}
			} else {
				// If prize exists, return err
				ctx.JSON(http.StatusConflict, errorCustomResponse("ItemName:["+req.ItemName+"] has already been created!"))
				return nil
			}

			ctx.JSON(http.StatusOK, item)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	case "update":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			item, err := server.store.GetItem(ctx, req.ItemId)
			if err != nil {
				if err == sql.ErrNoRows {
					// If item_name not exists, return an error
					ctx.JSON(http.StatusConflict, errorCustomResponse("ItemName:["+req.ItemName+"] has not been created!"))
					return nil
				} else {
					return err
				}
			} else {
				// If item exists, update it
				item, err = server.store.UpdateItem(ctx, db.UpdateItemParams{
					Itemname: req.ItemName,
					Describe: req.Describe,
					Price:    req.Price,
					Picpath:  req.PicPath,
					ItemID:   req.ItemId,
				})
				if err != nil {
					return err
				}
				fmt.Println("Item Updated")
			}

			ctx.JSON(http.StatusOK, item)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	fmt.Println("================================manageCatalogUpdateHandler: End================================")
}

/*  Manage Catalog DELETE handle function */
func (server *Server) manageCatalogDeleteHandler(ctx *gin.Context) {
	fmt.Println("================================manageCatalogDeleteHandler: Start================================")

	// Read frontend data
	itemIDStr := ctx.Query("itemId")
	if itemIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "itemId is required"})
		return
	}
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("itemID =", itemID)

	// Start database transaction
	err = server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read prize
		item, err := server.store.GetItem(ctx, itemID)
		if err != nil {
			if err == sql.ErrNoRows {
				// If item not exists, return an error
				ctx.JSON(http.StatusConflict, errorCustomResponse("ItemName:["+item.ItemName+"] has not been created!"))
				return nil
			} else {
				return err
			}
		} else {
			// If prize exists, delete it
			err = server.store.DeleteItem(ctx, itemID)
			if err != nil {
				return err
			}
			fmt.Println("Item Deleted")
		}

		ctx.JSON(http.StatusOK, nil)
		return nil
	})

	// Return response
	if err != nil {
		if err.Error() == "pq: update or delete on table \"items\" violates foreign key constraint \"inventories_item_id_fkey\" on table \"inventories\"" {
			ctx.JSON(http.StatusConflict, errorCustomResponse("This item is now possessed by users, so you can not delete it!"))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	fmt.Println("================================manageCatalogDeleteHandler: End================================")
}

/* Manage Diff GET handle function */
func (server *Server) manageDiffQueryHandler(ctx *gin.Context) {
	fmt.Println("================================catalogHandler: Start================================")

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
	fmt.Println("items count:", len(items))
	ctx.JSON(http.StatusOK, items)

	fmt.Println("================================catalogHandler: End================================")
}

/* Manage Diff Update POST received data */
type manageDiffUpdateRequest struct {
	Status       string `json:"status"`
	DiffLV       int32  `json:"diffLV"`
	AwardDensity int32  `json:"awardDensity"`
	EnemyDensity int32  `json:"enemyDensity"`
}

/*  Manage DiffLv Update POST handle function */
func (server *Server) manageDiffUpdateHandler(ctx *gin.Context) {
	fmt.Println("================================manageDiffUpdateHandler: Start================================")

	var req manageDiffUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("diffLV =", req.DiffLV)

	switch req.Status {
	case "create":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			diffLv, err := server.store.ListGameDiffSetsByDiffLv(ctx, req.DiffLV)
			if err != nil {
				if err == sql.ErrNoRows {
					// If diffLv not exists, create it
					diffLv, err = server.store.CreateDiffLv(ctx, db.CreateDiffLvParams{
						DiffLv:       req.DiffLV,
						AwardDensity: req.AwardDensity,
						EnemyDensity: req.EnemyDensity,
					})
					if err != nil {
						return err
					}
					fmt.Println("diffLv Created")
				} else {
					return err
				}
			} else {
				// If diffLv exists, return err
				ctx.JSON(http.StatusConflict, errorCustomResponse("This diffLv has already been created!"))
				return nil
			}

			ctx.JSON(http.StatusOK, diffLv)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	case "delete":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			diffLv, err := server.store.ListGameDiffSetsByDiffLv(ctx, req.DiffLV)
			if err != nil {
				if err == sql.ErrNoRows {
					// If diffLv not exists, return an error
					ctx.JSON(http.StatusConflict, errorCustomResponse("This diffLv has not been created!"))
					return nil
				} else {
					return err
				}
			} else {
				// If diffLv exists, delete it
				err = server.store.DeleteDiffLv(ctx, db.DeleteDiffLvParams{
					DiffLv:       req.DiffLV,
					AwardDensity: req.AwardDensity,
					EnemyDensity: req.EnemyDensity,
				})
				if err != nil {
					return err
				}
				fmt.Println("Prize Deleted")
			}

			ctx.JSON(http.StatusOK, diffLv)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	case "update":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Read prize
			diffLv, err := server.store.ListGameDiffSetsByDiffLv(ctx, req.DiffLV)
			if err != nil {
				if err == sql.ErrNoRows {
					// If diffLv not exists, return an error
					ctx.JSON(http.StatusConflict, errorCustomResponse("This diffLv has not been created!"))
					return nil
				} else {
					return err
				}
			} else {
				// If diffLv exists, update it
				diffLv, err = server.store.UpdateDiffLv(ctx, db.UpdateDiffLvParams{
					DiffLv:       req.DiffLV,
					Awarddensity: req.AwardDensity,
					Enemydensity: req.EnemyDensity,
				})
				if err != nil {
					return err
				}
				fmt.Println("DiffLv Updated")
			}

			ctx.JSON(http.StatusOK, diffLv)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		fmt.Println("================================manageDiffUpdateHandler: End================================")
	}
}
