package lggr

import (
	"log"

	"github.com/MegaBlackLabel/go-clean-architecture-testing-practice/utils/logs"
)

// // Logger インターフェイス
// type Logger interface {
// 	Errorf(format string, args ...interface{})
// 	Warnf(format string, args ...interface{})
// 	Infof(format string, args ...interface{})
// 	Debugf(format string, args ...interface{})
// }

// logの実装。ここではlogsを使っての実装。zapなどを使用する場合は別途logsインタフェースを実装する

// NewLogger adapter/loggerの実処理 [この関数をwireでDI]
func NewLogger() logs.Logger {
	return lggr{}
}

const (
	// LogPrefixError error prefix
	LogPrefixError = "[Error] "

	// LogPrefixWarn warn prefix
	LogPrefixWarn = "[Warnning] "

	// LogPrefixInfo info prefix
	LogPrefixInfo = "[Info] "

	// LogPrefixDebug debug prefix
	LogPrefixDebug = "[Debug] "
)

type lggr struct{}

func (l lggr) Errorf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixError)
	log.Printf(format, args...)
}

func (l lggr) Warnf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixWarn)
	log.Printf(format, args...)
}

func (l lggr) Infof(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixInfo)
	log.Printf(format, args...)
}

func (l lggr) Debugf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixDebug)
	log.Printf(format, args...)
}
