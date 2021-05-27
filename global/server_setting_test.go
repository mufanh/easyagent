package global

import "testing"

func TestSetupServerSetting(t *testing.T) {
	if err := SetupServerSetting("../configs/", "server"); err != nil {
		t.Fatalf("加载配置文件失败，详细信息:%+v", err)
	} else {
		t.Logf("加载配置文件成功，配置信息:%+v，日志配置信息:%+v", ServerConfig, ServerLogConfig)
	}
}
