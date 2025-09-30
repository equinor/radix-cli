package replicalog

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"
)

type colorFunc func(a ...interface{}) string

var (
	yellow  = color.New(color.FgHiYellow, color.BgBlack).SprintFunc()
	green   = color.New(color.FgHiGreen, color.BgBlack).SprintFunc()
	blue    = color.New(color.FgHiBlue, color.BgBlack).SprintFunc()
	cyan    = color.New(color.FgCyan, color.BgBlack).SprintFunc()
	red     = color.New(color.FgHiRed, color.BgBlack).SprintFunc()
	magenta = color.New(color.FgHiMagenta, color.BgBlack).SprintFunc()

	colors = []func(a ...interface{}) string{yellow, green, blue, cyan, red, magenta}

	writeMutex = sync.Mutex{}
)

// getColor Rotates color
func getColor(num int) colorFunc {
	return colors[num%len(colors)]
}

// printLine logs lines with color, safe for concurrent use by multiple goroutines
func printLine(w io.Writer, name string, logLine string, color colorFunc) {

	writeMutex.Lock()
	defer writeMutex.Unlock()

	logLine = strings.TrimSuffix(logLine, "\n")
	fmt.Fprintf(w, "[%s]: %s\n", color(name), logLine)
}
