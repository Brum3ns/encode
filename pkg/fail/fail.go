package fail

// List of errorcodes and their output:
type Errorcode int

var ERRORCODES = map[int]string{
	1001: "Invalid Stdin input was given",
	1002: "The encoder could not be found",
	1003: "Invalid number of threads",
}
