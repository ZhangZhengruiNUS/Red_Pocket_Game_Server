package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// receive data
type loginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (server *Server) loginHandler(ctx *gin.Context) {
	var req loginRequest
	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userName := req.UserName
	password := req.Password

	// temporary
	fmt.Println("password=" + password)
	fmt.Println("username=" + userName)

	account, err := server.store.GetUserByName(ctx, userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if account.Password == password {
		fmt.Println("password right")
		ctx.JSON(http.StatusOK, account)
	} else {
		fmt.Println("password wrong")
		ctx.JSON(http.StatusUnauthorized, commonResponse("userName & password denied"))
	}

}
