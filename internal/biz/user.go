package biz

import "time"

// UserEntity 用户实体.
type UserEntity struct {
	Id         int
	Email      string
	Token      string
	UserName   string
	Bio        string
	Image      string
	CeateTime  time.Time
	UpdateTime time.Time
}
