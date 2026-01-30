package logger

import (
	"log"
	"os"

	"crdx.org/col"
)

type StderrLogger struct {
	logger *log.Logger
}

func NewStderrLogger() *StderrLogger {
	return &StderrLogger{
		logger: log.New(os.Stderr, "", 0),
	}
}

func (self *StderrLogger) Added(s string, args ...any) {
	self.logger.Printf(col.Green("[+] "+s), args...)
}

func (self *StderrLogger) Removed(s string, args ...any) {
	self.logger.Printf(col.Yellow("[-] "+s), args...)
}

func (self *StderrLogger) Succeeded(s string, args ...any) {
	self.logger.Printf(col.Green("[✓] "+s), args...)
}

func (self *StderrLogger) Failed(s string, args ...any) {
	self.logger.Printf(col.Red("[𐄂] "+s), args...)
}

func (self *StderrLogger) Enabled(s string, args ...any) {
	self.logger.Printf(col.Green("[✓] "+s), args...)
}

func (self *StderrLogger) Disabled(s string, args ...any) {
	self.logger.Printf(col.Red("[𐄂] "+s), args...)
}

func (self *StderrLogger) Print(s string, args ...any) {
	self.logger.Printf(s, args...)
}

func (self *StderrLogger) Info(s string, args ...any) {
	self.logger.Printf(s, args...)
}

func (self *StderrLogger) Warn(s string, args ...any) {
	self.logger.Printf(col.Yellow("Warning: "+s), args...)
}

func (self *StderrLogger) Err(s string, args ...any) {
	self.logger.Printf(col.Red("Error: "+s), args...)
}
