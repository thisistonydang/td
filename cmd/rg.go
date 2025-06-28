package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rgCmd)
}

var rgCmd = &cobra.Command{
	Use:   "rg [FILE_EXTENSION(S)] SEARCH_STRING",
	Short: "Find files containing text, optionally filtered by file extensions",
	Long: `A streamlined wrapper around ripgrep that searches for files containing a specific string, 
optionally filtered by file extensions. Returns a list of matching filenames rather than file contents, 
making it ideal for quickly locating files that contain specific code or text patterns.

Supports single extensions (go) or multiple comma-separated extensions (go,ts,ex) to narrow your search 
scope efficiently. If no file extensions are provided, all files types will be searched.`,
	Example: `  td rg go string_to_search_for
  td rg go,ts,ex multiple_comma_separated_extensions
  td rg go "string with spaces"
`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			search("", args[0])
		} else {
			search(args[0], args[1])
		}
	},
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
}

func search(fileExtensions string, searchString string) {
	// Construct the rg command
	args := []string{"--color=always", "--files-with-matches", "--sort=path"}
	if fileExtensions != "" {
		args = append(args, "--glob=*.{"+fileExtensions+"}")
	}
	args = append(
		args,
		"--", // This is necessary to escape strings that may be interpreted as flags (e.g. --color)
		searchString,
	)
	cmd := exec.Command("rg", args...)

	// Captured stdout from rg command
	var stdout strings.Builder
	cmd.Stdout = &stdout

	// Captured stderr from rg command
	var stderr strings.Builder
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		// Check if it's an exit error to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
	}

	// Print the output captured from rg
	fmt.Print(stdout.String())
}
