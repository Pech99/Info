package us32

import (
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.MustLoadDLL("user32.dll")
	procGetWindowText       = user32.MustFindProc("GetWindowTextW")
	procGetWindowTextLength = user32.MustFindProc("GetWindowTextLengthW")
	procGetForegroundWindow = user32.MustFindProc("GetForegroundWindow")
	procMessageBox          = user32.MustFindProc("MessageBoxW")
	procGetClassNameW       = user32.MustFindProc("GetClassNameW")
	procGetClassNameA       = user32.MustFindProc("GetClassNameA")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

// restituisce la lunghezza del testo della finiesra con id hwnd
func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

// restituisce il testo della finiesra con id hwnd
func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen),
	)

	return syscall.UTF16ToString(buf)
}

func GetForegroundWindow() uintptr {
	hwnd, _, _ := procGetForegroundWindow.Call()
	return hwnd
}

func MessageBox(hwnd HWND, title, caption string, flags uint) int {
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags),
	)

	return int(ret)
}

func GetClassNameW(hwnd HWND) string {
	buf := make([]uint16, 255)
	procGetClassNameW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(255))

	return syscall.UTF16ToString(buf)
}

func GetClassNameA(hwnd HWND) string {
	buf := make([]uint16, 255)
	procGetClassNameA.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(255))

	return syscall.UTF16ToString(buf)
}
