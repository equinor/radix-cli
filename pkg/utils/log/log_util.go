package log

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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
func GetColor(num int) func(a ...interface{}) string {
	return Colors[num%len(Colors)]
}

// PrintLines logs lines with color, safe for concurrent use by multiple goroutines
func PrintLines(cmd *cobra.Command, name string, previousLogLines, logLines []string, color func(a ...interface{}) string) {
	for _, logLine := range logLines {
		if !logged(logLine, previousLogLines) {
			PrintLine(cmd, name, logLine, color)
		}
	}
}

// PrintLine logs lines with color, safe for concurrent use by multiple goroutines
func PrintLine(cmd *cobra.Command, name string, logLine string, color func(a ...interface{}) string) {
	print(cmd, name, logLine, color)
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
func print(cmd *cobra.Command, name, logLine string, color func(a ...interface{}) string) {
	writeMutex.Lock()
	defer writeMutex.Unlock()

	fmt.Fprintf(cmd.OutOrStdout(), "\r\n[%s]: %s", color(name), logLine)
}
