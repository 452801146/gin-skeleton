package initialize


import (
	"gin_skeleton/g"
	"io"
	"os"
)

var logDir = "./logs"

// 初始化日志路径
func InitLogDir() {
	// 判断文件夹是否存在
	_, err := os.Stat(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(logDir, 0666)
			if err != nil {
				panic("mkdir ./log  err")
			}
		}
	}
}

// 获得日志写入文件
func GetLogWriter(configName string) io.Writer {
	// 日志文件
	logPath := logDir+"/" + g.Config.GetString(configName)
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	return io.MultiWriter(file)
}
