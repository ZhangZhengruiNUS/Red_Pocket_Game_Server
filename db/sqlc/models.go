// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"database/sql"
	"time"
)

type GameDifficultySetting struct {
	DiffLv       int32         `json:"diffLv"`
	AwardDensity int32         `json:"awardDensity"`
	EnemyDensity int32         `json:"enemyDensity"`
	ReviserID    sql.NullInt64 `json:"reviserId"`
	ReviseTime   sql.NullTime  `json:"reviseTime"`
	CreatorID    int64         `json:"creatorId"`
	CreateTime   time.Time     `json:"createTime"`
}

type Inventory struct {
	UserID     int64     `json:"userId"`
	ItemID     int64     `json:"itemId"`
	Quantity   int32     `json:"quantity"`
	CreateTime time.Time `json:"createTime"`
}

type Item struct {
	ItemID     int64         `json:"itemId"`
	ItemName   string        `json:"itemName"`
	Describe   string        `json:"describe"`
	PicPath    string        `json:"picPath"`
	Price      int32         `json:"price"`
	ReviserID  sql.NullInt64 `json:"reviserId"`
	ReviseTime sql.NullTime  `json:"reviseTime"`
	CreatorID  int64         `json:"creatorId"`
	CreateTime time.Time     `json:"createTime"`
}

type Prize struct {
	PrizeID    int64         `json:"prizeId"`
	PrizeName  string        `json:"prizeName"`
	PicPath    string        `json:"picPath"`
	Weight     int32         `json:"weight"`
	ReviserID  sql.NullInt64 `json:"reviserId"`
	ReviseTime sql.NullTime  `json:"reviseTime"`
	CreatorID  int64         `json:"creatorId"`
	CreateTime time.Time     `json:"createTime"`
}

type User struct {
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Credit   int32  `json:"credit"`
	Coupon   int32  `json:"coupon"`
	// 0-normal user 1-admin
	RoleType   int32     `json:"roleType"`
	CreateTime time.Time `json:"createTime"`
}