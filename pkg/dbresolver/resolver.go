package dbresolver

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"realworld/internal/conf"
	"realworld/pkg/logger"
	"sync"
)

var ProviderSet = wire.NewSet(NewResolver)

type Resolver struct {
	dbMu      sync.RWMutex
	dbEngines map[string]*gorm.DB
	l         log.Logger
}

func NewResolver(l log.Logger) (*Resolver, error) {
	r := &Resolver{
		dbMu:      sync.RWMutex{},
		dbEngines: make(map[string]*gorm.DB),
		l:         l,
	}
	//@todo: 此处如何优雅处理一下
	r.initDb(conf.AllConf.DBConf)
	return r, nil
}

func (r *Resolver) initDb(dbconfs conf.DBConfs) error {
	for _, config := range dbconfs.Dbinfo {
		instanceName := config.Dbname
		//主库
		master := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", config.Master.User, config.Master.Pwd, config.Master.Host, config.Master.Port, config.Dbname, config.Charset))
		d, err := gorm.Open(master, &gorm.Config{
			Logger: &logger.DbLogger{
				Logger: r.l,
			},
		})

		if err != nil {
			return err
		}

		//从库
		var replicas []gorm.Dialector
		for _, c := range config.Replicas {
			replicas = append(replicas, mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", c.User, c.Pwd, c.Host, c.Port, config.Dbname, config.Charset)))
		}

		if err = d.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{master},
			Replicas: replicas,
			// sources/replicas 负载均衡策略
			Policy: dbresolver.RandomPolicy{},
		}).SetMaxIdleConns(int(config.MaxIdleConns)).SetMaxOpenConns(int(config.MaxOpenConns))); err != nil {
			return err
		}

		r.registerDbEngine(instanceName, d)
	}

	return nil
}

func (r *Resolver) registerDbEngine(instanceName string, db *gorm.DB) {
	r.dbMu.Lock()
	defer r.dbMu.Unlock()

	r.dbEngines[instanceName] = db
}

func (r *Resolver) ResolveDbEngine(instanceName string) *gorm.DB {
	r.dbMu.RLock()
	defer r.dbMu.RUnlock()

	if d, ok := r.dbEngines[instanceName]; ok {
		return d
	} else {
		panic(fmt.Sprintf("instanceName %s 不存在", instanceName))
	}
}
