package server

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handle error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Handle error custom response
func errorCustomResponse(msg string) gin.H {
	return errorResponse(errors.New(msg))
}

// Get extra game difficulty setting mode
func getExtraGameDifficultySettingMode(server *Server, ctx *gin.Context, user db.User) (ExtraDifficultyStrategy, error) {
	var extraDifficultyStrategy ExtraDifficultyStrategy
	// Get Average Coupon Count
	averageCouponCount, err := server.store.GetAverageCouponCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return extraDifficultyStrategy, err
	}
	log.Println("user.Coupon=", user.Coupon)
	log.Println("averageCouponCount=", averageCouponCount)

	// Judge extra game difficulty setting mode
	if float32(user.Coupon) > float32(averageCouponCount)*1.5 {
		// If (current user's coupon count)  > (average * 1.5), then adjust the difficulty setting to hard
		extraDifficultyStrategy = HardDiffStrategy{}
		log.Println("Hard mode")
	} else if float32(user.Coupon) <= float32(averageCouponCount)*0.5 {
		// If (current user's coupon count)  <= (average * 0.5), then adjust the difficulty setting to easy
		extraDifficultyStrategy = EasyDiffStrategy{}
		log.Println("Easy mode")
	} else {
		extraDifficultyStrategy = NormalDiffStrategy{}
		log.Println("Normal mode")
	}
	return extraDifficultyStrategy, nil
}
