package capnslog

import (
	"fmt"
	"os"
)

type PackageLogger struct {
	pkg   string
	level LogLevel
}

const calldepth = 3

func (p *PackageLogger) internalLog(depth int, inLevel LogLevel, entries ...interface{}) {
	if inLevel != CRITICAL && p.level < inLevel {
		return
	}
	logger.Lock()
	defer logger.Unlock()
	if logger.formatter != nil {
		logger.formatter.Format(p.pkg, inLevel, depth+1, entries...)
	}
}

func (p *PackageLogger) LevelAt(l LogLevel) bool {
	return p.level >= l
}

// Log a formatted string at any level between ERROR and TRACE
func (p *PackageLogger) Logf(l LogLevel, format string, args ...interface{}) {
	p.internalLog(calldepth, l, fmt.Sprintf(format, args...))
}

// Log a message at any level between ERROR and TRACE
func (p *PackageLogger) Log(l LogLevel, args ...interface{}) {
	p.internalLog(calldepth, l, fmt.Sprint(args...))
}

// log stdlib compatibility

func (p *PackageLogger) Println(args ...interface{}) {
	p.internalLog(calldepth, INFO, fmt.Sprintln(args...))
}

func (p *PackageLogger) Printf(format string, args ...interface{}) {
	p.internalLog(calldepth, INFO, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Print(args ...interface{}) {
	p.internalLog(calldepth, INFO, fmt.Sprint(args...))
}

// Panic and fatal

func (p *PackageLogger) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	p.internalLog(calldepth, CRITICAL, s)
	panic(s)
}

func (p *PackageLogger) Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	p.internalLog(calldepth, CRITICAL, s)
	panic(s)
}

func (p *PackageLogger) Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	p.internalLog(calldepth, CRITICAL, s)
	os.Exit(1)
}

func (p *PackageLogger) Fatal(args ...interface{}) {
	s := fmt.Sprint(args...)
	p.internalLog(calldepth, CRITICAL, s)
	os.Exit(1)
}

// Error Functions

func (p *PackageLogger) Errorf(format string, args ...interface{}) {
	p.internalLog(calldepth, ERROR, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Error(entries ...interface{}) {
	p.internalLog(calldepth, ERROR, entries...)
}

// Warning Functions

func (p *PackageLogger) Warningf(format string, args ...interface{}) {
	p.internalLog(calldepth, WARNING, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Warning(entries ...interface{}) {
	p.internalLog(calldepth, WARNING, entries...)
}

// Notice Functions

func (p *PackageLogger) Noticef(format string, args ...interface{}) {
	p.internalLog(calldepth, NOTICE, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Notice(entries ...interface{}) {
	p.internalLog(calldepth, NOTICE, entries...)
}

// Info Functions

func (p *PackageLogger) Infof(format string, args ...interface{}) {
	p.internalLog(calldepth, INFO, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Info(entries ...interface{}) {
	p.internalLog(calldepth, INFO, entries...)
}

// Debug Functions

func (p *PackageLogger) Debugf(format string, args ...interface{}) {
	p.internalLog(calldepth, DEBUG, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Debug(entries ...interface{}) {
	p.internalLog(calldepth, DEBUG, entries...)
}

// Trace Functions

func (p *PackageLogger) Tracef(format string, args ...interface{}) {
	p.internalLog(calldepth, TRACE, fmt.Sprintf(format, args...))
}

func (p *PackageLogger) Trace(entries ...interface{}) {
	p.internalLog(calldepth, TRACE, entries...)
}

func (p *PackageLogger) Flush() {
	logger.Lock()
	defer logger.Unlock()
	logger.formatter.Flush()
}
