package logger

/*
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	//"realworld/pkg/tracer"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	logger *logrus.Logger
}

var _ logrus.Formatter = (*LogFormatter)(nil)

type LogFormatter struct {
}

//{
//"x_param": "{\"post\":[],\"get\":{\"url\":\"usercoupon\\/newflag\\/del\"},\"body\":\"{\\\"app_id\\\":\\\"1\\\",\\\"user_id\\\":\\\"25459453\\\"}\"}",
//"x_response": "{\"stat\":1,\"msg\":\"succ\",\"data\":{},\"trace_id\":\"dayu_bb1189afb61941588db757d7b34ede8b\"}",
//"x_action": "cardcouponv2api.xesv5.com/usercoupon/newflag/del",
//"x_code": "200",
//"x_client_ip": "10.20.48.53",
//"x_duration": "0.0127",
//"x_msg": "",

//"x_trace_id": "dayu_bb1189afb61941588db757d7b34ede8b",
//"x_rpc_id": "1.4.1",

//"x_name": "log.info",
//"x_source": "",
//"x_timestamp": 1623467618,
//"x_server_ip": "bj-sjhl-api-cardcoupon-online-2",
//"x_version": "php-0.3",
//"x_department": "tal_dt_php"
//}
func (lf *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+6)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise err_encoder are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	host, _ := os.Hostname()
	data["x_name"] = "log." + entry.Level.String()
	data["x_department"] = "golang"
	data["x_hostname"] = host
	data["x_timestamp"] = time.Now().Unix()
	data["x_date"] = time.Now().Format("2006-01-02 15:04:05")
	data["x_version"] = "0.1"

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v\n", err)
	}

	return b.Bytes(), nil
}

func New(lvl string) *Logger {
	l := &Logger{}
	l.logger = logrus.New()
	l.logger.SetFormatter(&LogFormatter{})

	if lvl, err := logrus.ParseLevel(lvl); err != nil {
		panic(err.Error())
	} else {
		l.logger.SetLevel(lvl)
	}

	return l
}

func (l *Logger) SetOutput(writer io.Writer) *Logger {
	l.logger.SetOutput(writer)
	return l
}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.InfoM(ctx, map[string]interface{}{
		"x_msg": fmt.Sprintf(format, args...),
	})
}

func (l *Logger) InfoM(ctx context.Context, msgs map[string]interface{}) {
	//tn := tracer.ExtractTraceNode(ctx)

	//msgs["x_trace_id"] = tn.TraceId()
	//msgs["x_rpc_id"] = tn.RpcId()
	msgs["x_source"] = source()

	l.logger.WithFields(msgs).Info()

	//tn.IncrRpcId()
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.ErrorM(ctx, map[string]interface{}{
		"x_msg": fmt.Sprintf(format, args...),
	})
}

func (l *Logger) ErrorM(ctx context.Context, msgs map[string]interface{}) {
	//tn := tracer.ExtractTraceNode(ctx)
	frames, _ := json.Marshal(callFrames(15))

	//msgs["x_trace_id"] = tn.TraceId()
	//msgs["x_rpc_id"] = tn.RpcId()
	msgs["x_source"] = source()
	msgs["x_extra"] = string(frames)

	l.logger.WithFields(msgs).Error()

	//tn.IncrRpcId()
}

var packageName string
var once sync.Once

func source() string {
	return callFrames(10)[0]
}

func callFrames(maxDept int) []string {
	var stacks []string
	pcs := make([]uintptr, maxDept)
	depth := runtime.Callers(1, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	// cache this package's fully-qualified name
	once.Do(func() {
		packageName = getPackageName(runtime.FuncForPC(pcs[0]).Name())
	})

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != packageName {
			stacks = append(stacks, fmt.Sprintf("%s:%d", f.File, f.Line))
		}
	}

	return stacks
}

func getPackageName(absPath string) string {
	for {
		lastPeriod := strings.LastIndex(absPath, ".")
		lastSlash := strings.LastIndex(absPath, "/")
		if lastPeriod > lastSlash {
			absPath = absPath[:lastPeriod]
		} else {
			break
		}
	}

	return absPath
}
*/
