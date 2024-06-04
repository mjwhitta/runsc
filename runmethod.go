package runsc

type rFunc func(*state) (*state, error)

// RunMethod is simply a uint.
type RunMethod uint

var runMethods map[RunMethod]rFunc = map[RunMethod]rFunc{}
