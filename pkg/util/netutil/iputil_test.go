package netutil

import (
	"testing"
)

func TestGetLocalIPs(t *testing.T) {
	if ips, err := GetLocalIPs(); err != nil {
		t.Errorf("获取本地IP列表失败，失败原因:%+v", err)
	} else {
		t.Logf("获取本地IP列表成功，列表为:%+v", ips)
	}
}

func TestGetLocalIPsStr(t *testing.T) {
	if ips, err := GetLocalIPsStr(); err != nil {
		t.Errorf("获取本地IP列表失败，失败原因:%+v", err)
	} else {
		t.Logf("获取本地IP列表成功，列表为:%s", ips)
	}
}
