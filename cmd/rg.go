package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rgCmd)
}

var rgCmd = &cobra.Command{
	Use:   "rg FILE_EXTENSION(S) SEARCH_STRING",
	Short: "Find files containing text, filtered by file extensions",
	Long: `A streamlined wrapper around ripgrep that searches for files containing a specific string, 
filtered by file extensions. Returns a list of matching filenames rather than file contents, 
making it ideal for quickly locating files that contain specific code or text patterns.

Supports single extensions (go) or multiple comma-separated extensions (go,ts,ex) to narrow 
your search scope efficiently.`,
	Example: `  td rg go string_to_search_for
  td rg go,ts,ex multiple_comma_separated_extensions
  td rg go "string with spaces"
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		search(args[0], args[1])
	},
}

