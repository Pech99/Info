package main

import (
	"fmt"

	"github.com/Pech99/Info/us32"
	"golang.design/x/clipboard"
)

func main() {
	// https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ee391646(v=vs.85)

	hwnd := us32.GetForegroundWindow()
	clipboard.Write(clipboard.FmtText, []byte(us32.GetClassNameA(us32.HWND(hwnd))))

	info := fmt.Sprintln(
		" HWND:", hwnd, "\n",
		"ClassNameW:", us32.GetClassNameW(us32.HWND(hwnd)), "\n",
		"ClassNameA:", us32.GetClassNameA(us32.HWND(hwnd)), "\n",
		"WindowText:", us32.GetWindowText(us32.HWND(hwnd)),
	)
	us32.MessageBox(0, info, "Foreground Window Info", 0)

}
