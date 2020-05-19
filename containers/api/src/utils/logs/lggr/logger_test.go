package lggr

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	logger.Infof("test log Infof %v", "info")
	logger.Debugf("test log Debugf %v", "debug")
	logger.Warnf("test log Warnf %v", "warn")
	logger.Errorf("test log Errorf %v", "error")

	t.Log("loger OK")
}
