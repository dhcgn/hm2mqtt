package shared

// Event represents one event from the rpc interface
type Event struct {
	MethodName   string
	SerialNumber string
	Type         string
	DataValue    string
}
