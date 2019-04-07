package libs

import "github.com/Penglq/QLog"

func InitQLog() QLog.LoggerInterface {
	l := QLog.GetLogger()
	l.SetConfig(QLog.DEBUG, "",
		QLog.WithConsoleOPT(),
	)
	return l
}
