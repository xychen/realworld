package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/biz"
	"realworld/utils"

	pb "realworld/api/realworld/v1"
)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer
	biz    biz.Biz
	logger log.Logger
}

func NewRealWorldService(biz biz.Biz, l log.Logger) *RealWorldService {
	return &RealWorldService{
		biz:    biz,
		logger: l,
	}
}

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	info, err := s.biz.GetUserByEmail(ctx, req.User.Email)
	if err != nil {
		return nil, err
	}
	if req.User.Password != info.Token {
		return nil, errors.New(403, "", "login failed")
	}
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	user := &biz.UserEntity{
		Email:    req.User.Email,
		UserName: req.User.Username,
	}
	user.Token = utils.MD5([]byte(req.User.Password))
	user, err := s.biz.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Email:    user.Email,
			Token:    user.Token,
			Username: user.UserName,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) unFollowUser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesReply, error) {
	return &pb.ListArticlesReply{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesReply, error) {
	return &pb.ListArticlesReply{}, nil
}
func (s *RealWorldService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) AddCommentsToArticle(ctx context.Context, req *pb.AddCommentsRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetCommentsFromArticle(ctx context.Context, req *pb.GetCommentsFromArticleRequest) (*pb.MultiCommentsReply, error) {
	return &pb.MultiCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComments(ctx context.Context, req *pb.DeleteCommentsRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
