//go:build !windows
// +build !windows

package runsc

import hl "gitlab.com/mjwhitta/hilighter"

// WithNtAllocateVirtualMemory is only supported for Windows
func WithNtAllocateVirtualMemory(pid uint32, sc []byte) error {
	return hl.Errorf("runsc: unsupported OS")
}

// WithNtCreateSection is only supported for Windows
func WithNtCreateSection(pid uint32, sc []byte) error {
	return hl.Errorf("runsc: unsupported OS")
}

// WithNtQueueApcThread is only supported for Windows
func WithNtQueueApcThread(pid uint32, sc []byte) error {
	return hl.Errorf("runsc: unsupported OS")
}

// WithNtQueueApcThreadEx is only supported for Windows
func WithNtQueueApcThreadEx(pid uint32, sc []byte) error {
	return hl.Errorf("runsc: unsupported OS")
}
