package global

import (
	"testing"
)

func TestSetupLogger(t *testing.T) {
	if err := SetupLogger("../logs/", "test.log", 600, 10); err != nil {
		t.Fatalf("配置日志异常，异常信息:%+v", err)
	} else {
		Logger.Debug("DEBUG")
		Logger.Debugf("DEBUG:%s", "test")
		Logger.Info("INFO")
		Logger.Infof("INFO:%s", "test")
		Logger.Warn("WARN")
		Logger.Warnf("WARN:%s", "test")
		Logger.Error("ERROR")
		Logger.Errorf("ERROR:%s", "test")
	}
}
