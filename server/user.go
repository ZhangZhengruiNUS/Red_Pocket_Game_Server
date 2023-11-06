package server

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* UserModule */
type UserModule struct{}

func (m *UserModule) RegisterRoutes(server *Server, router *gin.Engine) {
	users := router.Group("")
	{
		users.POST("/login", server.loginHandler)
		users.POST("/signup", server.signupHandler)
	}
}

/* Login received data */
type loginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

/* Login handle function */
func (server *Server) loginHandler(ctx *gin.Context) {
	fmt.Println("================================loginHandler: Start================================")

	var req loginRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("username=", req.UserName)

	// Get data from database
	user, err := server.store.GetUserByName(ctx, req.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Verify password and return response
	if user.Password == req.Password {
		fmt.Println("password right")
		ctx.JSON(http.StatusOK, user)
	} else {
		fmt.Println("password wrong")
		ctx.JSON(http.StatusUnauthorized, errorCustomResponse("userName & password denied"))
	}

	fmt.Println("================================loginHandler: End================================")
}

/* Signup received data */
type signupRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

/* Signup handle function */
func (server *Server) signupHandler(ctx *gin.Context) {
	fmt.Println("================================signupHandler: Start================================")

	var req signupRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("username=", req.UserName)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user
		user, err := q.GetUserByName(ctx, req.UserName)
		if err != nil {
			if err == sql.ErrNoRows {
				// If username not exists, create it
				user, err = q.CreateUser(ctx, db.CreateUserParams{
					UserName: req.UserName,
					Password: req.Password,
					RoleType: 0,
				})
				if err != nil {
					return err
				}
				fmt.Println("User Created")
			} else {
				return err
			}
		} else {
			// If username exists, return err
			ctx.JSON(http.StatusConflict, errorCustomResponse("UserName:["+user.UserName+"] has already been signed up!"))
			return nil
		}

		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================signupHandler: End================================")
}
