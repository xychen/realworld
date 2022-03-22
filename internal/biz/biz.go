package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBiz)

type Repo interface {
	CreateUser(ctx context.Context, user *UserEntity) (*UserEntity, error)
	GetUserByName(ctx context.Context, username string) (*UserEntity, error)
	GetUserByEmail(ctx context.Context, email string) (*UserEntity, error)
}

type Biz interface {
	CreateUser(ctx context.Context, user *UserEntity) (*UserEntity, error)
	GetUserByName(ctx context.Context, username string) (*UserEntity, error)
	GetUserByEmail(ctx context.Context, email string) (*UserEntity, error)
}

type biz struct {
	repo Repo
}

func (b *biz) CreateUser(ctx context.Context, user *UserEntity) (*UserEntity, error) {
	return b.repo.CreateUser(ctx, user)
}

func (b *biz) GetUserByName(ctx context.Context, username string) (*UserEntity, error) {
	return b.repo.GetUserByName(ctx, username)
}

func (b *biz) GetUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	return b.repo.GetUserByEmail(ctx, email)
}

func NewBiz(repo Repo) Biz {
	return &biz{
		repo: repo,
	}
}

var _ Biz = (*biz)(nil)
