//go:build !windows
// +build !windows

package runsc

import "gitlab.com/mjwhitta/errors"

// WithNtAllocateVirtualMemory is only supported for Windows
func WithNtAllocateVirtualMemory(pid uint32, sc []byte) error {
	return errors.New("unsupported OS")
}

// WithNtCreateSection is only supported for Windows
func WithNtCreateSection(pid uint32, sc []byte) error {
	return errors.New("unsupported OS")
}

// WithNtQueueApcThread is only supported for Windows
func WithNtQueueApcThread(pid uint32, sc []byte) error {
	return errors.New("unsupported OS")
}

// WithNtQueueApcThreadEx is only supported for Windows
func WithNtQueueApcThreadEx(pid uint32, sc []byte) error {
	return errors.New("unsupported OS")
}
