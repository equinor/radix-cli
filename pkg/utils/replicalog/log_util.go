package replicalog

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"
)

type ColorFunc func(a ...interface{}) string

var (
	Yellow  = color.New(color.FgHiYellow, color.BgBlack).SprintFunc()
	Green   = color.New(color.FgHiGreen, color.BgBlack).SprintFunc()
	Blue    = color.New(color.FgHiBlue, color.BgBlack).SprintFunc()
	Cyan    = color.New(color.FgCyan, color.BgBlack).SprintFunc()
	Red     = color.New(color.FgHiRed, color.BgBlack).SprintFunc()
	Magenta = color.New(color.FgHiMagenta, color.BgBlack).SprintFunc()

	Colors = []func(a ...interface{}) string{Yellow, Green, Blue, Cyan, Red, Magenta}

	writeMutex = sync.Mutex{}
)

// GetColor Rotates color
func GetColor(num int) ColorFunc {
	return Colors[num%len(Colors)]
}

// PrintLines logs lines with color, safe for concurrent use by multiple goroutines
func PrintLines(w io.Writer, name string, previousLogLines, logLines []string, color ColorFunc) {
	for _, logLine := range logLines {
		if !logged(logLine, previousLogLines) {
			PrintLine(w, name, logLine, color)
		}
	}
}

// PrintLine logs lines with color, safe for concurrent use by multiple goroutines
func PrintLine(w io.Writer, name string, logLine string, color ColorFunc) {
	print(w, name, logLine, color)
}

func logged(logLine string, previousLogLines []string) bool {
	for _, previousLogLine := range previousLogLines {
		if strings.EqualFold(previousLogLine, logLine) {
			return true
		}
	}
	return false
}

// print Output string to standard output, safe for concurrent use by multiple goroutines
func print(w io.Writer, name, logLine string, color ColorFunc) {
	writeMutex.Lock()
	defer writeMutex.Unlock()

	logLine = strings.TrimSuffix(logLine, "\n")
	fmt.Fprintf(w, "[%s]: %s\n", color(name), logLine)
}
