package lumberjack

//Geoffrey is chill
var Geoffrey Lumberjack

//Lumberjack - our logger
type Lumberjack struct {
	//The Trap catches and handles Errors
	Trap ErrorHandler
	//The Axe makes Logs
	Axe LoggingHandler
}

//LumberData -
type LumberData interface {
	insert(interface{})
}

//Init -
func (l *Lumberjack) Init() {
}
