package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// AllConf 所有配置文件.
var AllConf *Conf

type Conf struct {
	BootstrapConf Bootstrap
	DBConf        DBConfs
	LogConf       LogConf
}

// ReadAllConf 根据目录或文件读取yaml配置
func ReadAllConf(confPath string) {
	c := config.New(
		config.WithSource(
			file.NewSource(confPath),
		),
	)
	defer c.Close()
	AllConf = &Conf{}

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(&AllConf.BootstrapConf); err != nil {
		panic(err)
	}

	//数据库配置
	if err := c.Scan(&AllConf.DBConf); err != nil {
		panic(err)
	}

	//	日志配置
	if err := c.Scan(&AllConf.LogConf); err != nil {
		panic(err)
	}

}
