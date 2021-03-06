package data

import (
	"context"
	"realworld/internal/biz"
	"time"
)

const TableUser = "user"

type User struct {
	Id         int       `gorm:"primaryKey"`
	Email      string    `gorm:"column:email"`
	Token      string    `gorm:"column:token"`
	UserName   string    `gorm:"column:username"`
	Bio        string    `gorm:"embedded" gorm:"column:bio"`
	Image      string    `gorm:"embedded" gorm:"column:image"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

// CreateUser 创建用户.
func (r *repo) CreateUser(ctx context.Context, user *biz.UserEntity) (*biz.UserEntity, error) {
	u := &User{
		Email:      user.Email,
		Token:      user.Token,
		UserName:   user.UserName,
		Bio:        user.Bio,
		Image:      user.Image,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	result := r.DB.WithContext(ctx).Table(TableUser).Create(u)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// GetUserByName 根据用户名获取用户.
func (r *repo) GetUserByName(ctx context.Context, username string) (*biz.UserEntity, error) {
	u := User{}
	result := r.DB.WithContext(ctx).Table(TableUser).Where("username = ?", username).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.UserEntity{
		Id:         u.Id,
		Email:      u.Email,
		Token:      u.Token,
		UserName:   u.UserName,
		Bio:        u.Bio,
		Image:      u.Image,
		CeateTime:  u.CreateTime,
		UpdateTime: u.UpdateTime,
	}, nil
}

// GetUserByEmail 根据用户名获取用户.
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*biz.UserEntity, error) {
	u := User{}
	result := r.DB.WithContext(ctx).Table(TableUser).Where("email = ?", email).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.UserEntity{
		Id:         u.Id,
		Email:      u.Email,
		Token:      u.Token,
		UserName:   u.UserName,
		Bio:        u.Bio,
		Image:      u.Image,
		CeateTime:  u.CreateTime,
		UpdateTime: u.UpdateTime,
	}, nil
}

func (r *repo) UpdateUserByEmail(ctx context.Context, email string, user *biz.UserEntity) (*biz.UserEntity, error) {
	return nil, nil
}
