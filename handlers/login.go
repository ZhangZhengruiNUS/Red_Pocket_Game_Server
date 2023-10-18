package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login input data
type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) loginHandler(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username := req.Username
	password := req.Password
	// temporary
	fmt.Println("password=" + password)
	fmt.Println("username=" + username)

	account, err := server.store.GetUserByName(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if account.Password == password {
		fmt.Println("password right")
		ctx.JSON(http.StatusOK, account)
	} else {
		fmt.Println("password wrong")
		ctx.JSON(http.StatusUnauthorized, commonResponse("username & password denied"))
	}

}
