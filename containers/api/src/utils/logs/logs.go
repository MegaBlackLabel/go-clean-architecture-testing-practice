//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/MegaBlackLabel/go-clean-architecture-testing-practice/utils/$GOPACKAGE

package logs

// ここではログ用のインターフェイスを定義utils/logsで実装する

// Logger インターフェイス
type Logger interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}
