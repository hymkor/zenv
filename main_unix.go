//go:build !windows
// +build !windows

package main

var flagDllDir = new(string)

func addDllDirectories(dllDirectories ...string) error {
	return nil
}
