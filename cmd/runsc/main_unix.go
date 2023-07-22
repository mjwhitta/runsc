//go:build !windows

package main

import hl "github.com/mjwhitta/hilighter"

func launch(sc []byte) error {
	return hl.Errorf("unsupported OS")
}
