package conf

import (
	"github.com/Penglq/QLog"
	"os"
)

func InitConf() {
	os.Getenv("AUTH_CENTER_ENV")
}

type AppConf struct {
	Env    string `json:"env"`
	Port   string `json:"port"`
	Logger Logger `json:"logger"`
}

type Logger struct {
	LogLevel     uint8
	IsConsole    bool
	IsFile       bool
	FilePath     string
	Filename     string
	FileSuffix   string
	FileMaxSize  int64
	FileMaxNSize int
	AlertConf    QLog.AlertApiConfig
}
