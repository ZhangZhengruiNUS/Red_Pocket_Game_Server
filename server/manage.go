package server

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"log"
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
	log.Println("================================manageRolltableUpdateHandler: Start================================")

	var req manageRolltableUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("prizeName =", req.PrizeName)

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
					log.Println("Prize Created")
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
				log.Println("Prize Deleted")
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
				log.Println("Prize Updated")
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

	log.Println("================================manageRolltableUpdateHandler: End================================")
}

/* Manage Catalog Update PUT received data */
type manageCatalogUpdateRequest struct {
	Status     string `json:"status"`
	ItemId     int64  `json:"itemId"`
	ItemName   string `json:"itemName"`
	Describe   string `json:"describe"`
	Price      int32  `json:"price"`
	PicPath    string `json:"picPath"`
	OperatorId int64  `json:"opeatorId"`
}

/*  Manage Catalog Update PUT handle function */
func (server *Server) manageCatalogUpdateHandler(ctx *gin.Context) {
	log.Println("================================manageCatalogUpdateHandler: Start================================")

	var req manageCatalogUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("itemName =", req.ItemName)

	switch req.Status {
	case "create":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Create a new instance of CreateItemParams
			httpStatus, createItemParams, err := db.CreateItemParamsFactory(server.store, ctx, req.ItemName, req.Describe, req.PicPath, req.Price, req.OperatorId)
			if err != nil {
				ctx.JSON(httpStatus, errorResponse(err))
				return err
			}
			// Insert an item in the DB
			item, err := server.store.CreateItem(ctx, createItemParams)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return err
			}
			log.Println("Item[" + item.ItemName + "] Created by [" + strconv.FormatInt(req.OperatorId, 10) + "]")
			ctx.JSON(http.StatusCreated, item)
			return nil
		})

		if err != nil {
			return
		}

	case "update":
		// Start database transaction
		err := server.store.ExecTx(ctx, func(q *db.Queries) error {
			// Create a new instance of UpdateItemParams
			httpStatus, updateItemParams, err := db.CreateUpdateItemParamsFactory(server.store, ctx, req.ItemId, req.ItemName, req.Describe, req.PicPath, req.Price, req.OperatorId)
			if err != nil {
				ctx.JSON(httpStatus, errorResponse(err))
				return err
			}
			// Update the item in the DB
			item, err := server.store.UpdateItem(ctx, updateItemParams)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return err
			}
			log.Println("Item[" + strconv.FormatInt(item.ItemID, 10) + "] Updated by [" + strconv.FormatInt(req.OperatorId, 10) + "]")
			ctx.JSON(http.StatusOK, item)
			return nil
		})
		// Return response
		if err != nil {
			return
		}
	}

	log.Println("================================manageCatalogUpdateHandler: End================================")
}

/*  Manage Catalog DELETE handle function */
func (server *Server) manageCatalogDeleteHandler(ctx *gin.Context) {
	log.Println("================================manageCatalogDeleteHandler: Start================================")

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
	log.Println("itemID =", itemID)

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
			log.Println("Item Deleted")
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

	log.Println("================================manageCatalogDeleteHandler: End================================")
}

/* Manage Diff GET handle function */
func (server *Server) manageDiffQueryHandler(ctx *gin.Context) {
	log.Println("================================manageDiffQueryHandler: Start================================")

	// Initialize query parameters
	params := db.ListGameDiffSetsParams{
		Page:     1,    //start record page
		Pagesize: 1000, //return record quantity
	}

	// Get data from database
	gameDifficultySettings, err := server.store.ListGameDiffSets(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	log.Println("count=", len(gameDifficultySettings))
	ctx.JSON(http.StatusOK, gameDifficultySettings)

	log.Println("================================manageDiffQueryHandler: End================================")
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
	log.Println("================================manageDiffUpdateHandler: Start================================")

	var req manageDiffUpdateRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("diffLV =", req.DiffLV)

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
					log.Println("diffLv Created")
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
				log.Println("Prize Deleted")
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
				log.Println("DiffLv Updated")
			}

			ctx.JSON(http.StatusOK, diffLv)
			return nil
		})
		// Return response
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		log.Println("================================manageDiffUpdateHandler: End================================")
	}
}
