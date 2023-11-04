package db

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Check user by userID exists
func CheckUserExistsByID(store *Store, ctx *gin.Context, userID int64) (User, int, error) {
	var user User
	user, err := store.GetUserById(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, http.StatusBadRequest, errors.New("userID[" + strconv.FormatInt(userID, 10) + "] not exists")
		}
		return user, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}
