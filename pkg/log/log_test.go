// author: maxf
// date: 2021-03-05 13:27
// version:

package log

import (
	"collector-sidecar-server/pkg/config"
	"collector-sidecar-server/pkg/constant"
	"testing"
)

func init() {
	defer Sync()
	c := config.Config{
		LogConfig: config.LogConfig{
			Level:      "debug",
			FileName:   "test.log",
			TimeFormat: constant.TimeLayout,
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     2,
			Compress:   false,
			LocalTime:  true,
			Console:    true,
		},
		AppName: "zapTest",
	}
	InitLogger(&(c.LogConfig), c.AppName)
}

func TestInfo(t *testing.T) {
	Info("test info", Pair("age", 20), Pair("name", "小明"))
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("test info", Pair("age", 20), Pair("name", "小明"))
	}
}
