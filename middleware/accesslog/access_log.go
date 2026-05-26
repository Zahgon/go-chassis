package accesslog

import (
	"fmt"
	"time"

	"github.com/go-chassis/openlog"

	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/lager"
	"github.com/go-chassis/go-chassis/v2/initiator"
)

// Record recorder
type Record func(startTime time.Time, i *invocation.Invocation)

var (
	instance = &accessLog{
		record: restfulRecord,
	}

	log openlog.Logger
)

const handlerNameAccessLog = "access-log"

func init() {
	if initiator.LoggerOptions == nil || len(initiator.LoggerOptions.AccessLogFile) == 0 {
		openlog.Info("lager.yaml non exist, skip init")
		return
	}

	if initiator.LoggerOptions.AccessLogFile == lager.Stdout {
		log = openlog.GetLogger()
	} else {
		var err error
		opts := &lager.Options{
			Writers:           lager.File,
			LoggerLevel:       lager.LevelInfo,
			LoggerFile:        initiator.LoggerOptions.AccessLogFile,
			LogFormatText:     initiator.LoggerOptions.LogFormatText,
			LogRotateAge:      initiator.LoggerOptions.LogRotateAge,
			LogRotateSize:     initiator.LoggerOptions.LogRotateSize,
			LogBackupCount:    initiator.LoggerOptions.LogBackupCount,
			LogRotateCompress: initiator.LoggerOptions.LogRotateCompress,
			LogRotateDisable:  initiator.LoggerOptions.LogRotateDisable,
		}
		log, err = lager.NewLog(opts)
		if err != nil {
			openlog.Error(fmt.Sprintf("new access log failed, %s", err.Error()))
			return
		}
	}

	err := handler.RegisterHandler(handlerNameAccessLog, func() handler.Handler {
		return instance
	})
	if err != nil {
		openlog.Error(fmt.Sprintf("register access log handler failed, %s", err.Error()))
	}
}

// Use support customize recorder
func Use(record Record) { _ = "STUB: not implemented"; return }

type accessLog struct {
	record func(time.Time, *invocation.Invocation)
}

// Handle ...
func (a *accessLog) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// Name ...
func (a *accessLog) Name() string { _ = "STUB: not implemented"; return "" }

func restfulRecord(startTime time.Time, i *invocation.Invocation) {
	_ = "STUB: not implemented"
	return
}
