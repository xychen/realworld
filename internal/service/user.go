package service

/*
import (
	"context"
	realworldv1 "realworld/api/realworld/v1"
	"realworld/internal/biz"
	"realworld/utils"
)

type UserService struct {
	biz biz.Biz
}

func NewUserService(biz biz.Biz) *UserService {
	return &UserService{
		biz: biz,
	}
}

// Register 用户注册.
func (u *UserService) Register(ctx context.Context, req *realworldv1.RegisterRequest) (*realworldv1.UserReply, error) {
	user := &biz.UserEntity{
		Email:    req.User.Email,
		UserName: req.User.Username,
	}
	user.Token = utils.MD5([]byte(req.User.Password))
	user, err := u.biz.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &realworldv1.UserReply{
		User: &realworldv1.UserReply_User{
			Email:    user.Email,
			Token:    user.Token,
			Username: user.UserName,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}
*/
