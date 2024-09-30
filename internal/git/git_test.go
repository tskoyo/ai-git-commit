package git

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGitRepo(t *testing.T) {
	execCommand = func(name string, arg ...string) *exec.Cmd {
		if name == "git" && strings.Join(arg, " ") == "rev-parse --is-inside-work-tree" {
			return exec.Command("echo", "true")
		}
		return exec.Command("false")
	}

	err := isGitRepo()
	assert.NoError(t, err, "expected nil, got error: %v", err)
}

func TestGetStagedDiff(t *testing.T) {
	execCommand = func(name string, arg ...string) *exec.Cmd {
		if name == "git" && strings.Join(arg, " ") == "diff --cached" {
			return exec.Command("echo", `diff --git a/main.go b/main.go
				index abcdef..123456 100644
				--- a/main.go
				+++ b/main.go
				@@ -1,4 +1,4 @@
				-package main
				+package main // Modified package declaration
				func main() {
						// Some code
				}`)
		}
		return exec.Command("false")
	}

	diff, err := getStagedDiff()
	assert.NoError(t, err, "expected nil, got error: %v", err)
	assert.Contains(t, diff, "Modified package declaration", "expected diff to contain changes, got: %s", diff)
}

func TestFormatDiff(t *testing.T) {
	diff := `diff --git a/main.go b/main.go
index abcdef..123456 100644
--- a/main.go
+++ b/main.go
@@ -1,4 +1,4 @@
-package main
+package main // Modified package declaration
func main() {
		// Some code
}`

	expectedOutput := "--- a/main.go +++ b/main.go @@ -1,4 +1,4 @@ -package main +package main // Modified package declaration func main() { // Some code }"
	formattedDiff := formatDiff(diff)

	assert.Equal(t, expectedOutput, formattedDiff, "expected formatted diff to match, got: %s", formattedDiff)
}
