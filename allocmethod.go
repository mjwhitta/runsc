package runsc

type aFunc func(*state) (*state, error)

// AllocMethod is simply a uint.
type AllocMethod uint

var allocMethods map[AllocMethod]aFunc = map[AllocMethod]aFunc{}
