//go:build !windows

package runsc

// Consts for supported allocation methods.
// const (
// 	InvalidAlloc AllocMethod = iota + 1
// )

var allocMethods map[AllocMethod]aFunc = map[AllocMethod]aFunc{}
