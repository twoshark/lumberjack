package lumberjack

import (
	"encoding/json"
	"runtime"
)

//WarningJSONParse -
var WarningJSONParse = NewError(-1, "Failed to Un/marshal or stringify JSON")

//GetCaller - return the name of caller of the current method
func GetCaller() string {
	// Skip GetCallerFunctionName and the function to get the caller of
	return getFrame(2).Function
}

//GetRuntimeFrameFunction -
func GetRuntimeFrameFunction(framesBack int) string {
	// Skip GetCallerFunctionName and the function to get the caller of
	return getFrame(framesBack + 1).Function
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

//JSON Marshals an object into a byte array
func JSON(obj interface{}) (out []byte, err error) {
	out, err = json.Marshal(obj)
	if err != nil {
		return
	}
	return
}

//StringJSON returns a stringified of a json string representing obj
func StringJSON(obj interface{}) (out string, err error) {
	bytes, err := JSON(obj)
	if err != nil {
		out = "ObjectFailedToParse"
		return
	}
	return string(bytes), err
}
