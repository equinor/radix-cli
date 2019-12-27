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

func GetColor(num int) func(a ...interface{}) string {
	return Colors[num%len(Colors)]
}

// From logs lines exceeding from
func From(cmd *cobra.Command, name string, from int, logLines []string, color func(a ...interface{}) string) int {
	logged := 0

	for num, logLine := range logLines {
		if num >= from {
			if !strings.EqualFold(strings.TrimSpace(logLine), "") && from > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "\r\n[%s] : %s", color(name), color(logLine))
			}

			logged++
		}
	}

	return logged
}
