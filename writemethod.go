package runsc

type wFunc func(*state, []byte) (*state, error)

// WriteMethod is simply a uint.
type WriteMethod uint

var writeMethods map[WriteMethod]wFunc = map[WriteMethod]wFunc{}
