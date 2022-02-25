package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBiz)

type Repo interface {
	CreateUser(user *UserEntity) (*UserEntity, error)
	GetUserByName(username string) (*UserEntity, error)
}

type Biz interface {
	CreateUser(user *UserEntity) (*UserEntity, error)
	GetUserByName(username string) (*UserEntity, error)
}

type biz struct {
	repo Repo
}

func (b *biz) CreateUser(user *UserEntity) (*UserEntity, error) {
	return b.repo.CreateUser(user)
}

func (b *biz) GetUserByName(username string) (*UserEntity, error) {
	return b.repo.GetUserByName(username)
}

func NewBiz(repo Repo) Biz {
	return &biz{
		repo: repo,
	}
}

var _ Biz = (*biz)(nil)
