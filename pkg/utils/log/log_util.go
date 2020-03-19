package log

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	Yellow  = color.New(color.FgHiYellow, color.BgBlack, color.Bold).SprintFunc()
	Green   = color.New(color.FgHiGreen, color.BgBlack, color.Bold).SprintFunc()
	Blue    = color.New(color.FgHiBlue, color.BgBlack, color.Underline).SprintFunc()
	Cyan    = color.New(color.FgCyan, color.BgBlack).SprintFunc()
	Red     = color.New(color.FgHiRed, color.BgBlack).Add(color.Italic).SprintFunc()
	Magenta = color.New(color.FgHiMagenta, color.BgBlack).Add(color.Italic).SprintFunc()

	Colors = []func(a ...interface{}) string{Yellow, Green, Blue, Cyan, Red, Magenta}
)

// GetColor Rotates color
func GetColor(num int) func(a ...interface{}) string {
	return Colors[num%len(Colors)]
}

// PrintLines logs lines with color
func PrintLines(cmd *cobra.Command, name string, previousLogLines, logLines []string, color func(a ...interface{}) string) {
	for _, logLine := range logLines {
		if !logged(logLine, previousLogLines) {
			Print(cmd, name, logLine, color)
		}
	}
}

func logged(logLine string, previousLogLines []string) bool {
	for _, previousLogLine := range previousLogLines {
		if strings.EqualFold(previousLogLine, logLine) {
			return true
		}
	}
	return false
}

// Print Output string to standard output
func Print(cmd *cobra.Command, name, logLine string, color func(a ...interface{}) string) {
	fmt.Fprintf(cmd.OutOrStdout(), "\r\n[%s] : %s", color(name), color(logLine))
}
