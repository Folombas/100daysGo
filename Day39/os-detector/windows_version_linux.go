//go:build !windows

package main

// getWindowsDetails возвращает заглушку для не-Windows систем
func getWindowsDetails() string {
	return "Windows"
}