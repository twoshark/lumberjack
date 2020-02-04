package lumberjack

import (
	"time"
)

//LoggingHandler -
type LoggingHandler []Log

//Log -
type Log struct {
	CreateDate time.Time
	Caller     string
	Messages   []string
}

//Log - it's big, it's heavy, it's wood. It's better than bad... it's good!
func (lh *LoggingHandler) Log(msgs ...string) {
	logEntry := Log{
		CreateDate: time.Now(),
		Caller:     GetCaller(),
		Messages:   msgs,
	}
	lh.insert(logEntry)
}

func (lh *LoggingHandler) insert(entry interface{}) {
	logEntry, ok := (entry).(Log)
	if ok {
		l := len(*lh)
		target := *lh
		if cap(*lh) == l {
			target = make([]Log, l+1, l+10)
			copy(target, *lh)
			target[l] = logEntry
		} else {
			target = append(target, logEntry)
		}
		lh = &target
		return
	}

	warning := WarningJSONParse
	warning.ErrorObj = entry
	Geoffrey.Trap.UhOh(warning)
}
