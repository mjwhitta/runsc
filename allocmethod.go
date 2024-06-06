package runsc

type aFunc func(*state) (*state, error)

// AllocMethod is simply a uint.
type AllocMethod uint
