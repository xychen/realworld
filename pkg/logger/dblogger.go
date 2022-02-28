package logger

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/logger"
	"time"
)

//gorm logger 需实现的接口
//type Interface interface {
//	LogMode(LogLevel) Interface
//	Info(context.Context, string, ...interface{})
//	Warn(context.Context, string, ...interface{})
//	Error(context.Context, string, ...interface{})
//	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
//}

var _ logger.Interface = (*DbLogger)(nil)

type DbLogger struct {
	Logger log.Logger
}

func (dl *DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, affected := fc()

	resp, _ := json.Marshal(map[string]interface{}{
		"affected": affected,
	})

	info := []interface{}{
		"x_msg", "",
		"x_param", sql,
		"x_response", string(resp),
		"x_action", "db.trace",
		"x_duration", time.Since(begin).Seconds(),
	}

	if err != nil {
		dl.Logger.Log(log.LevelError, info...)
	} else {
		//dl.Logger.InfoM(ctx, info)
		dl.Logger.Log(log.LevelInfo, info...)
	}
}

func (dl *DbLogger) LogMode(level logger.LogLevel) logger.Interface {
	return dl
}

func (dl *DbLogger) Warn(context.Context, string, ...interface{}) {

}

func (dl *DbLogger) Info(context.Context, string, ...interface{}) {

}

func (dl *DbLogger) Error(context.Context, string, ...interface{}) {

}
