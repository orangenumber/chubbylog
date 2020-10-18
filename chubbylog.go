// Written by Gon Yi

package chubbylog

import (
	"github.com/gonyyi/alog"
	"io"
	"log/syslog"
	"os"
)

const F_TIME = alog.F_TIME
const F_MMDD = alog.F_MMDD
const F_MICROSEC = alog.F_MICROSEC
const F_PREFIX = alog.F_PREFIX
const F_UTC = alog.F_UTC
const F_DATE = alog.F_DATE
const F_USE_BUF_1K = alog.F_USE_BUF_1K
const F_USE_BUF_2K = alog.F_USE_BUF_2K
const F_STD = alog.F_STD

// =====================================================================================================================
// CHUBBY LOG
// =====================================================================================================================
type AChubbyLogger struct {
	info  *alog.ALogger
	warn  *alog.ALogger
	error *alog.ALogger
	fatal *alog.ALogger
}

func (l *AChubbyLogger) Info(s ...interface{}) {
	l.info.Print(s...)
}
func (l *AChubbyLogger) Infof(format string, s ...interface{}) {
	l.info.Printf(format, s...)
}
func (l *AChubbyLogger) Infoj(prefix string, a interface{}) {
	l.info.Printj(prefix, a)
}
func (l *AChubbyLogger) Warn(s ...interface{}) {
	l.warn.Print(s...)
}
func (l *AChubbyLogger) Warnf(format string, s ...interface{}) {
	l.warn.Printf(format, s...)
}
func (l *AChubbyLogger) Warnj(prefix string, a interface{}) {
	l.warn.Printj(prefix, a)
}

func (l *AChubbyLogger) Error(s ...interface{}) {
	l.error.Print(s...)
}
func (l *AChubbyLogger) Errorf(format string, s ...interface{}) {
	l.error.Printf(format, s...)
}
func (l *AChubbyLogger) Errorj(prefix string, a interface{}) {
	l.error.Printj(prefix, a)
}

func (l *AChubbyLogger) Fatal(s ...interface{}) {
	l.fatal.Print(s...)
	l.close()
	os.Exit(1)
}
func (l *AChubbyLogger) Fatalf(format string, s ...interface{}) {
	l.fatal.Printf(format, s...)
	l.close()
	os.Exit(1)
}
func (l *AChubbyLogger) Fatalj(prefix string, a interface{}) {
	l.fatal.Printj(prefix, a)
	l.close()
	os.Exit(1)
}
func (l *AChubbyLogger) GetInfo() *alog.ALogger {
	return l.info
}
func (l *AChubbyLogger) GetWarn() *alog.ALogger {
	return l.warn
}
func (l *AChubbyLogger) GetError() *alog.ALogger {
	return l.error
}
func (l *AChubbyLogger) GetFatal() *alog.ALogger {
	return l.fatal
}
func (l *AChubbyLogger) ToSyslog(name string) error {
	fSyslog := func(acl *alog.ALogger, priority syslog.Priority) error {
		out, err := syslog.New(priority, name)
		acl.SetOutput(out)
		if err != nil {
			return err
		}
		acl.SetPrefix("") // syslog already have prefix
		return nil
	}
	if err := fSyslog(l.info, syslog.LOG_INFO); err != nil { return err }
	if err := fSyslog(l.warn, syslog.LOG_WARNING); err != nil { return err }
	if err := fSyslog(l.error, syslog.LOG_ERR); err != nil { return err }
	if err := fSyslog(l.fatal, syslog.LOG_CRIT); err != nil { return err }

	return nil
}

func (l *AChubbyLogger) close() {
	l.info.Close()
	l.warn.Close()
	l.error.Close()
	l.fatal.Close()
}
func New(output io.Writer, flag uint16) *AChubbyLogger {
	if output == nil {
		output = os.Stdout
	}

	return &AChubbyLogger{
		info:  alog.New(output, "[INFO]  ", flag),
		warn:  alog.New(output, "[WARN]  ", flag),
		error: alog.New(output, "[ERROR] ", flag),
		fatal: alog.New(output, "[FATAL] ", flag),
	}
}
