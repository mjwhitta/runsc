//+build !windows

package runsc

import hl "gitlab.com/mjwhitta/hilighter"

// WithNtAllocateVirtualMemory is only supported for Windows
func WithNtAllocateVirtualMemory(pid uint32, sc []byte) error {
	return hl.Errorf("Unsupported OS")
}

// WithNtCreateSection is only supported for Windows
func WithNtCreateSection(pid uint32, sc []byte) error {
	return hl.Errorf("Unsupported OS")
}

// WithNtQueueApcThread is only supported for Windows
func WithNtQueueApcThread(pid uint32, sc []byte) error {
	return hl.Errorf("Unsupported OS")
}

// WithNtQueueApcThreadEx is only supported for Windows
func WithNtQueueApcThreadEx(pid uint32, sc []byte) error {
	return hl.Errorf("Unsupported OS")
}
