// Copyright © 2023
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd
package cmd

import "testing"

func TestFormatJSONForShell(t *testing.T) {
	payload := `{"name":"app-fedcred","subject":"system:serviceaccount:default:app'oauth2"}`

	tests := []struct {
		name     string
		shell    terminalShell
		expected string
	}{
		{
			name:     "posix shell",
			shell:    terminalShellPosix,
			expected: `'{"name":"app-fedcred","subject":"system:serviceaccount:default:app'\''oauth2"}'`,
		},
		{
			name:     "powershell",
			shell:    terminalShellPowerShell,
			expected: `'{"name":"app-fedcred","subject":"system:serviceaccount:default:app''oauth2"}'`,
		},
		{
			name:     "windows cmd",
			shell:    terminalShellWindowsCmd,
			expected: `"{\"name\":\"app-fedcred\",\"subject\":\"system:serviceaccount:default:app'oauth2\"}"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := formatJSONForShell(payload, tt.shell)
			if actual != tt.expected {
				t.Fatalf("unexpected escaped payload\nexpected: %s\nactual:   %s", tt.expected, actual)
			}
		})
	}
}
