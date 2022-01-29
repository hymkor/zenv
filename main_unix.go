//go:build !windows
// +build !windows

package main

func addDllDirectories(dllDirectories ...string) error {
	return nil
}
