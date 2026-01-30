package logger

import (
	"fmt"
	"log/syslog"

	"github.com/samber/lo"
)

type SyslogLogger struct {
	infoLogger *syslog.Writer
	errLogger  *syslog.Writer
}

func NewSyslogLogger() *SyslogLogger {
	return &SyslogLogger{
		errLogger:  lo.Must(syslog.New(syslog.LOG_ERR, getProgramName())),
		infoLogger: lo.Must(syslog.New(syslog.LOG_INFO, getProgramName())),
	}
}

func (self *SyslogLogger) Added(s string, args ...any) {
	_ = self.infoLogger.Info("[+] " + fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Removed(s string, args ...any) {
	_ = self.infoLogger.Info("[-] " + fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Succeeded(s string, args ...any) {
	_ = self.infoLogger.Info("[✓] " + fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Failed(s string, args ...any) {
	_ = self.errLogger.Err("[𐄂] " + fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Enabled(s string, args ...any) {
	_ = self.infoLogger.Info("[✓] " + fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Disabled(s string, args ...any) {
	_ = self.errLogger.Err("[𐄂] " + fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Print(s string, args ...any) {
	_ = self.infoLogger.Info(fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Info(s string, args ...any) {
	_ = self.infoLogger.Info(fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Warn(s string, args ...any) {
	_ = self.infoLogger.Warning(fmt.Sprintf(s, args...))
}

func (self *SyslogLogger) Err(s string, args ...any) {
	_ = self.errLogger.Err(fmt.Sprintf(s, args...))
}
