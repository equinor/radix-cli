package log

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

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
