package data

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"realworld/internal/biz"
	"realworld/pkg/dbresolver"
)

var ProviderSet = wire.NewSet(NewRepo)

type repo struct {
	DB *gorm.DB
}

func NewRepo(resolver *dbresolver.Resolver) biz.Repo {
	return &repo{
		DB: resolver.ResolveDbEngine("realworld"),
	}
}

var _ biz.Repo = (*repo)(nil)
