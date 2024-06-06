package runsc

type rFunc func(*state) (*state, error)

// RunMethod is simply a uint.
type RunMethod uint
