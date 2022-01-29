package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	_LOAD_LIBRARY_SEARCH_DEFAULT_DIRS = 0x00001000
)

var flagDllDir = flag.String("L", "", "Set DLL directories list (seperated with "+string(os.PathListSeparator)+")")

var kernel32 = windows.NewLazySystemDLL("kernel32")
var procSetDefaultDllDirectories = kernel32.NewProc("SetDefaultDllDirectories")
var procAddDllDirectory = kernel32.NewProc("AddDllDirectory")

func failed(err error) bool {
	if err == nil {
		return false
	}
	if errno, ok := err.(syscall.Errno); ok && errno == 0 {
		return false
	}
	return true
}

func setDefaultDllDirectories() error {
	_, _, err := procSetDefaultDllDirectories.Call(_LOAD_LIBRARY_SEARCH_DEFAULT_DIRS)
	if failed(err) {
		return fmt.Errorf("SetDefaultDllDirectories: %w(%[1]t)", err)
	}
	return nil
}

func addDllDirectory(path string) error {
	pathW, err := windows.UTF16FromString(path)
	if failed(err) {
		return fmt.Errorf("UTF16FromString: %w", err)
	}
	_, _, err = procAddDllDirectory.Call(uintptr(unsafe.Pointer(&pathW[0])))
	if failed(err) {
		return fmt.Errorf(`AddDllDirectory: %w`, err)
	}
	return nil
}

func addDllDirectories(dllDirectories ...string) error {
	if err := setDefaultDllDirectories(); err != nil {
		return err
	}
	for _, dllDir := range dllDirectories {
		if err := addDllDirectory(dllDir); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", dllDir, err.Error())
		}
	}
	return nil
}
