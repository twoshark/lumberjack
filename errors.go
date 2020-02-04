package lumberjack

import (
	"log"
	"time"
)

//ErrorHandler -
type ErrorHandler []BadThing

//BadThing -
type BadThing struct {
	Code     int `yaml:"code"`
	ErrorObj interface{}
	Log      `yaml:"inline"`
}

//NewError - make a new Error
func NewError(code int, msgs ...string) BadThing {
	notGoodThing := BadThing{
		Code: code,
	}
	notGoodThing.Messages = msgs
	return notGoodThing
}

//Panic -
func (e *ErrorHandler) Panic(notGoodThing BadThing) {
	//Add Runtime details
	notGoodThing.CreateDate = time.Now()
	notGoodThing.Caller = GetCaller()
	//Insert pointer to completed error into the error log
	e.insert(&notGoodThing)
	panic(notGoodThing)
}

//UhOh - several new GoFireError to the log and Print it to the terminal
func (e *ErrorHandler) UhOh(notGoodThings ...BadThing) {
	//Local errors
	for _, notGoodThing := range notGoodThings {
		//Add Runtime details
		notGoodThing.CreateDate = time.Now()
		notGoodThing.Caller = GetCaller()
		//Insert pointer to completed error into the error log
		e.insert(&notGoodThing)
		json, err := StringJSON(notGoodThing)
		if err != nil {
			warning := WarningJSONParse
			warning.ErrorObj = notGoodThing
			e.insert(warning)
			return
		}
		log.Println(json)
	}
}

//insertError -
func (e *ErrorHandler) insert(notGoodThing interface{}) {
	badThing, ok := notGoodThing.(BadThing)
	if ok {
		l := len(*e)
		target := *e
		if cap(*e) == l {
			target = make([]BadThing, l+1, l+10)
			copy(target, *e)
			target[l] = badThing
		} else {
			target = append(target, badThing)
		}
		e = &target
		return
	}
	warning := NewError(-1, "Failed to Insert Error")
	warning.ErrorObj = notGoodThing
	Geoffrey.Trap.UhOh(warning)
}
